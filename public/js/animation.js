function OnLoadAnimation() {
  gsap.from("#create-todo", { opacity: 0, y: 100, duration: 0.5 });
}

const buttons = document.querySelectorAll(".deleteanimated");
buttons.forEach((element, i) => {
  console.log("element " + i);
  element.addEventListener("click", () => {
    console.log(element.id);
  });
});

window.addEventListener("load", OnLoadAnimation);
