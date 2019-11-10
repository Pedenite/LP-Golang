package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
    "strings"
    
    "encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const MAX = 10

type Palavra struct {
	palavraOriginal string `json:"Original"`
	palavraTraducao string `json:"Traducao"`
	anterior        *Palavra `json:"Anterior"`
	proxima         *Palavra `json:"Proxima"`
}

type PalavraJSON struct {
    Original string `json:"Original"`
    Traducao string `json:"Traducao"`
}
type ArrayPalavras []PalavraJSON

var objPalavras = ArrayPalavras {}

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

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")
}

func getPalavras(w http.ResponseWriter, r *http.Request) {
    allowCORS(w)
	json.NewEncoder(w).Encode(objPalavras)
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
            palavraJSON := PalavraJSON {
                Original: p1,
                Traducao: p2,
            }
            objPalavras = append(objPalavras, palavraJSON)
			conjuntoP.Append(&palavra)
		}
		i++
	}
    conjuntoP.Show()
    
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/palavras", getPalavras).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}