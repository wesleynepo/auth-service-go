package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pascaldekloe/jwt"
)

type HTTPHandler struct {}

func NewHTTPHandler() *HTTPHandler{
    return &HTTPHandler{}
}

func (handler *HTTPHandler) Login(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "token": generateToken(),
        "refresh": generateRefreshToken(),
    })
}

func (handler *HTTPHandler) Refresh(c *gin.Context) {
    body := RefreshRequest{}

    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "aqui": "debug"})
        return
    }

    response, err := refreshToken(body)

    if (err != nil) {
        c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

func refreshToken(request RefreshRequest) (*RefreshResponse, error) {
    claimsRefresh, err := jwt.HMACCheck([]byte(request.Refresh), []byte("test"))

    if (err != nil) {
        return nil, err
    }

    claimsToken, err := jwt.HMACCheck([]byte(request.Token), []byte("test"))

    if (err != nil) {
        return nil, err
    }

    if (claimsRefresh.Valid(time.Time{})) {
        return nil, errors.New("Refresh token expired")
    }

    token, err := createToken(claimsToken.Set["id"].(float64))

    if (err != nil) {
        return nil, err
    }

    return createRefreshResponse(token, request.Refresh), nil
}

func createRefreshResponse(token, refreshToken string) *RefreshResponse {
    return &RefreshResponse{Token: token, Refresh: refreshToken}
}

func createToken(userId float64) (string, error) {
    var claims jwt.Claims

    claims.Set = map[string]interface{}{"id": userId}
    claims.Issued  = jwt.NewNumericTime(time.Now().Round(time.Second))
    claims.Expires = jwt.NewNumericTime(time.Now().Add(time.Hour * 4))

    token, err := claims.HMACSign("HS256", []byte("test"))

    if (err != nil) {
        return "", err
    }

    return string(token), nil
}

func generateToken() (string) {
    token, err := createToken(123)

    if (err != nil) {
        log.Printf("error ocurred %v", err)
    }

    return token
}

func generateRefreshToken() (string) {
    var claims jwt.Claims

    claims.Set = map[string]interface{}{"id": 123}
    claims.Issued  = jwt.NewNumericTime(time.Now().Round(time.Second))
    claims.Expires = jwt.NewNumericTime(time.Now().Add(time.Hour * 24))

    token, err := claims.HMACSign("HS256", []byte("test"))

    if (err != nil) {
        log.Printf("error ocurred %v", err)
    }

    return string(token)
}

type RefreshRequest struct {
    Refresh string `json:"refresh"`
    Token string `json:"token"`
}

type RefreshResponse = RefreshRequest
