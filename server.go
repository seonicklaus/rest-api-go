package main

import (
	"os"

	"github.com/seonicklaus/rest-api-go/cache"
	"github.com/seonicklaus/rest-api-go/controller"
	router "github.com/seonicklaus/rest-api-go/http"
	"github.com/seonicklaus/rest-api-go/repository"
	"github.com/seonicklaus/rest-api-go/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.DELETE("/posts", postController.DeletePost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
