package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"

	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

type PalavraJSON struct {
	Original string `json:"Original"`
	Traducao string `json:"Traducao"`
}
type ArrayPalavras []PalavraJSON

var objPalavras = ArrayPalavras{}

var conjuntoP = &ConjuntoPalavras{}

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

func (palavras *ConjuntoPalavras) ShowAndUpdate() {
	objPalavras = ArrayPalavras{}
	contador := 0
	var palavraInicial = palavras.inicio
	for contador < palavras.tamanho {
		fmt.Printf("Palavra %v: %v\n", contador, palavraInicial)
		palavraJSON := PalavraJSON{
			Original: palavraInicial.palavraOriginal,
			Traducao: palavraInicial.palavraTraducao,
		}
		objPalavras = append(objPalavras, palavraJSON)
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
		linha = strings.TrimSpace(linha)                //Para eliminar espacos desnecessarios no inicio e fim de strings
		noDashes := strings.Replace(linha, "-", "", -1) //Elimina o caractere '-'

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

func getPalavra(w http.ResponseWriter, r *http.Request) {
	allowCORS(w)
	randomNumber := rand.Intn(conjuntoP.tamanho)
	json.NewEncoder(w).Encode(objPalavras[randomNumber])
}

func main() {

	var frases []string

	files, err := ioutil.ReadDir("files/") //Armazena todos os nomes de arquivos do diretório files no array de strings files
	if err != nil {
		log.Fatal(err)
	}

	qtd := 1
	fmt.Println("Foram encontrados os seguintes arquivos:")
	for _, f := range files {
		fmt.Printf("%d - ", qtd)
		fmt.Println(f.Name())
		qtd++
	}

	if qtd == 1 { //caso nenhum arquivo seja encontrado, sai com a msg de erro
		panic("Nenhum arquivo encontrado")
	}

	for _, f := range files { //lê todos os arquivos do diretório files

		temp := leArquivo("files/" + f.Name())
		i := 0
		for range temp {
			frases = append(frases, temp[i])
			i++
		}

	}
	
	var p1 string
	var p2 string
	i := 0
	for range frases {
		//fmt.Println(frases[i])
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
	conjuntoP.ShowAndUpdate()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/palavras", getPalavras).Methods("GET")
	router.HandleFunc("/palavra", getPalavra).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
