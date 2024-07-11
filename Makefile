run: 
	./runner.sh

css:
	npx tailwindcss build -i views/css/main.css -o public/style.css --watch

templ:
	templ generate --watch 
