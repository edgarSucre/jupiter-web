document.addEventListener("DOMContentLoaded", () => {
    const toggleDarkModeCheck = document.querySelector("#darkModeToggle");

    toggleDarkModeCheck.addEventListener("change", (e) => {
        component = document.body;
        component.classList.toggle("dark");
        e.stopPropagation();
    });

    const toggleSideBarButton = document.querySelector("#toggleSideBar");
    const sideBarHamburgerSpan = document.querySelector("#sideBarHamburger");

    toggleSideBarButton.addEventListener("click", ()=> {
        sideBarHamburgerSpan.classList.toggle("show-sidebar")
    });

    const userMenuLink = document.querySelector("#userMenuLink");
    const userMenuDropdown = document.querySelector("#userMenuDropdown");

    userMenuLink.addEventListener("click", (e) => {
        userMenuDropdown.classList.toggle("hidden");
        e.stopPropagation();
    });

    document.addEventListener("click", () => {
        userMenuDropdown.classList.add("hidden");
    });
});