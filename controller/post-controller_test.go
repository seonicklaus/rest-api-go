package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/seonicklaus/rest-api-go/entity"
	"github.com/seonicklaus/rest-api-go/repository"
	"github.com/seonicklaus/rest-api-go/service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	// Create a new HTTP POST request
	var jsonStr = []byte(`{"title":"` + TITLE + `","text":"` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// Assign HTTP Handler Dunction (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add Assertions on the HTTP Status Code and response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; got %v but want %v", status, http.StatusOK)
	}

	// Decode HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// Clean up database
	cleanUp(&post)
}

func TestGetPost(t *testing.T) {
	// Insert New Post
	setup()

	// Create a HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP Handler Dunction (controller AddPost function)
	handler := http.HandlerFunc(postController.GetPosts)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	// Add Assertions on the HTTP Status Code and response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; got %v but want %v", status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// Clean up database
	cleanUp(&posts[0])
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}

	postRepo.Save(&post)
}

func cleanUp(post *entity.Post) {
	postRepo.Delete(post)
}
