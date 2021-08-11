package repository

import "github.com/seonicklaus/rest-api-go/entity"

type PostsRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type report struct {
}
