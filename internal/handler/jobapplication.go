package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
)

type JobApplicationHandler struct {
	service service.JobApplicationService
}

func NewJobApplicationHandler(service service.JobApplicationService) *JobApplicationHandler {
	return &JobApplicationHandler{service: service}
}
func (h *JobApplicationHandler) Create(c *gin.Context) {
	var jobApplicationRequest request.CreateJobApplication
	if err := c.ShouldBindJSON(&jobApplicationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("currentUser")
	jobApplication, err := h.service.Create(c.Request.Context(), jobApplicationRequest, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobApplication)
}

func (h *JobApplicationHandler) FindAllByUserId(c *gin.Context) {
	userId := c.GetString("currentUser")
	jobApplications, err := h.service.FindAllByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobApplications)
}

func (h *JobApplicationHandler) FindById(c *gin.Context) {
	jobApplicationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobApplication, err := h.service.FindById(c.Request.Context(), jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jobApplication)
}

func (h *JobApplicationHandler) Update(c *gin.Context) {
	var jobApplicationRequest request.UpdateJobApplication
	if err := c.ShouldBindJSON(&jobApplicationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobApplicationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedJob, err := h.service.Update(c.Request.Context(), jobApplicationId, jobApplicationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedJob)
}

func (h *JobApplicationHandler) Delete(c *gin.Context) {
	jobApplicationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Delete(c.Request.Context(), jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "job application deleted successfully"})
}
