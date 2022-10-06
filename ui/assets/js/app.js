import "../css/app.scss";

document.addEventListener("DOMContentLoaded", () => {
    document.querySelector('#hamburger').addEventListener('click', () => {
        const bars = document.querySelectorAll('.hamburger-bar');
        bars.forEach((bar) => {
            bar.classList.toggle('open');
        })

        document.querySelector('#navigation').classList.toggle('open');
    })
});