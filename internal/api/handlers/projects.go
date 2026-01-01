package handlers

import (
	"net/http"

	"github.com/dotbinio/taskwarrior-api/internal/taskwarrior"
	"github.com/gin-gonic/gin"
)

// ProjectHandler handles project-related requests
type ProjectHandler struct {
	client *taskwarrior.Client
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(client *taskwarrior.Client) *ProjectHandler {
	return &ProjectHandler{
		client: client,
	}
}

// ListProjects handles GET /api/v1/projects
func (h *ProjectHandler) ListProjects(c *gin.Context) {
	projects, err := h.client.GetProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve projects",
			"code":  "PROJECT_LIST_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"projects": projects,
		"count":    len(projects),
	})
}

// GetProjectTasks handles GET /api/v1/projects/:name/tasks
func (h *ProjectHandler) GetProjectTasks(c *gin.Context) {
	projectName := c.Param("name")

	if projectName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "project name is required",
			"code":  "MISSING_PROJECT_NAME",
		})
		return
	}

	// Sanitize project name
	projectName = taskwarrior.SanitizeInput(projectName)

	tasks, err := h.client.Export("project:" + projectName + " status:pending")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve project tasks",
			"code":  "PROJECT_TASKS_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"project": projectName,
		"tasks":   tasks,
		"count":   len(tasks),
	})
}
