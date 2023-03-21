const headerBurgerButton = document.querySelector('.header__navigation__left__burger');
const headerLeftBar = document.querySelector('.left-pannel__popup');
const popupListItems = document.querySelectorAll('.left-pannel__popup__content__list-item');

headerBurgerButton.addEventListener('click', () => {
    headerLeftBar.classList.toggle('active');
    headerBurgerButton.classList.toggle('active');
})

for (let popupListItem of popupListItems)
    popupListItem.addEventListener('click', () => {
        headerLeftBar.classList.remove('active');
        headerBurgerButton.classList.remove('active');
})