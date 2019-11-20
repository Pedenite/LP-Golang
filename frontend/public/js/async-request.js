let palavra;
let tamanhoLista;
let aprendidas = 0;
let emAndamento = 0;

getRandomPalavras();

function getRandomPalavras(peso) {
    if (palavra) {
        $.ajax({
            url: 'http://localhost:8080/alterar-peso',
            type: 'post',
            dataType: 'html',
            data: {
                peso: peso,
                palavra: palavra.Original
            },
            success: function (data) {
                if (data) {
                    alert('Deck finalizado');
                }
            },
            error: function (e) {
                console.log(e);
            }
        });
    }

    getTamanhoLista();

    $.ajax({
        type: 'GET',
        url: 'http://localhost:8080/palavra',
        dataType: 'json',
        success: function (data) {
            palavra = data
        },
        error: function (e) {
        }
    })
        .then(() => {
            if (cardPosition == 1) {
                // secondRotateCard is a card-animation.js function
                secondRotateCard();
                cardFront.text(palavra.Original);
                setTimeout(function () {
                    cardBack.text(palavra.Traducao);
                }, 1000);
            } else {
                cardFront.text(palavra.Original);
                cardBack.text(palavra.Traducao);
            }
        });
}


function getTamanhoLista() {
    $.ajax({
        type: 'GET',
        url: 'http://localhost:8080/tamanho-lista',
        dataType: 'json',
        success: function (data) {
            if ((tamanhoLista - data) > aprendidas) {
                aprendidas = tamanhoLista - data;
            }
            if (!tamanhoLista || data > tamanhoLista) {
                tamanhoLista = data;
            }
            if ((aprendidas + data) > tamanhoLista) {
                tamanhoLista = aprendidas + data
            }

            qntLoader('totalQnt', tamanhoLista, 'Total', "#ffffff", "#083975", "#083975", 0);
            qntLoader('aprendidasQnt', ((aprendidas*100)/tamanhoLista), 'Aprendidas',"#ffffff", "#00f7ff", "#ffffff", null);
            qntLoader('emAndamentoQnt', 100 - ((aprendidas*100)/tamanhoLista), 'Andamento',"#a6ffce","#00ff72","#a6ffce", null);

            $('.total-value').text(tamanhoLista);
            $('.aprendidas-value').text(aprendidas);
            $('.andamento-value').text(data);
        },
        error: function (e) {
            console.log(e);
        }
    });
}

function addFrase(original, traducao) {
    $.ajax({
        url: 'http://localhost:8080/nova-frase',
        type: 'post',
        dataType: 'html',
        data: {
            original, traducao
        },
        success: function (data) {
        },
        error: function (e) {
            console.log(e);
        }
    });
}
function qntLoader(idContainer, progress, title, colorNum, colorProgress, colorBarrier, progressVal) {
    $(`#${idContainer}`).circularloader({
        backgroundColor: "#083975",//background colour of inner circle
        fontColor: colorNum,//font color of progress text
        fontSize: "25px",//font size of progress text
        radius: 40,//radius of circle
        progressBarBackground: colorBarrier ,//background colour of circular progress Bar
        progressBarColor: colorProgress,//colour of circular progress bar
        progressBarWidth: 15,//progress bar width
        progressPercent: progress,//progress percentage out of 100
        progressValue: progressVal,//diplay this value instead of percentage
        showText: true,//show progress text or not
        title: title,//show header title for the progress bar
    });
}