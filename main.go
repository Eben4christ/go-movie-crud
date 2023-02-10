// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"github.com/gorilla/mux"
// )

// type Movie struct{
// 	ID string `json:"id"`
// 	Isbn string `json:"isbn"`
// 	Title string `json:"title"`
// 	Director *Director `json:"director"`
// }

// type Director struct{
// 	Firstname string `json:"firstname"`
// 	Lastname string `json:"lastname"`
// }

// var movies []Movie

// func getMovies(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(movies)
// }

// func deleteMovie(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {

// 		if item.ID == params["id"]{
// 			movies = append(movies[:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(movies)
// }

// func getMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range movies {
// 		if item.ID == params["id"]{
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// }

// func createMovie(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	var movie Movie
// 	_ = json.NewDecoder(r.Body).Decode(&movie)
// 	movie.ID = strconv.Itoa(rand.Intn(100000000))
// 	movies = append(movies, movie)
// 	json.NewEncoder(w).Encode(movie)
// }

// func updateMovie(w http.ResponseWriter, r *http.Request){
// 	// set json content type
// 	w.Header().Set("Content-Type", "application/json")
// 	// params
// 	params := mux.Vars(r)
// 	// loop over the movies, range
// 	for index, item := range movies{
// 		if item.ID == params["id"]{
// 			movies = append(movies[:index], movies[index+1:]...)
// 			var movie Movie
// 			// delete the movie with the ID that you you have sent
// 			_ = json.NewDecoder(r.Body).Decode(&movie)
// 			movie.ID = params["id"]
// 			// add a new movie - the movie that we send in the postman
// 			movies = append(movies, movie)
// 			json.NewEncoder(w).Encode(movie)
// 			return
// 		}
// 	}
// }

// func main(){
// 	r := mux.NewRouter()

// 	movies = append(movies, Movie{ID: "1",Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
// 	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Eben", Lastname: "Steph"}})

// 	r.HandleFunc("/movies", getMovies).Methods("GET")
// 	r.HandleFunc("/movies{id}", getMovie).Methods("GET")
// 	r.HandleFunc("/movies", createMovie).Methods("POST")
// 	r.HandleFunc("/movies{id}", updateMovie).Methods("PUT")
// 	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

// 	fmt.Print("Starting server at port 8000\n")
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, movie := range movies {
		if movie.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for _, movie := range movies {
		if movie.ID == movieId {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movieId := params["id"]
	for index, item := range movies {
		if item.ID == movieId {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			item.ID = movieId
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "456", Title: "The One", Director: &Director{Firstname: "Ross", Lastname: "Ricky"}})
	movies = append(movies, Movie{ID: "2", Isbn: "236", Title: "Matrix", Director: &Director{Firstname: "John", Lastname: "Wick"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}