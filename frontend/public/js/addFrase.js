const btnAddFrase = $('#btn-addFrase');
const inputPalavraOriginal = $('input[name="fraseOriginal"]');
const inputPalavraTraducao = $('input[name="fraseTraducao"]');

btnAddFrase.on('click', () => {
    addFrase(inputPalavraOriginal.val(), inputPalavraTraducao.val());
    getTamanhoLista();
});