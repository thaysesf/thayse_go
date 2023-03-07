package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rotas := mux.NewRouter().StrictSlash(true)

	rotas.HandleFunc("/api/comment/list/1", getAll).Methods("GET")
	rotas.HandleFunc("/api/comment/list/2", getAll2).Methods("GET")
	rotas.HandleFunc("/api/comment/new", create).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, rotas))

}

type Comentario struct {
	Email      string
	Comentario string
	Content_id int
}

var Comentarios1 = []Comentario{}
var Comentarios2 = []Comentario{}

func getAll(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Comentarios1)
}
func getAll2(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Comentarios2)
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var p Comentario

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &p); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, &p)

	if p.Content_id == 1 {
		Comentarios1 = append(Comentarios1, p)
	} else {
		Comentarios2 = append(Comentarios2, p)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(p); err != nil {
		panic(err)
	}
}
