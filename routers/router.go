package routers

import (
	"DTSGolang/Kelas2/Assignment8/controllers"

	"github.com/gin-gonic/gin"

	_ "DTSGolang/Kelas2/Assignment8/docs"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a simple service for managing authors and books
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host 127.0.0.1:8000
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Post A Author
	router.POST("/author", controllers.CreateAuthor)

	// Post A Book
	router.POST("/book", controllers.CreateBook)

	// Read All Authors
	router.GET("/authors", controllers.GetAuthors)

	// Read All Books
	router.GET("/books", controllers.GetBooks)

	// Read A Author
	router.GET("/author/:id", controllers.GetAuthorById)

	// Read A Book
	router.GET("/book/:id", controllers.GetBookById)

	// Update A Author
	router.PUT("/author/:id", controllers.UpdateAuthorById)

	// Update A Book
	router.PUT("/book/:id", controllers.UpdateBookById)

	// Delete A Book
	router.DELETE("/book/:id", controllers.DeleteBookByID)

	// Read the authors and related books
	router.GET("/authorandbook", controllers.GetAuthorWithBook)

	return router
}
