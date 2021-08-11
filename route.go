package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/seonicklaus/rest-api-go/entity"
	"github.com/seonicklaus/rest-api-go/repository"
)

var (
	repo repository.PostRepository = repository.NewPostsRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
