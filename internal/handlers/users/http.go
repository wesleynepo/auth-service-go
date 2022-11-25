package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleynepo/auth-service-go/internal/core/ports"
)

type HTTPHandler struct {
    usersService ports.UserService
}

func NewHTTPHandler(usersService ports.UserService) *HTTPHandler{
    return &HTTPHandler{
        usersService: usersService,
    }
}

func (handler *HTTPHandler) Create(c *gin.Context) {
    body := BodyCreateUser{}

    if err := c.BindJSON(&body); err != nil {
        c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
        return
    }

    err := handler.usersService.Create(body.Email, body.Password, body.ConfirmPassword)

    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusCreated)
}
