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

	router.Use(CORS())

	router.POST(PREFIX_PATH+"/get-map", server.getMap)
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

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
