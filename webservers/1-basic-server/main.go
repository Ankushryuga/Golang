package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`		//first latter is capital b/c we r exporting it.
	Isbn string `json:"ispn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
} 


var movies []Movie;	//slice of movie.


func getAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode({});
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// movies = append(movies, Movie{ID:r.id, Isbn: r.isbn, Title: r.title, Director: &Director{}})
}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID:"282882", Isbn:"191919nsnss0", Title:"Movie-one", Director: &Director{FirstName: "ankush", LastName: "raj"}})
	movies = append(movies, Movie{ID:"449949", Isbn:"4994944nfnf4", Title:"Movie-Two", Director: &Director{FirstName: "ravi", LastName: "kaushal"}})
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/createMovie", createMovie).Methods("POST")
	fmt.Print("Server started at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
