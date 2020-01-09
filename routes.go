package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"./entity"
)

var (
	dao FirestoreDao = FirestoreDao{}
)

// TODO: Implement Clean Architecture Server -> Controller -> Service -> Repository

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(dao.FindAll())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling data"}`))
	}
	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func addPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	post.ID = rand.Int63()
	dao.Save(&post)
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)
}
