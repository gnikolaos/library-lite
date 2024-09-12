(() => {
    const homeMenu = document.getElementById("homeMenu");
    const openHomeMenu = document.getElementById("openHomeMenu");
    const closeHomeMenu = document.getElementById("closeHomeMenu");

    function toggleHomeMenu() {
        if (homeMenu.style.display === "block") {
            homeMenu.style.display = "none";
            return;
        }

        homeMenu.style.display = "block";
        closeHomeMenu.addEventListener("click", toggleHomeMenu, { once: true });
    }

    openHomeMenu.addEventListener("click", toggleHomeMenu);
})();
