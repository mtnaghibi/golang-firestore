package main

import (
	"encoding/json"
	"net/http"
	"github.com/mrnaghibi/firebase/repository"
	"github.com/mrnaghibi/firebase/entity"
	"math/rand"
)



var (
repo repository.PostRepository = repository.NewPostRepository() 
)


func getPosts(response http.ResponseWriter , request *http.Request)  {
	response.Header().Set("Content-Type","application/json")
	posts , err := repo.FindAll()
	if err !=nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error Getting the Posts"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func addPost(response http.ResponseWriter , request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var post entity.Post
	err:= json.NewDecoder(request.Body).Decode(&post)
	if err != nil{
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Error Unmarshalling Data"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}