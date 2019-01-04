package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GinServer struct {
	*gin.Engine
	Server *http.Server
}

var server *GinServer

func init() {
	server = defaultEngine()
}

func Start() {
	go func() {
		log.Printf("[GIN-debug] Listening and serving HTTP on %s\n", server.Server.Addr)
		if err := server.Server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

}

func Stop() error {
	return server.Server.Shutdown(nil)
}

func defaultEngine() *GinServer {
	eng := gin.Default()
	return &GinServer{
		Engine: eng,
		Server: &http.Server{
			Addr:    ":7788",
			Handler: eng,
		},
	}
}
