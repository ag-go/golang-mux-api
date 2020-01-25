package main

import (
	"os"

	"gitlab.com/pragmaticreviews/golang-mux-api/controller"
	router "gitlab.com/pragmaticreviews/golang-mux-api/http"
	"gitlab.com/pragmaticreviews/golang-mux-api/repository"
	"gitlab.com/pragmaticreviews/golang-mux-api/service"
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

	httpRouter.SERVE(os.Getenv("PORT"))
}
