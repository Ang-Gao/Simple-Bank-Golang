package api

import (
	db "myproject/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{
		store: store,
	}
	server.setupRouter()
	return server
}
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}

func (server *Server) setupRouter() {
	router := gin.Default()
	server.router = router
	router.POST("/create_account", server.CreateAccount)
}

//H is a shortcut for map[string]interface{}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
