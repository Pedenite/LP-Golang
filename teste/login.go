/*package main

// ARQUIVO QUE PEGA UM VALOR DE UM FORMULARIO E REDIRECIONA A PAGINA DEPOIS DE PEGAR
//ESTA COMENTADO PQ JA FOI IMPLEMENTADO NO PROJ.GO
import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Ouvindo porta 8080")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/login", getFormulario)
	router.HandleFunc("/main", pgMain)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("frontend/index.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}
func pgMain(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("frontend/main.html")

	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, nil)
}

func getFormulario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("txtEmail")
	fmt.Println(email)
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

var wg sync.WaitGroup////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	canal := make(chan string)
	for _, f := range files {
		go func(frases *[]string, wg *sync.WaitGroup, canal chan string) {
			wg.Add(len(files))
			canal <- "."
			defer wg.Done()
			temp := leArquivo(dir + f.Name())
			i := 0
			for range temp {
				*frases = append(*frases, temp[i])
				i++
			}
			wg.Wait()
		}(&frases, &wg, canal)
	}
	wg.Wait()
	fmt.Println(<-canal)
*/
