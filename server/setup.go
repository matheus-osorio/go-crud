package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var app = mux.NewRouter()

func Setup() {
	app.HandleFunc("/movies", getMovieLists).Methods("GET")
	app.HandleFunc("/movies/{id}", getSingleMovie).Methods("GET")
	app.HandleFunc("/movies", addMovie).Methods("POST")
	app.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	app.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Finished server setup")
}

func Run(port string) {
	log.Println("Initializing Server")
	err := http.ListenAndServe(":"+port, app)
	log.Fatal(err)
}
