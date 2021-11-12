package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var port = ":8080"
var tpl *template.Template

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func artistsPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "artists.html", data)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}

}

func oneArtistPage(w http.ResponseWriter, r *http.Request) {
	a, err1 := strconv.Atoi(r.URL.Query().Get("id"))
	if err1 != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	if a > len(data) {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}

	err := tpl.ExecuteTemplate(w, "artist_profile.html", data[a-1])
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func randomPage(w http.ResponseWriter, r *http.Request) {
	min := 1
	max := len(data)
	myId := (rand.Intn(max-min) + min)
	err := tpl.ExecuteTemplate(w, "artist_profile.html", data[myId-1])
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func authorPage(w http.ResponseWriter, r *http.Request) {
	authors := "Made by Aleksandr, Peenu, GertIndrek"
	err := tpl.ExecuteTemplate(w, "index.html", authors)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func handleRequests() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/random", randomPage)
	http.HandleFunc("/artists/", oneArtistPage)
	http.HandleFunc("/authors", authorPage)
	http.HandleFunc("/artists", artistsPage)
	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	giveData()
	fmt.Println("Visit for result http://localhost" + port)
	handleRequests()
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	if status == http.StatusNotFound {
		http.Error(w, "Ooooups! 404 Page not found", 404)
	}
	if status == http.StatusInternalServerError {
		http.Error(w, "500 Internal serrrrrver error", 500)
	}
	if status == http.StatusBadRequest {
		http.Error(w, "400 Bad request", 400)
	}
}
