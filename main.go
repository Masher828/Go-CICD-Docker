package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello From Go")
	})

	app.GET("/bye", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Bye From Go")
	})

	err := app.Run(":80")
	if err != nil {
		fmt.Println("Error while starting server", err)
	}
}
