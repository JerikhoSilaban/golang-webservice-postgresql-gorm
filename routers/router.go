package routers

import (
	"DTSGolang/Kelas2/Assignment8/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/author", controllers.CreateAuthor)
	router.POST("/book", controllers.CreateBook)

	router.GET("/authors", controllers.GetAuthors)
	router.GET("/books", controllers.GetBooks)

	router.GET("/author/:id", controllers.GetAuthorById)
	router.GET("/book/:id", controllers.GetBookById)

	router.PUT("/author/:id", controllers.UpdateAuthorById)
	router.PUT("/book/:id", controllers.UpdateBookById)

	router.DELETE("/book/:id", controllers.DeleteBookByID)

	router.GET("/authorandbook", controllers.GetAuthorWithBook)

	return router
}
