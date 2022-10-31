package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleynepo/auth-service-go/internal/handlers/auth"
)

func main() {
    authHandler := auth.NewHTTPHandler()
    router := gin.Default()

    router.GET("/ping", func (c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    router.POST("/login", authHandler.Login)
    router.POST("/refresh", authHandler.Refresh)

    router.Run()
}
