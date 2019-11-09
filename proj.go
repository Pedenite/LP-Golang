package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const MAX = 10

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

	i := 0
	for range frases {
		fmt.Println(frases[i])
		i++
	}
}
