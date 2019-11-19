// -- Pop Up animation --
const authBtn = $('#auth-btn');
const popUp = $('.pop-up');
const userInput = $('#user-input');
const userLabel = $('#user-label');

if (!localStorage.getItem('email')) {
    popUp.css('display', 'flex');
} else {
    showAllElements();
}

authBtn.on('click', () => {
    localStorage.setItem("email", userInput.val());
    fadePopUp();
});

userInput.on('focus', () => {
    userLabel.css({
        'top': '-1.7rem',
        'left': '.5rem',
        'font-size': '1.4rem'
    })
});

function fadePopUp() {
    popUp.fadeTo("slow", 0, (done) => {
        popUp.css('display', 'none');
        showAllElements();
    });
}

function showAllElements() {
    $('.main__content > *').fadeTo("fast", 1);
    $('.main__sidebar > *').fadeTo("fast", 1);
}