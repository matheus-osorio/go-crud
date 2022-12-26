package server

import (
	"encoding/json"
	"gocrud/database"
	"net/http"

	"github.com/gorilla/mux"
)

func getMovieLists(writer http.ResponseWriter, request *http.Request) {
	writer = setDefaultHeaders(writer)
	body := database.GetAllMovies()
	json.NewEncoder(writer).Encode(body)
}

func getSingleMovie(writer http.ResponseWriter, request *http.Request) {
	writer = setDefaultHeaders(writer)
	params := mux.Vars(request)
	movie := database.GetSingleMovie(params["id"])
	json.NewEncoder(writer).Encode(movie)
}

func addMovie(writer http.ResponseWriter, request *http.Request) {
	var movie database.Movie
	json.NewDecoder(request.Body).Decode(&movie)
	database.AddMovie(movie)
	json.NewEncoder(writer).Encode(true)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	var movie database.Movie
	params := mux.Vars(request)
	json.NewDecoder(request.Body).Decode(&movie)
	movie.Id = params["id"]
	database.UpdateMovie(movie)
	json.NewEncoder(writer).Encode(true)
}

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	database.DeleteMovie(params["id"])
	json.NewEncoder(writer).Encode(true)
}
