package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.authHandler.Index)
	auth := r.Group("/auth")
	{
		auth.GET("/:provider", s.authHandler.SignIn)
		auth.GET("/:provider/callback", s.authHandler.AuthCallback)
		auth.GET("/refresh", s.authHandler.RotateRefreshToken)
		auth.GET("/signout", s.authHandler.SignOut)
	}

	return r
}
