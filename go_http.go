package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)


type Page struct {

	Title string

	Body  []byte

}

func (p *Page) save() error {

	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)

}


func loadPage(title string) (*Page, error) {

	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}


// Acceuil : handler sur :1111/ ooo



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Fluubi's shitty website (Covide 666)</h1>")
fmt.Fprintf(w, "Hello, bienvenu à la racine, ceci est un test de serveur web en GoLang! ")
}


///Cette fonction répond à tous les appels avec voir dans le nom 
func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/voir/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}


func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/voir/", viewHandler)
    log.Fatal(http.ListenAndServe(":1111", nil))
}
