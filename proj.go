package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"

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
		//fmt.Printf("Palavra %v: %v\n", contador, palavraInicial)
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
func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("frontend/index.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}
func pgMain(w http.ResponseWriter, r *http.Request) {
	teste, err := r.Cookie("Email")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if teste.Value != "" {
		tpl, _ := template.ParseFiles("frontend/main.html")
		w.WriteHeader(http.StatusOK)
		tpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func getFormulario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostForm.Get("txtEmail")
	fmt.Println("Email logado:", email)
	lembrarDeMim := r.PostForm.Get("lembrar")
	fmt.Println("lembrar De Mim?", lembrarDeMim)
	expirar := time.Now()
	if lembrarDeMim == "on" {
		expirar = time.Now().Add(24 * 60 * 60 * time.Second) //Lembrar por uma dia
	} else {
		expirar = time.Now().Add(120 * time.Second) //Lembrar por 2 minutis
	}
	cookie := http.Cookie{Name: "Email", Value: email, Expires: expirar}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

func main() {

	var frases []string

	files, err := ioutil.ReadDir("files/") //Armazena todos os nomes de arquivos do diretório files no array de strings files
	if err != nil {
		log.Fatal(err)
	}

	qtd := 1
	//fmt.Println("Foram encontrados os seguintes arquivos:")
	for _, f := range files {
		fmt.Println("indice do arquivo:", qtd)
		fmt.Println("Arquivo encontrado:", f.Name(), "\n")
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
	router.HandleFunc("/", index)
	router.HandleFunc("/login", getFormulario)
	router.HandleFunc("/main", pgMain)
	log.Fatal(http.ListenAndServe(":8080", router))
}
