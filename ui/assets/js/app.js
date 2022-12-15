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
    document.querySelector('#hamburger').addEventListener('click', handleHamburgerClicked);
    document.querySelector('#backdrop').addEventListener('click', handleHamburgerClicked);

    window.addEventListener('resize', (e) => {
        const navigation = document.querySelector('#navigation')
        const backdrop = document.querySelector('#backdrop')
        const bars = document.querySelectorAll('.hamburger-bar');

        if (e.target.innerWidth >= 700 && navigation.classList.contains('open')) {
            navigation.classList.remove('open');
            backdrop.classList.remove('open');
            bars.forEach((bar) => {
                bar.classList.remove('open');
            })
        }
        // document.querySelector('#navigation').classList.toggle('open');
        // document.querySelector('#backdrop').classList.toggle('open');
    })
});