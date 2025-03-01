package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
)

type InterviewHandler struct {
	service service.InterviewService
}

func NewInterviewHandler(service service.InterviewService) *InterviewHandler {
	return &InterviewHandler{service: service}
}

func (h *InterviewHandler) Create(c *gin.Context) {
	var request request.CreateInterview
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobApplicationId, err := strconv.Atoi(c.Param("jobApplicationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	interview, err := h.service.Create(c.Request.Context(), request, jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, interview)
}

func (h *InterviewHandler) FindById(c *gin.Context) {
	interviewId, err := strconv.Atoi(c.Param("interviewId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interview, err := h.service.FindById(c.Request.Context(), interviewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, interview)
}

func (h *InterviewHandler) FindAllByApplicationId(c *gin.Context) {
	jobApplicationId, err := strconv.Atoi(c.Param("jobApplicationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	interviews, err := h.service.FindAllByApplicationId(c.Request.Context(), jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, interviews)
}

func (h *InterviewHandler) Update(c *gin.Context) {
	var request request.UpdateInterview
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interviewId, err := strconv.Atoi(c.Param("interviewId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interview, err := h.service.Update(c.Request.Context(), interviewId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, interview)
}

func (h *InterviewHandler) Delete(c *gin.Context) {
	interviewId, err := strconv.Atoi(c.Param("interviewId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Delete(c.Request.Context(), interviewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "interview deleted successfully"})
}
