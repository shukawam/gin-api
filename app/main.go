package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

type InternalServerError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

func main() {
	router := gin.Default()
	router.
		GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, Message{Message: "pong"})
		}).
		GET("/sleep", func(ctx *gin.Context) {
			time.Sleep(45 * time.Second)
			ctx.JSON(http.StatusOK, Message{Message: "enough sleep?"})
		}).
		GET("/error", func(ctx *gin.Context) {
			ctx.JSON(http.StatusInternalServerError, InternalServerError{
				Message: "internal server error",
				Details: "An error happened",
			})
		})
	router.Run(":8000")
}
