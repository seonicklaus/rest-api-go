package main

import (
	"os"

	"github.com/seonicklaus/rest-api-go/controller"
	router "github.com/seonicklaus/rest-api-go/http"
	"github.com/seonicklaus/rest-api-go/repository"
	"github.com/seonicklaus/rest-api-go/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.DELETE("/posts", postController.DeletePost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
