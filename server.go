package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequest() {
	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/api/posts", getPosts).Methods("GET")
	router.HandleFunc("/api/posts", addPost).Methods("POST")

	log.Println("Server Listening on Port",port)
	log.Fatalln(http.ListenAndServe(port, router))
}

func main() {
	handleRequest()
}
