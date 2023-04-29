package controllers

import (
	"DTSGolang/Kelas2/Assignment8/database"
	"DTSGolang/Kelas2/Assignment8/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthor(ctx *gin.Context) {
	db := database.GetDB()

	type AuthorInput struct {
		Name string `json:"name" binding:"required"`
	}

	var newAuthor AuthorInput

	if err := ctx.ShouldBindJSON(&newAuthor); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	author := models.Author{
		Name: newAuthor.Name,
	}

	err := db.Create(&author).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func GetAuthorById(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()

	author := models.Author{}

	err := db.First(&author, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Author Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, author)
}

func GetAuthors(ctx *gin.Context) {
	db := database.GetDB()

	var authors []models.Author

	err := db.Find(&authors).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Authors Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func UpdateAuthorById(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()

	type AuthorInput struct {
		Name string `json:"name" binding:"required"`
	}

	var updateAuthor AuthorInput
	if err := ctx.ShouldBindJSON(&updateAuthor); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	author := models.Author{}
	authorUpdate := models.Author{
		Name: updateAuthor.Name,
	}

	err := db.Model(&author).Where("id = ?", id).Updates(authorUpdate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Author Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.First(&authorUpdate, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Author Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, authorUpdate)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	type BookInput struct {
		AuthorID  uint   `json:"author_id" binding:"required"`
		Publisher string `json:"publisher" binding:"required"`
		Title     string `json:"title" binding:"required"`
	}

	var newBook BookInput
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{
		AuthorID:  newBook.AuthorID,
		Publisher: newBook.Publisher,
		Title:     newBook.Title,
	}

	err := db.Create(&book).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()

	book := models.Book{}

	err := db.First(&book, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func GetBooks(ctx *gin.Context) {
	db := database.GetDB()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Books Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func UpdateBookById(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()

	type BookInput struct {
		AuthorID  uint   `json:"author_id" binding:"required"`
		Publisher string `json:"publisher" binding:"required"`
		Title     string `json:"title" binding:"required"`
	}

	var updateBook BookInput
	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book := models.Book{}

	bookUpdate := models.Book{
		AuthorID:  updateBook.AuthorID,
		Publisher: updateBook.Publisher,
		Title:     updateBook.Title,
	}

	err := db.Model(&book).Where("id = ?", id).Updates(bookUpdate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, bookUpdate)
}

func DeleteBookByID(ctx *gin.Context) {
	id := ctx.Param("id")

	db := database.GetDB()

	book := models.Book{}

	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Book Not Found")
			return
		}

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book with id %d has been successfully deleted",
	})
}

func GetAuthorWithBook(ctx *gin.Context) {
	db := database.GetDB()

	authors := models.Author{}

	// Nama tabel untuk parameter fungsi Preload() haruslah uppercase untuk awal nama table-nya
	err := db.Preload("Books").Find(&authors).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get author data with books"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"authors": authors,
	})
}
