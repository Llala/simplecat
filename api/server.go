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
	router.GET("/translation", server.GetTranslation)
	router.PATCH("/translation_unit", server.updateTranslationUnit)
	router.GET("/text_units", server.ListTextUnits)
	// router.POST("/users/login", server.loginUser)
	// router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.ListAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
