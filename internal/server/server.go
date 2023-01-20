package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {

	// 
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// start server
	go Start(server)
	// go func() {
	// 	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalln(err)
	// 	}
	// }()

	// graceful shutdown
	GracefulShutdown(server)

}

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// api route
	api := router.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// serve static files
	router.Use(static.Serve("/", static.LocalFile("./build", true)))
	router.NoRoute(func(c *gin.Context) { // fallback
		c.File("./build/index.html")
	})
	return router

}

func Start(server *http.Server) {
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}


func Shutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
	log.Println("Server exiting")
}