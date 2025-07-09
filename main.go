package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.
		GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		}).
		GET("/sleep", func(ctx *gin.Context) {
			time.Sleep(30 * time.Second)
			ctx.String(http.StatusOK, "enough sleep?")
		}).
		GET("/error", func(ctx *gin.Context) {
			panic("An error happened")
		})
	router.Run("0.0.0.0:8080")
}
