package api

import (
	db "github.com/Llala/simplecat/db/sqlc"
	"github.com/Llala/simplecat/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/application", server.createApplication)
	router.GET("/application", server.ListApplication)
	router.DELETE("/application", server.DeleteApplication)
	router.GET("/translation", server.GetTranslation)
	router.PATCH("/translation_unit", server.updateTranslationUnit)
	router.GET("/text_units", server.ListTextUnits)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
