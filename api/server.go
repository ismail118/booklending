package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ismail118/booklending/db/sql"
	"github.com/ismail118/booklending/token"
)

type Server struct {
	querier db.Querier
	router  *gin.Engine
	paseto  *token.Paseto
}

func NewServer(querier db.Querier, paseto *token.Paseto) *Server {
	server := &Server{
		querier: querier,
		paseto:  paseto,
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

	router.POST("/user/login", server.loginUser)

	authRouth := router.Group("/")
	authRouth.Use(authMiddleware(server.paseto))
	authRouth.POST("/borrowbook", server.createBorrowBook)
	authRouth.POST("/returnbook", server.returnBook)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
