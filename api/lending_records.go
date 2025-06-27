package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	db "github.com/ismail118/booklending/db/sql"
	"log"
	"net/http"
	"time"
)

var ErrHaveReachMaxBorrow = errors.New("have reach max borrow")

const numMaxBorrows = 5

type createBorrowBookRequest struct {
	Book     int64 `json:"book" binding:"required,min=1"`
	Borrower int64 `json:"borrower" binding:"required,min=1"`
}

func (server *Server) createBorrowBook(c *gin.Context) {
	var req createBorrowBookRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLendingRecordsParams{
		Book:     req.Book,
		Borrower: req.Borrower,
	}

	// check active landing books
	listLandingRecord, err := server.querier.GetListLendingRecordByBorrower(context.Background(), req.Borrower)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		} else {
			// user have no active landing books
			log.Println("user can land a book cause have no active landing books")
		}
	}

	if len(listLandingRecord) > numMaxBorrows {
		c.JSON(http.StatusForbidden, errorResponse(ErrHaveReachMaxBorrow))
		return
	}

	newId, err := server.querier.CreateLendingRecords(context.Background(), arg)
	if err != nil {
		if db.IsSqlErr(err, db.ErrNumForeignKeyViolation) {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg1 := db.UpdateBookQtyParams{
		Qty: -1,
		ID:  req.Book,
	}
	err = server.querier.UpdateBookQty(context.Background(), arg1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	landingRecord, err := server.querier.GetLendingRecord(context.Background(), newId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, landingRecord)
}

type returnBookRequest struct {
	ID   int64 `json:"id" binding:"required,min=1"`
	Book int64 `json:"book" binding:"required,min=1"`
}

func (server *Server) returnBook(c *gin.Context) {
	var req returnBookRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ReturnBookParams{
		ID:         req.ID,
		ReturnDate: time.Now(),
	}

	err = server.querier.ReturnBook(context.Background(), arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg1 := db.UpdateBookQtyParams{
		Qty: 1,
		ID:  req.Book,
	}

	err = server.querier.UpdateBookQty(context.Background(), arg1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	book, err := server.querier.GetBook(context.Background(), req.Book)
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
