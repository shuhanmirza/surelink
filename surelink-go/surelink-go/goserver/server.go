package goserver

import (
	"github.com/gin-gonic/gin"
	gedis "surelink-go/redisStore"
	db "surelink-go/sqlc"
)

const PREFIX_PATH = "/redirection"

type Server struct {
	store      *db.Store
	router     *gin.Engine
	redisStore *gedis.RedisStore
}

func NewServer(store *db.Store, redisStore *gedis.RedisStore) *Server {
	server := &Server{store: store, redisStore: redisStore}
	router := gin.Default()

	router.GET(PREFIX_PATH+"/get-map", server.getMap)
	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
