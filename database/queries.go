package database

import (
	"github.com/google/uuid"
	"github.com/imdario/mergo"
)

var database = map[string]Movie{
	"ba403c24-d18a-4c5f-b968-c78f6d47b509": {
		Id:    "ba403c24-d18a-4c5f-b968-c78f6d47b509",
		Isbn:  "978-3-16-148410-0",
		Title: "An Example Movie",
		Director: Director{
			Person: Person{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		Actors: []Actor{
			{
				Person: Person{
					FirstName: "John",
					LastName:  "Doe",
				},
			},
			{
				Person: Person{
					FirstName: "John2",
					LastName:  "Doe2",
				},
			},
		},
	},
}

func AddMovie(movie Movie) {
	id := uuid.New().String()
	movie.Id = id
	database[id] = movie
}

func GetAllMovies() map[string]Movie {
	return database
}

func GetSingleMovie(id string) Movie {
	return database[id]
}

func DeleteMovie(id string) {
	delete(database, id)
}

func UpdateMovie(movie Movie) {
	srcMovie := GetSingleMovie(movie.Id)
	mergo.Merge(&srcMovie, movie, mergo.WithOverride)
	database[movie.Id] = srcMovie
}
