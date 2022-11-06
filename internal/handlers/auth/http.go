package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleynepo/auth-service-go/internal/core/ports"
)

type HTTPHandler struct {
    authService ports.AuthService
}

func NewHTTPHandler(authService ports.AuthService) *HTTPHandler{
    return &HTTPHandler{
        authService: authService,
    }
}

func (handler *HTTPHandler) Login(c *gin.Context) {
    body := BodyLogin{}
    c.BindJSON(&body)

    auth, err := handler.authService.Login(body.Email, body.Password)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, auth)
}

func (handler *HTTPHandler) Refresh(c *gin.Context) {
    body := BodyRefresh{}
    c.BindJSON(&body)

    auth, err := handler.authService.Refresh(body.Refresh)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, auth)
}
