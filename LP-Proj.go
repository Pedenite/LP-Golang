package main

import (
	"database/sql"
	"fmt"
	"text/template"
	"time"

	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type PalavraJSON struct {
	Original     string `json:"Original"`
	Traducao     string `json:"Traducao"`
	Id           string `json:"Id"`
	EmailCriacao string `json:"Email"`
	Peso         int    `json:"Peso"`
	Data         string `json:"Data"`
}
type ArrayPalavras []PalavraJSON

var objPalavras = ArrayPalavras{}

const dbUsuario string = "root" //Usuario de acesso do Banco de dados
const dbSenha string = "senha"  //senha de acesso do Banco de dados
const dbBanco string = "LP"

func allowCORS(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
}
func setObjPalavrasById(id string) {

	palavraJSON := PalavraJSON{
		Original:     getOriginalById(id),
		Traducao:     getTraducaoById(id),
		Id:           id,
		EmailCriacao: getEmailById(id),
		Peso:         getPesoById(id),
		Data:         getDataById(id),
	}
	objPalavras = append(objPalavras, palavraJSON)

}

func getPalavras(w http.ResponseWriter, r *http.Request) {
	allowCORS(w)
	json.NewEncoder(w).Encode(objPalavras)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("frontend/index.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func pgMain(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("Email")
	setObjPalavrasById(getIdByPeso(w, r))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if email.Value != "" {
		tpl, _ := template.ParseFiles("frontend/main.html")
		w.WriteHeader(http.StatusOK)
		tpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func pgMainGeral(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("Email")
	setObjPalavrasById(getIdByPesoGeral())
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if email.Value != "" {
		tpl, _ := template.ParseFiles("frontend/main.html")
		w.WriteHeader(http.StatusOK)
		tpl.Execute(w, nil)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func btnProximo(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("UPDATE palavras SET peso = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	id := getIdByPeso(w, r)
	peso := getPesoById(id) - 1

	stmtIns.Exec(peso, id)
	fmt.Println(peso, id)
	defer stmtIns.Close()

	http.Redirect(w, r, "/main", http.StatusSeeOther)
}
func btnProximoGeral(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("UPDATE palavras SET peso = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	id := getIdByPesoGeral()
	peso := getPesoById(id) - 1

	stmtIns.Exec(peso, id)
	fmt.Println(peso, id)
	defer stmtIns.Close()

	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

func novasPalavras(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("frontend/indexNewWord.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
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

func getFrases(w http.ResponseWriter, r *http.Request) { //Adiciona novas palavras ao programa
	r.ParseForm()
	email, err := r.Cookie("Email")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	original := r.PostForm.Get("newWord")
	translated := r.PostForm.Get("translate")

	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmtIns, err := db.Prepare("INSERT INTO `LP`.`palavras` (`original`, `tradução`, `emailCriacao`) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	stmtIns.Exec(original, translated, email.Value)
	defer stmtIns.Close()
	//fmt.Println(email.Value, original, translated)

	http.Redirect(w, r, "/NovasPalavras", http.StatusSeeOther)
}
func getOriginalById(id string) string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var original string
	consulta := db.QueryRow("SELECT original FROM palavras WHERE id = " + id)
	consulta.Scan(&original)
	return original
}
func getTraducaoById(id string) string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var original string
	consulta := db.QueryRow("SELECT tradução FROM palavras WHERE id = " + id)
	consulta.Scan(&original)
	return original
}
func getEmailById(id string) string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var original string
	consulta := db.QueryRow("SELECT emailCriacao FROM palavras WHERE id = " + id)
	consulta.Scan(&original)
	return original
}
func getPesoById(id string) int {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var original int
	consulta := db.QueryRow("SELECT peso FROM palavras WHERE id = " + id)
	consulta.Scan(&original)
	return original
}

func getDataById(id string) string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var original string
	consulta := db.QueryRow("SELECT data FROM palavras WHERE id = " + id)
	consulta.Scan(&original)
	return original
}
func getIdByPeso(w http.ResponseWriter, r *http.Request) string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	email, err := r.Cookie("Email")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var id string
	consulta := db.QueryRow("select id from palavras where emailCriacao = '" + email.Value + "'  order by rand(),peso desc")
	consulta.Scan(&id)
	//fmt.Println(id)
	return id
}
func getIdByPesoGeral() string {
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var id string
	consulta := db.QueryRow("select id from palavras  order by rand(),peso desc")
	consulta.Scan(&id)
	fmt.Println(id)
	return id
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	go router.HandleFunc("/palavras", getPalavras).Methods("GET")
	go router.HandleFunc("/", index)
	go router.HandleFunc("/login", getLogin)
	go router.HandleFunc("/main", pgMain)
	go router.HandleFunc("/mainGeral", pgMainGeral)
	go router.HandleFunc("/mainProximo", btnProximo)
	go router.HandleFunc("/sendWord", getFrases)
	go router.HandleFunc("/NovasPalavras", novasPalavras)
	log.Fatal(http.ListenAndServe(":8080", router))
}
