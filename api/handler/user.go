package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetUser(c *gin.Context) {
	u, ok := c.Get("user")

	if !ok {
		c.JSON(500, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(200, u)
}
