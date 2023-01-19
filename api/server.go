package api

import (
	db "myproject/db/sqlc"

	"github.com/gin-gonic/gin"
	binding "github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

//create New server
func NewServer(store *db.SQLStore) *Server {
	server := &Server{
		store: store,
	}
	server.setupRouter()
	return server
}

//server start function
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

//server setup and router path
func (server *Server) setupRouter() {
	router := gin.Default()
	//validate user input
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("supportedCurrency", validCurrency)
	}

	server.router = router

	//--------------------------------------------------------------------
	//Path//
	router.POST("/create_account", server.CreateAccount)
	router.POST("/transfers", server.CreateTransfer)
	//--------------------------------------------------------------------
}

//H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
