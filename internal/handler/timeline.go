package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
)

type TimelineHandler struct {
	timelineService service.TimelineService
}

func NewTimelineHandler(service service.TimelineService) *TimelineHandler {
	return &TimelineHandler{timelineService: service}
}

func (h *TimelineHandler) Create(c *gin.Context) {
	var request request.CreateTimeline
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobApplicationId, err := strconv.Atoi(c.Param("jobApplicationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timeline, err := h.timelineService.Create(c.Request.Context(), request, jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, timeline)
}

func (h *TimelineHandler) FindAllByApplicationId(c *gin.Context) {
	jobApplicationId, err := strconv.Atoi(c.Param("jobApplicationId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timelines, err := h.timelineService.FindAllByApplicationId(c.Request.Context(), jobApplicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timelines)
}

func (h *TimelineHandler) FindById(c *gin.Context) {
	timelineId, err := strconv.Atoi(c.Param("timelineId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timeline, err := h.timelineService.FindById(c.Request.Context(), timelineId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *TimelineHandler) Update(c *gin.Context) {
	var request request.UpdateTimeline
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timelineId, err := strconv.Atoi(c.Param("timelineId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeline, err := h.timelineService.Update(c.Request.Context(), timelineId, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timeline)
}

func (h *TimelineHandler) Delete(c *gin.Context) {
	timelineId, err := strconv.Atoi(c.Param("timelineId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.timelineService.Delete(c.Request.Context(), timelineId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "timeline deleted successfully"})
}
