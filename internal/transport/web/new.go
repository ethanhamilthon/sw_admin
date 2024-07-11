package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"templtest/internal/services"

	"time"
)

type Server struct {
	mux     *http.ServeMux
	service *services.RegisterService
}

func New(service *services.RegisterService) *Server {
	s := &Server{service: service}
	s.register()
	return s
}

func (s *Server) register() {
	handlers := http.NewServeMux()
	main := http.NewServeMux()
	fsHandler := http.FileServer(http.Dir("./public"))
	main.Handle("/public/", http.StripPrefix("/public/", fsHandler))
	main.HandleFunc("/", s.index)
	handlers.Handle("/metrics/", http.StripPrefix("/metrics", main))
	s.mux = handlers
}

func (s *Server) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal to close app
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":4000",
		Handler: s.mux,
	}

	go func() {
		log.Printf("Starting server on port 4000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-osSignal
	log.Println("Shutdown signal received. Shutting down...")

	// Timeout to close handlers
	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
	defer shutdownCancel()

	// Stop the server
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

}
