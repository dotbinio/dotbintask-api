package handlers

import (
	"net/http"

	"github.com/dotbinio/taskwarrior-api/internal/taskwarrior"
	"github.com/gin-gonic/gin"
)

// ReportHandler handles report-related requests
type ReportHandler struct {
	client *taskwarrior.Client
}

// NewReportHandler creates a new report handler
func NewReportHandler(client *taskwarrior.Client) *ReportHandler {
	return &ReportHandler{
		client: client,
	}
}

// NextReport handles GET /api/v1/reports/next
func (h *ReportHandler) NextReport(c *gin.Context) {
	tasks, err := h.client.Export("status:pending")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve tasks",
			"code":  "REPORT_FAILED",
		})
		return
	}

	// Sort by urgency and return top tasks
	// Taskwarrior already provides urgency scores
	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"report": "next",
	})
}

// ActiveReport handles GET /api/v1/reports/active
func (h *ReportHandler) ActiveReport(c *gin.Context) {
	tasks, err := h.client.Export("status:pending +ACTIVE")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve tasks",
			"code":  "REPORT_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"report": "active",
	})
}

// CompletedReport handles GET /api/v1/reports/completed
func (h *ReportHandler) CompletedReport(c *gin.Context) {
	tasks, err := h.client.Export("status:completed")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve tasks",
			"code":  "REPORT_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"report": "completed",
	})
}

// WaitingReport handles GET /api/v1/reports/waiting
func (h *ReportHandler) WaitingReport(c *gin.Context) {
	tasks, err := h.client.Export("status:waiting")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve tasks",
			"code":  "REPORT_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"report": "waiting",
	})
}

// AllReport handles GET /api/v1/reports/all
func (h *ReportHandler) AllReport(c *gin.Context) {
	tasks, err := h.client.Export("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve tasks",
			"code":  "REPORT_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks":  tasks,
		"count":  len(tasks),
		"report": "all",
	})
}
