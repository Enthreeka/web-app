// <--------------------------- LEFT BURGER BAR --------------------------->
const burgerMenuButton = document.querySelector('.header__navigation__left__burger');
const burgerMenuBar = document.querySelector('.header__left__bar');

burgerMenuButton.addEventListener('click', () => {
  burgerMenuBar.classList.toggle('active');
  burgerMenuButton.classList.toggle('active');
})

const taskDescriptionSpan = document.querySelector('.main__view__content__list__item__span__desc');
const taskSubDescriptionSpan = document.querySelector('.main__view__content__list__item__span__sub__desc');

// <--------------------------- POPUP - HOW MUCH TASKS MADE --------------------------->

const howMuchTasksMadeButton = document.querySelector('.header__navigation__right__made');
const howMuchTasksMadePopup = document.querySelector('.header__right__made__popup');

howMuchTasksMadeButton.addEventListener('click', () => {
  howMuchTasksMadeButton.classList.remove('triggered');
  howMuchTasksMadePopup.classList.toggle('active');
  howMuchTasksMadeButton.classList.toggle('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  burgerMenuBar.classList.remove('active');  
  burgerMenuButton.classList.remove('active');
})

document.addEventListener('click', (e) => {
  if (e.target.closest('.header__navigation__right__made') === null && e.target.closest('.header__right__made__popup') === null) {
    howMuchTasksMadeButton.classList.remove('active');
    howMuchTasksMadePopup.classList.remove('active');
  }
})

// <--------------------------- POPUP - HOW MUCH TASKS MADE - INNER BUTTONS --------------------------->

const dailyDoneTasksList = document.querySelector('.header__right__made__popup__list__daily');

const dailyDoneTasksButton = document.querySelector('.header__right__made__popup__buttons__daily');

const howMuchMadeCloseButton = document.querySelector('.header__right__made__popup__button__mobile');

howMuchMadeCloseButton.addEventListener('click', () => {
  howMuchTasksMadePopup.classList.remove('active');
})

// <--------------------------- POPUP - USER ICON --------------------------->

const userIconPopup = document.querySelector('.header__right__user__popup');
const userIconButton = document.querySelector('.header__navigation__right__user');

userIconButton.addEventListener('click', (e) => {
  userIconButton.classList.toggle('active');
  userIconPopup.classList.toggle('active');
  howMuchTasksMadePopup.classList.remove('active');
  howMuchTasksMadeButton.classList.remove('active');
  burgerMenuBar.classList.remove('active');  
  burgerMenuButton.classList.remove('active');
})

document.addEventListener('click', (e) => {
  if (e.target.closest('.header__right__user__popup') === null && e.target.closest('.header__navigation__right__user') === null) {
    userIconPopup.classList.remove('active');
    userIconButton.classList.remove('active');
  }
})

// <--------------------------- POPUP - USER ICON - INNER BUTTONS --------------------------->

const openSettingsButton = document.querySelector('.header__right__user__popup__list__item__settings__link');
const openMainButton = document.querySelector('.header__right__user__popup__list__item__mains__link');
const openThemesButton = document.querySelector('.header__right__user__popup__list__item__themes__link');
const openProductivityButton = document.querySelector('.header__right__user__popup__list__item__productivity__link');
const openNotificationsButton = document.querySelector('.header__right__user__popup__list__item__notifications__link');

const closeSettingsButton = document.querySelector('.header__right__user__settings__popup__content__right__button');

const settingsPopup = document.querySelector('.header__right__user__settings__popup');

openSettingsButton.addEventListener('click', () => {
  settingsPopup.classList.add('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  document.body.style.overflow = 'hidden';

  innerPopupAccountButton.classList.add('active');
  innerPopupMainButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  innerPopupAccount.classList.add('active');
  innerPopupMain.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');
})

openMainButton.addEventListener('click', () => {
  settingsPopup.classList.add('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  document.body.style.overflow = 'hidden';

  innerPopupAccountButton.classList.remove('active');
  innerPopupMainButton.classList.add('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.add('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');
})

openThemesButton.addEventListener('click', () => {
  settingsPopup.classList.add('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  document.body.style.overflow = 'hidden';

  innerPopupAccountButton.classList.remove('active');
  innerPopupMainButton.classList.remove('active');
  innerPopupThemeButton.classList.add('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupTheme.classList.add('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');
})

openProductivityButton.addEventListener('click', () => {
  settingsPopup.classList.add('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  document.body.style.overflow = 'hidden';

  innerPopupAccountButton.classList.remove('active');
  innerPopupMainButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.add('active');
  innerPopupNotificationsButton.classList.remove('active');

  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.add('active');
  innerPopupNotifications.classList.remove('active');
})

openNotificationsButton.addEventListener('click', () => {
  settingsPopup.classList.add('active');
  userIconPopup.classList.remove('active');
  userIconButton.classList.remove('active');
  document.body.style.overflow = 'hidden';

  innerPopupAccountButton.classList.remove('active');
  innerPopupMainButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.add('active');

  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.add('active');

  
})

closeSettingsButton.addEventListener('click', () => {
  settingsPopup.classList.remove('active');
  document.body.style.overflow = 'visible';
})

document.addEventListener('click', (e) => {
  if (e.target.closest('.header__right__user__settings__popup__content') === null 
  && e.target.closest('.header__right__user__popup__list__item__settings__link') === null
  && e.target.closest('.header__right__user__popup__list__item__mains__link') === null
  && e.target.closest('.header__right__user__popup__list__item__productivity__link') === null
  && e.target.closest('.header__right__user__popup__list__item__themes__link') === null
  && e.target.closest('.header__right__user__popup__list__item__notifications__link') === null) {
    settingsPopup.classList.remove('active');
    document.body.style.overflow = 'visible';
  }
})

// <--------------------------- POPUP - USER ICON - SETTINGS - INNER BUTTONS --------------------------->

const innerPopupAccountButton = document.querySelector('.user-account');
const innerPopupAccount = document.querySelector('.header__right__user__settings__popup__content__right__account');

const innerPopupMainButton = document.querySelector('.user-main');
const innerPopupMain = document.querySelector('.header__right__user__settings__popup__content__right__main');

const innerPopupSubscriptionButton = document.querySelector('.user-subscription');
const innerPopupSubscription = document.querySelector('.header__right__user__settings__popup__content__right__subscription');

const innerPopupThemeButton = document.querySelector('.user-theme');
const innerPopupTheme = document.querySelector('.header__right__user__settings__popup__content__right__themes');

const innerPopupProductivityButton = document.querySelector('.user-productivity');
const innerPopupProductivity = document.querySelector('.header__right__user__settings__popup__content__right__productivity');

const innerPopupNotificationsButton = document.querySelector('.user-notifications');
const innerPopupNotifications = document.querySelector('.header__right__user__settings__popup__content__right__notifications');

innerPopupAccountButton.addEventListener('click', () => {
  innerPopupAccount.classList.add('active');
  innerPopupMain.classList.remove('active');
  innerPopupSubscription.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');

  innerPopupAccountButton.classList.add('active')
  innerPopupMainButton.classList.remove('active');
  innerPopupSubscriptionButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

innerPopupMainButton.addEventListener('click', () => {
  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.add('active');
  innerPopupSubscription.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');

  innerPopupAccountButton.classList.remove('active')
  innerPopupMainButton.classList.add('active');
  innerPopupSubscriptionButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

innerPopupSubscriptionButton.addEventListener('click', () => {
  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupSubscription.classList.add('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');

  innerPopupAccountButton.classList.remove('active')
  innerPopupMainButton.classList.remove('active');
  innerPopupSubscriptionButton.classList.add('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

innerPopupThemeButton.addEventListener('click', () => {
  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupSubscription.classList.remove('active');
  innerPopupTheme.classList.add('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.remove('active');

  innerPopupAccountButton.classList.remove('active')
  innerPopupMainButton.classList.remove('active');
  innerPopupSubscriptionButton.classList.remove('active');
  innerPopupThemeButton.classList.add('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.remove('active');

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

innerPopupProductivityButton.addEventListener('click', () => {
  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupSubscription.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.add('active');
  innerPopupNotifications.classList.remove('active');

  innerPopupAccountButton.classList.remove('active')
  innerPopupMainButton.classList.remove('active');
  innerPopupSubscriptionButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.add('active');
  innerPopupNotificationsButton.classList.remove('active');

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

innerPopupNotificationsButton.addEventListener('click', () => {
  innerPopupAccount.classList.remove('active');
  innerPopupMain.classList.remove('active');
  innerPopupSubscription.classList.remove('active');
  innerPopupTheme.classList.remove('active');
  innerPopupProductivity.classList.remove('active');
  innerPopupNotifications.classList.add('active');

  innerPopupAccountButton.classList.remove('active')
  innerPopupMainButton.classList.remove('active');
  innerPopupSubscriptionButton.classList.remove('active');
  innerPopupThemeButton.classList.remove('active');
  innerPopupProductivityButton.classList.remove('active');
  innerPopupNotificationsButton.classList.add('active');

  buttonForNotifications.style.opacity = '1';

  if (document.documentElement.scrollWidth < 520) {
    leftBarPannel.style.left = '-100%';
    mobileMenuButton.style.opacity = '1';
    mobileMenuButton.style.visibility = 'visible';
  }
});

// <--------------------------- TEXTAREA AUTOSIZING --------------------------->

const mainTaskDescription = document.querySelector('.main__view__content__form__task__description');
const mainTaskSubDescription = document.querySelector('.main__view__content__form__task__sub__description');

mainTaskDescription.addEventListener('input', autosize);
mainTaskSubDescription.addEventListener('input', autosize);

function autosize() {
  this.style.height = 'auto';
  let applyNow = this.style.offsetHeight;
  this.style.height = this.scrollHeight - 20 + 'px';
}

// <--------------------------- TASK ADDING / REMVOING  --------------------------->

const mainTaskTemplate = document.querySelector('.main__tasks__template').content;
const mainItemTemplate = mainTaskTemplate.querySelector('.main__view__content__list__item');
const taskList = document.querySelector('.main__view__content__list');
const items = taskList.children;
const newTaskForm = document.querySelector('.main__view__content__form');
const mainTaskSubDescriptionInput = newTaskForm.querySelector('.main__view__content__form__task__sub__description');
const addingButton = document.querySelector('.main__view__content__form__button__add');
const cleaningButton = document.querySelector('.main__view__content__form__button__clean');
const howMushTasksMadeCounter = document.querySelector('.header__navigation__right__made__header__btn__counter');
const howMuchTasksForToday = document.querySelector('.main__view__content__heading');

const onlyImportantTasks = document.querySelector('.left__bar__important__link');
const onlyImportantTasksSpan = document.querySelector('.left__bar__important__link__span');
const allTasks = document.querySelector('.left__bar__all__tasks'); 
const allTasksSpan = document.querySelector('.left__bar__all__tasks__span');
const onlyNonImportantTasks = document.querySelector('.left__bar__non__important__link');
const onlyNonImportantTasksSpan = document.querySelector('.left__bar__non__important__link__span');

const howMuchTasks = document.querySelector('.left__bar__all__tasks__span');
const howMuchImportantTasks = document.querySelector('.left__bar__important__link__span');
const howMuchCommonTasks = document.querySelector('.left__bar__non__important__link__span');

const howMuchTasksMadeInner = document.querySelector('.header__right__made__popup__list__daily__paragraph__span');
const howMuchTasksMessage = document.querySelector('.header__right__made__popup__list__daily__span__hidden');

const todayTasksGraph = document.querySelector('.header__right__made__popup__list__daily__item');

let counter = 0;
let counterForTasks = 0;

let counterForCommonTasks = 0;
let counterForImportantTasks = 0;
let counterForAllTasks = 0;

let count = 0;

let importantForChart = 0;
let simpleForChart = 0;

let howMuchHaveYouDoneToday = () => {
  counter++;
  return counter;
}

let addingChartData = function () {
  myChart.data.datasets[0].data.push(importantForChart, simpleForChart);
  myChart.data.datasets[0].data.reverse();
  myChart.data.datasets[0].data.length = 2;
  myChart.update();
}

let deleteTaskHandler = (item) => {
  let taskListCloseButton = item.querySelector('.main__view__content__list__item__button__done');
  
    taskListCloseButton.addEventListener('click', () => {
      item.remove()
      howMuchHaveYouDoneToday(0)
      howMushTasksMadeCounter.textContent = counter;
      howMuchTasksMadeInner.textContent = counter + '/5';

      if (counter >= 5) {
        howMuchTasksMessage.classList.remove('hidden');
      }

      if (counter === 1) {
        productivityTip.classList.add('active');
        howMuchTasksMadeButton.classList.add('triggered');
      }

      if (counter === 5) {
        var myCanvas = document.createElement('canvas');

        myCanvas.style.position = 'absolute';
        myCanvas.style.top = '0';  
        myCanvas.style.left = '0';
        myCanvas.style.width = '100%';
        myCanvas.style.height = '100vh';
        myCanvas.style.zIndex = '-10';
        document.body.appendChild(myCanvas);
        
        var myConfetti = confetti.create(myCanvas, {
          resize: true,
          useWorker: true
        });
        myConfetti({
          particleCount: 500,
          spread: 1000
        });

        const doneSound = document.querySelector('.done-sound');
      doneSound.play();
      }

      if (counter === 5 && productivityTip.classList.contains('active')) {
        purposeDone.style.top = '105px';
        purposeDone.classList.add('active');
      } else if (counter === 5 && !productivityTip.classList.contains('active')) {
        purposeDone.classList.add('active');
      }

      counterForTasks--;
      counterForTasks > 0 ? howMuchTasksForToday.textContent = 'Задач на сегодня: ' + counterForTasks : howMuchTasksForToday.textContent = 'Все задачи выполнены!';

      counterForAllTasks--;
      counterForCommonTasks > 0 ? counterForCommonTasks = counterForAllTasks - counterForImportantTasks : counterForCommonTasks = 0;
      counterForImportantTasks >= 0 ? counterForImportantTasks = counterForAllTasks - counterForCommonTasks : counterForImportantTasks = 0;

      howMuchCommonTasks.textContent = counterForCommonTasks;
      howMuchImportantTasks.textContent = counterForImportantTasks;
      howMuchTasks.textContent = counterForAllTasks;

      if (count < 200) {
        count += 40;
      }

      item.classList.contains('important') ? importantForChart++ : simpleForChart++;
      addingChartData();

      if (burgerMenuButton.classList.contains('default')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(173, 51, 51) ' + count + 'px, white ' + count +'px, white 200px)';
      }
      if (burgerMenuButton.classList.contains('green')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(141, 173, 51) ' + count + 'px, white ' + count +'px, white 200px)';
      }
      if (burgerMenuButton.classList.contains('emerald')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(51, 173, 151) ' + count + 'px, white ' + count +'px, white 200px)';
      }
      if (burgerMenuButton.classList.contains('purple')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(119, 51, 173 ' + count + 'px, white ' + count +'px, white 200px)';
      }
      if (burgerMenuButton.classList.contains('brown')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(173, 117, 51) ' + count + 'px, white ' + count +'px, white 200px)';
      }
      if (burgerMenuButton.classList.contains('black')) {
        todayTasksGraph.style.background = 'linear-gradient(to right, rgb(44, 42, 40) ' + count + 'px, white ' + count +'px, white 200px)';
      }
    });
};

let isTaskImportant = (item) => {
  let taskIsImportant = item.querySelector('.main__view__content__form__important');

  taskIsImportant.addEventListener('click', () => {
    item.classList.toggle('important');
    taskIsImportant.classList.toggle('important');

    if (homeButton.classList.contains('default')) {
      item.classList.add('default');
      item.classList.remove('green', 'emerald', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('green')) {
      item.classList.add('green');
      item.classList.remove('default', 'emerald', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('emerald')) {
      item.classList.add('emerald');
      item.classList.remove('default', 'green', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('purple')) {
      item.classList.add('purple');
      item.classList.remove('default', 'emerald', 'green', 'brown', 'black');
    }
    if (homeButton.classList.contains('brown')) {
      item.classList.add('brown');
      item.classList.remove('default', 'emerald', 'purple', 'green', 'black');
    }
    if (homeButton.classList.contains('black')) {
      item.classList.add('black');
      item.classList.remove('default', 'emerald', 'purple', 'brown', 'green');
    }

    if (item.classList.contains('important')) {
      counterForImportantTasks++;
    } else if (!item.classList.contains('important') && counterForImportantTasks > 0) {
      counterForImportantTasks--;
    }

    counterForCommonTasks >= 0 ? counterForCommonTasks = counterForAllTasks - counterForImportantTasks : counterForCommonTasks = 0;
    counterForImportantTasks >= 0 ? counterForImportantTasks = counterForAllTasks - counterForCommonTasks : counterForImportantTasks = 0;

    howMuchCommonTasks.textContent = counterForCommonTasks;
    howMuchImportantTasks.textContent = counterForImportantTasks;
    howMuchTasks.textContent = counterForAllTasks;
  })
}

let isTaskEditable = (item) => {
  let taskIsEditable = item.querySelector('.main__view__content__form__edit__button');

  const TaskDescriptionSpan = item.querySelector('.main__view__content__list__item__span__desc');
  const TaskSubDescriptionSpan = item.querySelector('.main__view__content__list__item__span__sub__desc');

  taskIsEditable.addEventListener('click', () => {
    item.classList.toggle('editable');
    taskIsEditable.classList.toggle('editable');

    if (item.classList.contains('editable')) {
      TaskDescriptionSpan.setAttribute('contentEditable', true);
      TaskSubDescriptionSpan.setAttribute('contentEditable', true);
      TaskDescriptionSpan.classList.add('editable');
      TaskSubDescriptionSpan.classList.add('editable');
      TaskDescriptionSpan.focus();
    } else {
      TaskDescriptionSpan.setAttribute('contentEditable', false);
      TaskSubDescriptionSpan.setAttribute('contentEditable', false);
      TaskDescriptionSpan.classList.remove('editable');
      TaskSubDescriptionSpan.classList.remove('editable');
    }

    if (homeButton.classList.contains('default')) {
      item.classList.add('default');
      item.classList.remove('green', 'emerald', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('green')) {
      item.classList.add('green');
      item.classList.remove('default', 'emerald', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('emerald')) {
      item.classList.add('emerald');
      item.classList.remove('default', 'green', 'purple', 'brown', 'black');
    }
    if (homeButton.classList.contains('purple')) {
      item.classList.add('purple');
      item.classList.remove('default', 'emerald', 'green', 'brown', 'black');
    }
    if (homeButton.classList.contains('brown')) {
      item.classList.add('brown');
      item.classList.remove('default', 'emerald', 'purple', 'green', 'black');
    }
    if (homeButton.classList.contains('black')) {
      item.classList.add('black');
      item.classList.remove('default', 'emerald', 'purple', 'brown', 'green');
    }

    if (TaskDescriptionSpan.textContent === '' && TaskSubDescriptionSpan.textContent === '') {
      item.remove();
      counterForTasks--;
      counterForTasks > 0 ? howMuchTasksForToday.textContent = 'Задач на сегодня: ' + counterForTasks : howMuchTasksForToday.textContent = 'Все задачи выполнены! Или их и не было...';
    }
  });
};

const showAllTasks = (item) => {
  allTasks.addEventListener('click', () => {
    item.classList.remove('none');
    allTasks.classList.add('active');
    onlyImportantTasks.classList.remove('active');
    onlyNonImportantTasks.classList.remove('active');
  })
};

const showimportantOnly = (item) => {
  onlyImportantTasks.addEventListener('click', () => {
    allTasks.classList.remove('active');
    onlyImportantTasks.classList.add('active');
    onlyNonImportantTasks.classList.remove('active');

    item.classList.contains('important') ? item.classList.remove('none') : item.classList.add('none');
  })
};

const showNonImportantOnly = (item) => {
  onlyNonImportantTasks.addEventListener('click', () => {
    allTasks.classList.remove('active');
    onlyImportantTasks.classList.remove('active');
    onlyNonImportantTasks.classList.add('active');
    
    item.classList.contains('important') ? item.classList.add('none') : item.classList.remove ('none');
  })
};

for (let i = 0; i < items.length; i++) {
  deleteTaskHandler(items[i]);
};
if(document.querySelectorAll('.main__view__content__form__button__clean').length) {
    cleaningButton.addEventListener('click', (evt) => {
      evt.preventDefault();
      mainTaskDescription.value = '';
      mainTaskSubDescriptionInput.value = '';
    });
}


addingButton.addEventListener('click', (e) => {
  e.preventDefault();

  let mainTaskDescriptionText = mainTaskDescription.value;
  let mainTaskSubDescription = mainTaskSubDescriptionInput.value;

  mainTaskTemplate.querySelector('.main__view__content__list__item__span__desc').textContent = mainTaskDescriptionText;
  mainTaskTemplate.querySelector('.main__view__content__list__item__span__sub__desc').textContent = mainTaskSubDescription;

  let newMainTask = mainItemTemplate.cloneNode(true);

  if (mainTaskDescriptionText !== '' || mainTaskSubDescription !== '') {
  taskList.appendChild(newMainTask);
  counterForTasks++;
  counterForAllTasks++;
  counterForTasks > 0 ? howMuchTasksForToday.textContent = 'Задач на сегодня: ' + counterForTasks : howMuchTasksForToday.textContent = 'Все задачи выполнены! Или их и не было...';
  }

  counterForCommonTasks >= 0 ? counterForCommonTasks = counterForAllTasks - counterForImportantTasks : counterForCommonTasks = 0;
  counterForImportantTasks >= 0 ? counterForImportantTasks = counterForAllTasks - counterForCommonTasks : counterForImportantTasks = 0;

  howMuchCommonTasks.textContent = counterForCommonTasks;
  howMuchImportantTasks.textContent = counterForImportantTasks;
  howMuchTasks.textContent = counterForAllTasks;

//   mainTaskDescription.value = '';
//   mainTaskSubDescriptionInput.value = '';

  showAllTasks(newMainTask);
  showimportantOnly(newMainTask);
  showNonImportantOnly(newMainTask);
  isTaskImportant(newMainTask);
  isTaskEditable(newMainTask);
  deleteTaskHandler(newMainTask);
});

window.onload = function() {
  let mainItems = document.querySelectorAll('.main__view__content__list__item')
  itemsLoaded(mainItems)
  function itemsLoaded(mainItems) {
    let mainTaskDescriptionText = mainTaskDescription.value;
    let mainTaskSubDescription = mainTaskSubDescriptionInput.value;
    if (mainTaskDescriptionText !== '' || mainTaskSubDescription !== '') {
      taskList.appendChild(newMainTask);
      counterForTasks++;
      counterForAllTasks++;
      counterForTasks > 0 ? howMuchTasksForToday.textContent = 'Задач на сегодня: ' + counterForTasks : howMuchTasksForToday.textContent = 'Все задачи выполнены! Или их и не было...';
      }

      counterForCommonTasks >= 0 ? counterForCommonTasks = counterForAllTasks - counterForImportantTasks : counterForCommonTasks = 0;
      counterForImportantTasks >= 0 ? counterForImportantTasks = counterForAllTasks - counterForCommonTasks : counterForImportantTasks = 0;

      howMuchCommonTasks.textContent = counterForCommonTasks;
      howMuchImportantTasks.textContent = counterForImportantTasks;
      howMuchTasks.textContent = counterForAllTasks;

      mainTaskDescription.value = '';
      mainTaskSubDescriptionInput.value = '';
    for (let i = 0; i < mainItems.length; i++) {
      showAllTasks(mainItems[i]);
      showimportantOnly(mainItems[i]);
      showNonImportantOnly(mainItems[i]);
      isTaskImportant(mainItems[i]);
      isTaskEditable(mainItems[i]);
      deleteTaskHandler(mainItems[i]);
    };
  }
let btnsDel = document.querySelectorAll('.main__view__content__list__item__button__done');
btnsDel.forEach(btn => {
  btn.addEventListener("click", (e) => {
    let taskId = btn.closest('form')
    formDel(e,taskId)
});
});
}

const mobileMenuButton = document.querySelector('.header__navigation__left__burger__settings');
const leftBarPannel = document.querySelector('.header__right__user__settings__popup__content__left');
const mobileMenuCloseButton = document.querySelector('.header__right__user__settings__popup__content__right__button__mobile');


const mobileMenuCloseButtonUser = document.querySelector('.header__right__user__popup__button__mobile');
const mobileMenuUser = document.querySelector('.header__right__user__popup');

mobileMenuButton.addEventListener('click', () => {
  leftBarPannel.style.left = '0';
  mobileMenuButton.style.opacity = '0';
  mobileMenuButton.style.visibility = 'hidden';
  buttonForNotifications.style.opacity = '0';
  
});

mobileMenuCloseButton.addEventListener('click', () => {
  leftBarPannel.style.left = '-100%';
  mobileMenuButton.style.opacity = '1';
  mobileMenuButton.style.visibility = 'visible';
  buttonForNotifications.style.opacity = '1';
});

mobileMenuCloseButtonUser.addEventListener('click', () => {
  mobileMenuUser.classList.remove('active');
})

// <--------------------------- CURRENT DATE --------------------------->

const todayDate = new Date();

const todayDay = todayDate.getDate();
const todayMonth = todayDate.getMonth();
const todayYear = todayDate.getFullYear();

const spanForTodayDate = document.querySelector('.main__view__content__heading__span');
let currentMonth;

switch (todayMonth) {
  case 0:
    currentMonth = 'Янв';
    break;
  case 1:
    currentMonth = 'Фев';
    break;
  case 2:
    currentMonth = 'Мар';
    break;
  case 3:
    currentMonth = 'Апр';
    break;
  case 4:
    currentMonth = 'Мая';
    break;
  case 5:
    currentMonth = 'Июня';
    break; 
  case 6:
    currentMonth = 'Июля';
    break;
  case 7:
    currentMonth = 'Авг';
    break;
  case 8:
    currentMonth = 'Сен';
    break;
  case 9:
    currentMonth = 'Окт';
    break;
  case 10:
    currentMonth = 'Ноя';
    break;
  case 11:
    currentMonth = 'Дек';
    break;
}

spanForTodayDate.textContent = todayDay + ' ' + currentMonth + ' ' + todayYear;

// <--------------------------- THEMES --------------------------->

const themeColors = [
  '#AD3333', // красный
  '#8DAD33', // зеленый
  '#33AD97', // изумрудный 
  '#7733AD', // фиолетовый
  '#AD7533', // коричневый
  '#2C2A28', // черный
];

const headingUnderline = document.querySelector('.main__view__content__top');

const defaultThemeButton = document.querySelector('.default__btn');
const greenThemeButton = document.querySelector('.green__btn');
const emeraldThemeButton = document.querySelector('.emerald__btn');
const purpleThemeButton = document.querySelector('.purple__btn');
const brownThemeButton = document.querySelector('.brown__btn');
const blackThemeButton = document.querySelector('.black__btn');

const defaultThemeDiv = document.querySelector('.default__theme');
const greenThemeDiv = document.querySelector('.green__theme'); 
const emeraldThemeDiv = document.querySelector('.emerald__theme');
const purpleThemeDiv = document.querySelector('.purple__theme'); 
const brownThemeDiv = document.querySelector('.brown__theme');
const blackThemeDiv = document.querySelector('.black__theme');

const pageHeader = document.querySelector('.header');
const settingsPopupLowerButton = document.querySelector('.header__right__user__settings__popup__content__right__account__photo__lower__button');
const settingsPopupEmailButton = document.querySelector('.header__right__user__settings__popup__content__right__account__change__email');
const mainViewContent = document.querySelector('.main__view__content__heading');

const homeButton = document.querySelector('.header__navigation__left__home');

defaultThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[0];
  settingsPopupEmailButton.style.backgroundColor = themeColors[0];
  settingsPopupLowerButton.style.backgroundColor = themeColors[0];
  addingButton.style.backgroundColor = themeColors[0];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[0];

  burgerMenuButton.classList.add('default');
  howMuchTasksMadeButton.classList.add('default');
  userIconButton.classList.add('default');
  homeButton.classList.add('default');

  burgerMenuButton.classList.remove('green', 'emerald', 'purple', 'brown', 'black')
  howMuchTasksMadeButton.classList.remove('green', 'emerald', 'purple', 'brown', 'black')
  userIconButton.classList.remove('green', 'emerald', 'purple', 'brown', 'black')
  homeButton.classList.remove('green', 'emerald', 'purple', 'brown', 'black')

  defaultThemeDiv.classList.add('active');
  greenThemeDiv.classList.remove('active');
  emeraldThemeDiv.classList.remove('active');
  purpleThemeDiv.classList.remove('active');
  brownThemeDiv.classList.remove('active');
  blackThemeDiv.classList.remove('active');
});

greenThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[1];
  settingsPopupEmailButton.style.backgroundColor = themeColors[1];
  settingsPopupLowerButton.style.backgroundColor = themeColors[1];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[1];
  addingButton.style.backgroundColor = themeColors[1];

  burgerMenuButton.classList.add('green');
  howMuchTasksMadeButton.classList.add('green');
  userIconButton.classList.add('green');
  homeButton.classList.add('green');

  burgerMenuButton.classList.remove('default', 'emerald', 'purple', 'brown', 'black')
  howMuchTasksMadeButton.classList.remove('default', 'emerald', 'purple', 'brown', 'black')
  userIconButton.classList.remove('default', 'emerald', 'purple', 'brown', 'black')
  homeButton.classList.remove('default', 'emerald', 'purple', 'brown', 'black')

  defaultThemeDiv.classList.remove('active');
  greenThemeDiv.classList.add('active');
  emeraldThemeDiv.classList.remove('active');
  purpleThemeDiv.classList.remove('active');
  brownThemeDiv.classList.remove('active');
  blackThemeDiv.classList.remove('active');
});

emeraldThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[2];
  settingsPopupEmailButton.style.backgroundColor = themeColors[2];
  settingsPopupLowerButton.style.backgroundColor = themeColors[2];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[2];
  addingButton.style.backgroundColor = themeColors[2];

  burgerMenuButton.classList.add('emerald');
  howMuchTasksMadeButton.classList.add('emerald');
  userIconButton.classList.add('emerald');
  homeButton.classList.add('emerald');

  burgerMenuButton.classList.remove('green', 'default', 'purple', 'brown', 'black')
  howMuchTasksMadeButton.classList.remove('green', 'default', 'purple', 'brown', 'black')
  userIconButton.classList.remove('green', 'default', 'purple', 'brown', 'black')
  homeButton.classList.remove('green', 'default', 'purple', 'brown', 'black')

  defaultThemeDiv.classList.remove('active');
  greenThemeDiv.classList.remove('active');
  emeraldThemeDiv.classList.add('active');
  purpleThemeDiv.classList.remove('active');
  brownThemeDiv.classList.remove('active');
  blackThemeDiv.classList.remove('active');
});

purpleThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[3];
  settingsPopupEmailButton.style.backgroundColor = themeColors[3];
  settingsPopupLowerButton.style.backgroundColor = themeColors[3];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[3];
  addingButton.style.backgroundColor = themeColors[3];

  burgerMenuButton.classList.add('purple');
  howMuchTasksMadeButton.classList.add('purple');
  userIconButton.classList.add('purple');
  homeButton.classList.add('purple');

  burgerMenuButton.classList.remove('green', 'emerald', 'default', 'brown', 'black')
  howMuchTasksMadeButton.classList.remove('green', 'emerald', 'default', 'brown', 'black')
  userIconButton.classList.remove('green', 'emerald', 'default', 'brown', 'black')
  homeButton.classList.remove('green', 'emerald', 'default', 'brown', 'black')

  defaultThemeDiv.classList.remove('active');
  greenThemeDiv.classList.remove('active');
  emeraldThemeDiv.classList.remove('active');
  purpleThemeDiv.classList.add('active');
  brownThemeDiv.classList.remove('active');
  blackThemeDiv.classList.remove('active');
});

brownThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[4];
  settingsPopupEmailButton.style.backgroundColor = themeColors[4];
  settingsPopupLowerButton.style.backgroundColor = themeColors[4];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[4];
  addingButton.style.backgroundColor = themeColors[4];

  burgerMenuButton.classList.add('brown');
  howMuchTasksMadeButton.classList.add('brown');
  userIconButton.classList.add('brown');
  homeButton.classList.add('brown');

  burgerMenuButton.classList.remove('green', 'emerald', 'purple', 'default', 'black')
  howMuchTasksMadeButton.classList.remove('green', 'emerald', 'purple', 'default', 'black')
  userIconButton.classList.remove('green', 'emerald', 'purple', 'default', 'black')
  homeButton.classList.remove('green', 'emerald', 'purple', 'default', 'black')

  defaultThemeDiv.classList.remove('active');
  greenThemeDiv.classList.remove('active');
  emeraldThemeDiv.classList.remove('active');
  purpleThemeDiv.classList.remove('active');
  brownThemeDiv.classList.add('active');
  blackThemeDiv.classList.remove('active');
});

blackThemeButton.addEventListener('click', () => {
  pageHeader.style.backgroundColor = themeColors[5];
  settingsPopupEmailButton.style.backgroundColor = themeColors[5];
  settingsPopupLowerButton.style.backgroundColor = themeColors[5];
  headingUnderline.style.borderBottom = '1px solid' + themeColors[5];
  addingButton.style.backgroundColor = themeColors[5];

  burgerMenuButton.classList.add('black');
  howMuchTasksMadeButton.classList.add('black');
  userIconButton.classList.add('black');
  homeButton.classList.add('black');

  burgerMenuButton.classList.remove('green', 'emerald', 'purple', 'brown', 'default')
  howMuchTasksMadeButton.classList.remove('green', 'emerald', 'purple', 'brown', 'default')
  userIconButton.classList.remove('green', 'emerald', 'purple', 'brown', 'default')
  homeButton.classList.remove('green', 'emerald', 'purple', 'brown', 'default')

  defaultThemeDiv.classList.remove('active');
  greenThemeDiv.classList.remove('active');
  emeraldThemeDiv.classList.remove('active');
  purpleThemeDiv.classList.remove('active');
  brownThemeDiv.classList.remove('active');
  blackThemeDiv.classList.add('active');
});

// <--------------------------- DATE TO DEVIDE TASKS --------------------------->

const newDayNewTasks = new Date();

let currentTime = newDayNewTasks.toLocaleTimeString();
let currentDay = todayDate.getDay();

const taskListDateTemplate = document.querySelector('.main__tasks__template__date').content;
const taskListDateChild = taskListDateTemplate.querySelector('.main__view__content__list__item__date');
const taskListDateParagraph = taskListDateTemplate.querySelector('.main__view__content__list__item__paragraph__date');

const weekDays = [
  'Воскресенье',
  'Понедельник',
  'Вторник',
  'Среда',
  'Четверг',
  'Пятница',
  'Суббота',
]

if (currentTime === '00:00:00' && taskList.length) {
  taskList.appendChild(taskListDateChild)
  taskListDateParagraph.textContent = todayDate.getDate() + ' · ' + currentMonth + ' · ' + weekDays[currentDay];
  howMushTasksMadeCounter.textContent = '0';
  counter = 0;
  count = 0;
}

// <------------------------ PRODUCTIVITY GRAPH -------------------------->

const ctx = document.getElementById('myChart');

const myChart = new Chart(ctx, {
    type: 'doughnut',
    data: {
      labels: ['Неважные задачи', 'Важные задачи'],
      datasets: [{

        data: [],
        borderWidth: 1,
        backgroundColor: [
          'rgb(173, 51, 51)',
          'rgba(173, 51, 51, 0.4)',
        ],
      }]
    },
    options: {
      plugins: {
        legend: {
          display: false
        }
      },
      scales: {
        y: {
          beginAtZero: true
        }
      },
    }
  });


// <----------------------------- TIPS ------------------------------->

const productivityTipButton = document.querySelector('.productivity-tip__close-button');
const productivityTip = document.querySelector('.productivity-tip');

productivityTipButton.addEventListener('click', () => {
  productivityTip.classList.remove('active');
  purposeDone.style.top = '50px';
  howMuchTasksMadeButton.classList.remove('triggered');
})

const purposeDoneButton = document.querySelector('.purpose-done__close-button');
const purposeDone = document.querySelector('.purpose-done');

purposeDoneButton.addEventListener('click', () => {
  purposeDone.classList.remove('active');
})

// <---------------------------- NOTIFICATIONS -------------------------->

const notificationsForm = document.querySelector('.header__right__user__settings__popup__content__right__notifications__form');
const buttonForNotifications = document.querySelector('.header__right__user__settings__popup__content__right__notofications__button');
const notificationsBall = document.querySelector('.button-ball');
const firstNotificationsInput = document.querySelector('.notifications-first');
const secondNotificationsInput = document.querySelector('.notifications-second');
const notificationsCondition = document.querySelector('.is-notifications-on');

notificationsForm.addEventListener('submit', () => {
    buttonForNotifications.classList.toggle('active');
    notificationsBall.classList.toggle('active');

    if (buttonForNotifications.classList.contains('active')) {
      notificationsCondition.textContent = 'Уведомления включены';
    } else {
      notificationsCondition.textContent = 'Уведомления выключены';
    }
})
addingButton.onclick = function(e) {
    formSend(e,newTaskForm)
}
async function formSend(e,form) { //сама функция
    e.preventDefault();
    let formData = new FormData(form)
      form.classList.add('_sending');
        let response = await fetch('/dashboard/add' ,{
            method: 'POST',
            body: formData, // тут по идеи должны быть инпуты
        });
          mainTaskDescription.value = '';
          mainTaskSubDescriptionInput.value = '';
}
async function formDel(e,form) { //сама функция
    e.preventDefault();
    let formData = new FormData(form)
    console.log(form)
      form.classList.add('_sending');
        let response = await fetch('/dashboard/delete' ,{
            method: 'DELETE',
            body: formData, // тут по идеи должны быть инпуты
        });
}

const imageLoad = document.querySelector(".load-image-file")

imageLoad.addEventListener('input', function(e){
let form = imageLoad.closest("form")
            formImage(form)
})



async function formImage(form) { //сама функция
    let formData = new FormData(form)
    console.log(form)

      form.classList.add('_sending');
        let response = await fetch('/dashboard/image' ,{
            method: 'POST',
            body: formData, // тут по идеи должны быть инпуты
        });
}

 const saveBtn = document.querySelector('.save-name');

  saveBtn.addEventListener('click', function() {
    let from = saveBtn.closest('form')
    let input = from.querySelector('input')
    const value = input.value;

        console.log(value)
        fetch('/dashboard/update/name', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({value})
        });

  });

const leveaBtn = document.querySelector(".header__right__user__popup__list__item__sign-out__link")

leveaBtn.onclick = function(e){
let form = leveaBtn.closest("form")
console.log(form)
formLeave(form)
}


async function formLeave(form) { //сама функция
    let formData = new FormData(form)
    console.log(form)
      form.classList.add('_sending');
        let response = await fetch('/dashboard/leave' ,{
            method: 'POST',
            body: formData, // тут по идеи должны быть инпуты
        });
}

const editBtn = document.querySelector(".main__view__content__form__edit__button")

editBtn.onclick = function(e){
let form = editBtn.closest("form")
console.log(form)
formEdit(form)
}

async function formEdit(form) { //сама функция
    let formData = new FormData(form)
    console.log(form)
      form.classList.add('_sending');
        let response = await fetch('/dashboard/edit' ,{
            method: 'PUT',
            body: formData, // тут по идеи должны быть инпуты
        });
}
