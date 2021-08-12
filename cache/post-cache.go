package cache

import "github.com/seonicklaus/rest-api-go/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
