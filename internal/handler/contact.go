package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hadihalimm/jobtagger-backend/internal/model/request"
	"github.com/hadihalimm/jobtagger-backend/internal/service"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func (h *ContactHandler) Create(c *gin.Context) {
	var req request.CreateContact
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetString("currentUser")
	contact, err := h.contactService.Create(c.Request.Context(), req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contact)
}

func (h *ContactHandler) FindById(c *gin.Context) {
	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := h.contactService.FindById(c.Request.Context(), contactId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) FindAllByUserId(c *gin.Context) {
	userId := c.GetString("currentUser")
	contacts, err := h.contactService.FindAllByUserId(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

func (h *ContactHandler) Update(c *gin.Context) {
	var req request.UpdateContact
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedContact, err := h.contactService.Update(c.Request.Context(), contactId, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedContact)
}

func (h *ContactHandler) Delete(c *gin.Context) {
	contactId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.contactService.Delete(c.Request.Context(), contactId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
}
