// -- Pop Up animation --
const authBtn = $('#auth-btn');
const popUp = $('.pop-up');
const userInput = $('#user-input');
const userLabel = $('#user-label');

authBtn.on('click', () => {
    popUp.fadeTo("slow", 0, (done) => {
        popUp.css('display', 'none');

        $('.main__content > *').fadeTo("fast", 1);
        $('.main__sidebar > *').fadeTo("fast", 1);
    });
});

userInput.on('focus', () => {
    userLabel.css({
        'top': '-1.7rem',
        'left': '.5rem',
        'font-size': '1.4rem'
    })
});