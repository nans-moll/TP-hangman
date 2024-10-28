package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type Etudiant struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   string
}

type Classe struct {
	NomClasse       string
	Filiere         string
	Niveau          string
	NombreEtudiants int
	Etudiants       []Etudiant
}

var counter int
var mutex sync.Mutex // Mutex pour protéger le compteur

func promoHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/promo.html"))

	classe := Classe{
		NomClasse:       "B1 Informatique",
		Filiere:         "Informatique",
		Niveau:          "Bachelor 1",
		NombreEtudiants: 4,
		Etudiants: []Etudiant{
			{Nom: "white", Prenom: "walter", Age: 20, Sexe: "M"},
			{Nom: "Smith", Prenom: "Jane", Age: 19, Sexe: "F"},
			{Nom: "Pereira", Prenom: "Alex", Age: 37, Sexe: "M"},
			{Nom: "White", Prenom: "Emily", Age: 22, Sexe: "F"},
		},
	}

	tmpl.Execute(w, classe)
}

func changeHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	viewCount := counter
	mutex.Unlock()

	tmpl := template.Must(template.ParseFiles("./templates/change.html"))

	var message string
	var bgColor string
	if viewCount%2 == 0 {
		message = fmt.Sprintf("Le nombre de vues est pair : %d", viewCount)
		bgColor = "lightblue"
	} else {
		message = fmt.Sprintf("Le nombre de vues est impair : %d", viewCount)
		bgColor = "lightcoral"
	}

	tmpl.Execute(w, map[string]string{
		"Message": message,
		"BgColor": bgColor,
	})
}

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/promo", promoHandler)

	http.HandleFunc("/change", changeHandler)

	fmt.Println("Le serveur est démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
