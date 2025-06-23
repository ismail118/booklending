package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ismail118/booklending/db/sql"
)

type Server struct {
	querier db.Querier
	router  *gin.Engine
}

func NewServer(querier db.Querier) *Server {
	server := &Server{
		querier: querier,
	}

	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/books", server.createBook)
	router.GET("/books/:id", server.getBook)
	router.PUT("/books", server.updateBook)
	router.DELETE("/books/:id", server.getBook)
	router.GET("/books", server.getListBooks)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
