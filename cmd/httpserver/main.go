package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wesleynepo/auth-service-go/internal/core/service/authsrv"
	"github.com/wesleynepo/auth-service-go/internal/handlers/auth"
	"github.com/wesleynepo/auth-service-go/internal/handlers/health"
	"github.com/wesleynepo/auth-service-go/pkg/jwt"
)

func main() {
    authService := authsrv.New(jwt.New())
    authHandler := auth.NewHTTPHandler(authService)
    healthHandler := health.NewHTTPHandler()
    router := gin.Default()

    router.GET("/healthcheck", healthHandler.Get)

    router.POST("/login", authHandler.Login)
    router.POST("/refresh", authHandler.Refresh)

    router.Run()
}
