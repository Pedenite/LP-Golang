let palavra;

getRandomPalavras();

function getRandomPalavras(peso) {
    if (palavra) {
        $.ajax({
            url: 'http://localhost:8080/alterar-peso',
            type: 'post',
            dataType: 'html',
            data : { 
                peso: peso,
                palavra: palavra.Original
            },
            success : function(data) {
                if (data) {
                    alert('Deck finalizado');
                }
            },
            error: function(e){
                console.log(e);
            } 
        });
    }
    
    $.ajax({
        type: 'GET',
        url : 'http://localhost:8080/tamanho-lista',
        dataType: 'json',
        success: function(data) {
            $('.total-value').text(data);
        },
        error: function(e){
            console.log(e);
        }
    })

    $.ajax({
        type: 'GET',
        url : 'http://localhost:8080/palavra',
        dataType: 'json',
        success: function(data) {
            palavra = data
        },
        error: function(e){
        }
    })
    .then(() => {
        if (cardPosition == 1) {
            // secondRotateCard is a card-animation.js function
            secondRotateCard();
            cardFront.text(palavra.Original);
            setTimeout(function() {
                cardBack.text(palavra.Traducao);
            }, 1000);
        } else {
            cardFront.text(palavra.Original);
            cardBack.text(palavra.Traducao);
        }
    })
}