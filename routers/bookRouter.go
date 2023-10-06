package routers

import (
	"github.com/gin-gonic/gin"
	"simple-rest-api-go/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.AddBook)

	router.PUT("/books/:bookID", controllers.UpdateBook)

	router.GET("/books/:bookID", controllers.GetBookById)

	router.GET("/books", controllers.GetBook)

	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}

func main() {

}
