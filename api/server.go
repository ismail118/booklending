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

	router.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Ok"})
	})
	router.POST("/books", server.createBook)
	router.GET("/books/:id", server.getBook)
	router.PUT("/books", server.updateBook)
	router.DELETE("/books/:id", server.deleteBook)
	router.GET("/books", server.getListBooks)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
