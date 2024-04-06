package api

import (
	"github.com/gin-gonic/gin"
	db "kodboris/db/sqlc"
	"kodboris/util"
)

type Server struct {
	router *gin.Engine
	db     *db.Store
	config util.Config
}

func NewServer(config util.Config, store *db.Store) *Server {
	return &Server{
		db:     store,
		config: config,
	}
}

func (server *Server) setupRouter() {
	router := gin.Default()
	//router.POST("api/v1/accounts", server.createAccount)
	router.GET("/", server.homeProbe)
	router.POST("/member", server.createMember)
	server.router = router

}

// start the HTTP server
func (server *Server) Start(address string, store *db.Store) error {
	server.setupRouter()
	return server.router.Run(address)
}

// error handler
func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
