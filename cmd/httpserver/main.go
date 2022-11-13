package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wesleynepo/auth-service-go/internal/core/service/authsrv"
	"github.com/wesleynepo/auth-service-go/internal/core/service/usersrv"
	"github.com/wesleynepo/auth-service-go/internal/handlers/auth"
	"github.com/wesleynepo/auth-service-go/internal/handlers/health"
	"github.com/wesleynepo/auth-service-go/internal/handlers/users"
	"github.com/wesleynepo/auth-service-go/internal/repositories/usersrepo"
	"github.com/wesleynepo/auth-service-go/pkg/database"
	"github.com/wesleynepo/auth-service-go/pkg/hash"
	"github.com/wesleynepo/auth-service-go/pkg/jwt"
)

func main() {
    storage := database.New("postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable&application_name=GOGOS")
    defer storage.Close()

    usersRepository := usersrepo.NewRelational(storage.Get())
    usersService := usersrv.New(usersRepository, hash.New())
    authService := authsrv.New(jwt.New(), usersService)
    authHandler := auth.NewHTTPHandler(authService)
    healthHandler := health.NewHTTPHandler()
    usersHandler := users.NewHTTPHandler(usersService)
    router := gin.Default()

    router.GET("/healthcheck", healthHandler.Get)

    router.POST("/login", authHandler.Login)
    router.POST("/refresh", authHandler.Refresh)
    router.POST("/create", usersHandler.Create)

    router.Run()
}
