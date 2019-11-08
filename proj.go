package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var frases []string

	file, err := os.Open("fileName.txt")
	if err != nil { //tratamento de erro
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha) //Para eliminar espacos desnecessarios

		frases = append(frases, linha)

		if err == io.EOF {
			break
		}

	}
	for i := 0; i < 3; i++ {
		fmt.Println(frases[i])
	}
	file.Close()
}
