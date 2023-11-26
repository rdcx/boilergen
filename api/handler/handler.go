package handler

import (
	"errors"
	"lightban/api/db"
	"lightban/api/model"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func user(c *gin.Context) (*model.User, error) {
	u, ok := c.Get("user")

	if !ok {
		return nil, errors.New("User not found")
	}

	return u.(*model.User), nil
}

func handleError(c *gin.Context, err error, status int) bool {
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return false
	}
	return true
}
