package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/neepoo/go-web/db/sqlc"
)

// Server servers all HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{
		store: store,
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start run the HTTP server on specify address
func (server Server) Start(addr string) error {
	return server.router.Run(addr)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
