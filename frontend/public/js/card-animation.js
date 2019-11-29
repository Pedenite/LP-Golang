let cardPosition = 0;
const cardFront = $('#card-front');
const cardBack = $('#card-back');
const btnVerTraducao = $('#btn-verTraducao');

btnVerTraducao.on('click', () => {
    if (cardPosition == 0) {
        firstRotateCard();
    } else {
        secondRotateCard();
    }
});

function firstRotateCard() {
    $('.card--front').css('transform', 'rotateY(-180deg)');
    $('.card--back').css('transform', 'rotateY(0)');
    cardPosition = 1;
}

function secondRotateCard() {
    $('.card--front').css('transform', 'rotateY(0)');
    $('.card--back').css('transform', 'rotateY(180deg)');
    cardPosition = 0;
}