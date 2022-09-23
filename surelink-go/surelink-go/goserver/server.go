package goserver

import (
	"github.com/gin-gonic/gin"
	gedis "surelink-go/redisStore"
	db "surelink-go/sqlc"
)

const REDIRECTION_PATH_PREFIX = "/redirection"
const CAPCTHA_PATH_PREFIX = "/captcha"

type Server struct {
	store      *db.Store
	router     *gin.Engine
	redisStore *gedis.RedisStore
}

func NewServer(store *db.Store, redisStore *gedis.RedisStore) *Server {
	server := &Server{store: store, redisStore: redisStore}
	router := gin.Default()

	router.Use(CORS())

	router.GET(REDIRECTION_PATH_PREFIX+"/get-map", server.getMap)
	router.POST(REDIRECTION_PATH_PREFIX+"/set-map", server.setMap)

	router.GET(CAPCTHA_PATH_PREFIX+"/get-captcha", server.getCaptcha)

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
