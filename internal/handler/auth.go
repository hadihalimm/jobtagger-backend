package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	provider := c.Param("provider")
	fmt.Println(provider)
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	h.service.SignIn(c.Writer, c.Request)
}

func (h *AuthHandler) AuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()
	user, err := h.service.AuthCallback(c.Writer, c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := h.service.GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := h.service.GenerateRefreshToken(c.Request, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("refresh_token", refreshToken, int(time.Hour*24*14), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (h *AuthHandler) RotateRefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	savedRefreshToken, err := h.service.ValidateRefreshToken(c.Request, refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	newAccessToken, err := h.service.GenerateAccessToken(savedRefreshToken.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	err := h.service.SignOut(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.service.RevokeRefreshToken(c.Request, refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *AuthHandler) Index(c *gin.Context) {
	tmpl, _ := template.New("foo").Parse(indexTemplate)
	tmpl.Execute(c.Writer, nil)
}

var indexTemplate = `<p><a href="/auth/google">Login with Google</a></p>`
