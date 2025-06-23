package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	db "github.com/ismail118/booklending/db/sql"
	"net/http"
)

type createBookRequest struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	ISBN     string `json:"ISBN" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
	Category string `json:"category" binding:"required"`
}

func (server *Server) createBook(c *gin.Context) {
	var req createBookRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBookParams{
		Title:    req.Title,
		Author:   req.Author,
		ISBN:     req.ISBN,
		Quantity: req.Quantity,
		Category: req.Category,
	}

	// TODO: handle proper error such duplicate unique field
	// mysql error ref: https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
	book, err := server.querier.CreateBook(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

type getBookRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getBook(c *gin.Context) {
	var req getBookRequest

	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.querier.GetBook(context.Background(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

type updateBookRequest struct {
	ID       int64  `json:"id" binding:"required,min=1"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	ISBN     string `json:"ISBN" binding:"required"`
	Quantity int64  `json:"quantity" binding:"required,min=1"`
	Category string `json:"category" binding:"required"`
}

func (server *Server) updateBook(c *gin.Context) {
	var req updateBookRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBookParams{
		ID:       req.ID,
		Title:    req.Title,
		Author:   req.Author,
		ISBN:     req.ISBN,
		Quantity: req.Quantity,
		Category: req.Category,
	}

	// TODO: add propher error handling
	book, err := server.querier.UpdateBook(context.Background(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

type deleteBookRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteBook(c *gin.Context) {
	var req deleteBookRequest

	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.querier.DeleteBook(context.Background(), req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, "Success")
}

type getListBooksRequest struct {
	Limit int64 `form:"limit" binding:"required,min=5,max=10"`
	Page  int64 `form:"page" binding:"required,min=1"`
}

func (server *Server) getListBooks(c *gin.Context) {
	var req getListBooksRequest

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetListBookParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	books, err := server.querier.GetListBook(context.Background(), arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, books)
}
