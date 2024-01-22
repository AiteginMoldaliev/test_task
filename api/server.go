package api

import (
	db "github.com/AiteginMoldaliev/test-task/db/sqlc"
	"github.com/AiteginMoldaliev/test-task/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config      util.Config
	store       *db.Store
	router      *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config:      config,
		store:       store,
	}

	server.setupRouter()
	
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	api := router.Group("/api/v1")

	// persons api
	api.GET("/persons/:id", server.GetPerson)
	api.POST("/persons", server.CreatPerson)
	api.POST("/persons/query", server.QueryPerson)
	api.POST("/persons/update", server.UpdatePerson)
	api.POST("/persons/list", server.GetPersonsList)
	api.DELETE("/persons/:id", server.DeletePerson)
	
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}