package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleynepo/auth-service-go/internal/handlers/auth"
	"github.com/wesleynepo/auth-service-go/internal/handlers/health"
)

func main() {
    authHandler := auth.NewHTTPHandler()
    healthHandler := health.NewHTTPHandler()
    router := gin.Default()

    router.GET("/healthcheck", healthHandler.Get)

    router.POST("/login", authHandler.Login)
    router.POST("/refresh", authHandler.Refresh)

    router.Run()
}
