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
	job := r.Group("/job")
	{
		job.GET("", s.RequireAccessToken, s.jobApplicationHandler.FindAllByUserId)
		job.GET("/:id", s.RequireAccessToken, s.jobApplicationHandler.FindById)
		job.POST("", s.RequireAccessToken, s.jobApplicationHandler.Create)
		job.PATCH("/:id", s.RequireAccessToken, s.jobApplicationHandler.Update)
		job.DELETE("/:id", s.RequireAccessToken, s.jobApplicationHandler.Delete)
	}
	interview := r.Group("/interview")
	{
		interview.POST("/:jobApplicationId", s.RequireAccessToken, s.interviewHandler.Create)
		interview.GET("/:jobApplicationId/:interviewId", s.RequireAccessToken, s.interviewHandler.FindById)
		interview.GET("/:jobApplicationId", s.RequireAccessToken, s.interviewHandler.FindAllByApplicationId)
		interview.PATCH("/:jobApplicationId/:interviewId", s.RequireAccessToken, s.interviewHandler.Update)
		interview.DELETE("/:jobApplicationId/:interviewId", s.RequireAccessToken, s.interviewHandler.Delete)
	}

	contact := r.Group("/contact")
	{
		contact.POST("", s.RequireAccessToken, s.contactHandler.Create)
		contact.GET("", s.RequireAccessToken, s.contactHandler.FindAllByUserId)
		contact.GET("/:id", s.RequireAccessToken, s.contactHandler.FindById)
		contact.PATCH("/:id", s.RequireAccessToken, s.contactHandler.Update)
		contact.DELETE(":id", s.RequireAccessToken, s.contactHandler.Delete)
	}

	timeline := r.Group("/timeline/:jobApplicationId")
	{
		timeline.POST("", s.RequireAccessToken, s.timelineHandler.Create)
		timeline.GET("", s.RequireAccessToken, s.timelineHandler.FindAllByApplicationId)
		timeline.GET("/:timelineId", s.RequireAccessToken, s.timelineHandler.FindById)
		timeline.PATCH("/:timelineId", s.RequireAccessToken, s.timelineHandler.Update)
		timeline.DELETE("/:timelineId", s.RequireAccessToken, s.timelineHandler.Delete)
	}

	return r
}
