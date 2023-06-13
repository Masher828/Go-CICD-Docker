package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello From Go updated")
	})

	app.GET("/bye", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Bye From Go updated")
	})

	app.GET("/hi", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hi From Go updated")
	})

	app.GET("/hola", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hola From Go updated")
	})

	err := app.Run(":80")
	if err != nil {
		fmt.Println("Error while starting server", err)
	}
}
