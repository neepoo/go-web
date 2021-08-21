package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "github.com/neepoo/go-web/db/sqlc"
)

// Server servers all HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)

	router.POST("/users", server.createUser)

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
