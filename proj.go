package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const MAX = 10

type Palavra struct {
	palavraOriginal string
	palavraTraducao string
	anterior        *Palavra
	proxima         *Palavra
}

type ConjuntoPalavras struct {
	tamanho int
	inicio  *Palavra
	fim     *Palavra
}

func (palavras *ConjuntoPalavras) Append(novaPalavra *Palavra) {
	if palavras.tamanho == 0 {
		palavras.inicio = novaPalavra
		palavras.fim = novaPalavra
	} else {
		ultimaPalavra := palavras.fim
		ultimaPalavra.proxima = novaPalavra
		novaPalavra.anterior = ultimaPalavra
		palavras.fim = novaPalavra
	}
	palavras.tamanho++
}

func (palavras *ConjuntoPalavras) Show() {
	contador := 0
	var palavraInicial = palavras.inicio
	for contador < palavras.tamanho {
		fmt.Printf("Palavra %v: %v\n", contador, palavraInicial)
		palavraInicial = palavraInicial.proxima
		contador++
	}
}

func leArquivo(arquivo string) []string {
	var frases []string
	file, err := os.Open(arquivo)
	if err != nil { //tratamento de erro
		fmt.Println(err)
		return nil
	}
	defer file.Close() //vai ser executado apenas ao fim da função por causa do defer
	reader := bufio.NewReader(file)
	for {
		linha, err := reader.ReadString('-')
		linha = strings.TrimSpace(linha) //Para eliminar espacos desnecessarios no inicio e fim de strings
		noDashes := strings.Replace(linha, "-", "", -1)
		frases = append(frases, noDashes)

		linha, err = reader.ReadString('\n')
		linha = strings.TrimSpace(linha) //Para eliminar espacos desnecessarios no inicio e fim de strings

		frases = append(frases, linha)

		if err == io.EOF {
			break
		}

	}
	return frases
}

func main() {

	var qtd int8 = 0
	var userFile string
	var frases []string
	for qtd <= 0 || qtd > MAX {
		fmt.Printf("Digite a quantidade de arquivos a serem lidos: ")
		fmt.Scan(&qtd)
	}
	for qtd > 0 {
		fmt.Printf("Indique o nome do arquivo: ")
		fmt.Scan(&userFile)
		temp := leArquivo("files/" + userFile)
		i := 0
		for range temp {
			frases = append(frases, temp[i])
			i++
		}
		qtd--
	}
	conjuntoP := &ConjuntoPalavras{}
	var p1 string
	var p2 string
	i := 0
	for range frases {
		fmt.Println(frases[i])
		if i%2 == 0 {
			p1 = frases[i]
			p2 = frases[i+1]
			palavra := Palavra{
				palavraOriginal: p1,
				palavraTraducao: p2,
			}
			conjuntoP.Append(&palavra)
		}
		i++
	}
	conjuntoP.Show()

}
