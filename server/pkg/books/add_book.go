package books

import (
	"net/http"

	"github.com/claranceliberi/ampersand/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddBookRequestBody struct {
    Title       string `json:"title"`
    Author      string `json:"author"`
    Description string `json:"description"`
}

func (h handler) AddBook(c *gin.Context) {
    body := AddBookRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var book models.Book

    book.Title = body.Title
    book.Author = body.Author
    book.Description = body.Description

    if result := h.DB.Create(&book); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &book)
}