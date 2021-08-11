package main

import (
	"fmt"
	"net/http"

	"github.com/seonicklaus/rest-api-go/controller"
	router "github.com/seonicklaus/rest-api-go/http"
	"github.com/seonicklaus/rest-api-go/repository"
	"github.com/seonicklaus/rest-api-go/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port = "8000"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server has started...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
