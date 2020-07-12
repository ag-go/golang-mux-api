package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/pragmaticreviews/golang-mux-api/routes"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}/", routes.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
