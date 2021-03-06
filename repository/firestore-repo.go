package repository

import (
	"context"
	"log"

	"github.com/seonicklaus/rest-api-go/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	projectID      string = "rest-api-be2c9"
	collectionName string = "posts"
)

type firestoreRepo struct{}

// Firestore Constructor
func NewFirestoreRepository() PostRepository {
	return &firestoreRepo{}
}

func (*firestoreRepo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to add a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*firestoreRepo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// Delete function :TODO
func (*firestoreRepo) Delete(post *entity.Post) error {
	return nil
}
