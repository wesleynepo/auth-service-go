package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {}

func NewHTTPHandler() *HTTPHandler{
    return &HTTPHandler{}
}

func (handler HTTPHandler) Get(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "available",
    })
}
