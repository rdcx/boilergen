package handler

import (
	"lightban/api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createProjectRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

func (h *Handler) GetProject(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	usr := u.(*model.User)
	// refresh the user
	user, err := h.db.GetUser(usr.ID)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	project, err := h.db.GetProject(uint(id))

	if err != nil {
		c.JSON(404, gin.H{"error": "Project not found"})
		return
	}

	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.JSON(200, project)
}

func (h *Handler) GetProjects(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	projects, err := h.db.GetProjectsByUserID(user.ID)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}

	c.JSON(200, projects)
}

func (h *Handler) CreateProject(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project := &model.Project{
		Name:   req.Name,
		UserID: user.ID,
	}

	if err := h.db.CreateProject(project); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, project)
}

func (h *Handler) DeleteProject(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	project, err := h.db.GetProject(uint(id))
	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.db.DeleteProjectByID(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, gin.H{"message": "Project deleted"})
}

type updateProjectRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

func (h *Handler) UpdateProject(c *gin.Context) {
	u, ok := c.Get("user")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	user := u.(*model.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	project, err := h.db.GetProject(uint(id))
	if project.UserID != user.ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var req updateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	project.Name = req.Name

	if err := h.db.UpdateProject(project); err != nil {
		c.JSON(500, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(200, project)
}
