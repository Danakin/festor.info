import "../css/app.scss";

const handleHamburgerClicked = () => {
    const bars = document.querySelectorAll('.hamburger-bar');
    bars.forEach((bar) => {
        bar.classList.toggle('open');
    })

    document.querySelector('#navigation').classList.toggle('open');
    document.querySelector('#backdrop').classList.toggle('open');
}

document.addEventListener("DOMContentLoaded", () => {
    document.querySelector('#hamburger').addEventListener('click', handleHamburgerClicked)
    document.querySelector('#backdrop').addEventListener('click', handleHamburgerClicked)
});