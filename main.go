package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Welcome to GO movies api")
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "45267", Title: "Lord of the Rings", Director: &Director{Firstname: "John", Lastname: "Amaka"}})
	movies = append(movies, Movie{ID: "2", Isbn: "67890", Title: "The Hobbits", Director: &Director{Firstname: "Ade", Lastname: "Dayo"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func deleteMovie(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(movies)
}

func updateMovies(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}

	}
}

func createMovies(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(1000000000)))
	movies = append(movies, movie)
	json.NewEncoder(writer).Encode(movie)
}

func getMovieById(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getMovies(writer http.ResponseWriter, r *http.Request) {
	writer.Header().Set("content-Type", "application/json")
	json.NewEncoder(writer).Encode(movies)
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie
