package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func main() {

}

func AddBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	maxID := 0
	for _, book := range BookDatas {
		if book.ID > maxID {
			maxID = book.ID
		}
	}

	newBook.ID = maxID + 1

	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created",
		"book":    newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("bookID")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bookID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid ID Format",
		})
		return
	}

	for i, book := range BookDatas {
		if bookID == book.ID {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].ID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with Id %v has been successfully Updated", bookID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with ID %v has been successfully updated", bookID),
	})
}

func GetBookById(ctx *gin.Context) {
	id := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		bookID, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_status":  "Bad Request",
				"error_message": "Invalid ID format",
			})
			return
		}
		if bookID == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func GetBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"book": BookDatas,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		bookID, err := strconv.Atoi(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_status":  "Bad Request",
				"error_message": "Invalid ID format",
			})
			return
		}
		if bookID == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with ID %v Not Found", id),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"Message": fmt.Sprintf("Book with ID %v has been successfully Deleted", id),
	})
}
