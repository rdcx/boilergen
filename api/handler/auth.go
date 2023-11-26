package handler

import (
	"lightban/util"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Auth(f func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" || len(token) < 7 {
			c.JSON(401, gin.H{
				"message": "Missing authorization header",
			})
			return
		}

		// remove the "Bearer " prefix
		token = token[7:]

		claims, err := util.ParseClaims(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		// check if the user_id claim exists
		if _, ok := claims["user_id"]; !ok {
			c.JSON(401, gin.H{
				"message": "Missing user_id claim",
			})
			return
		}

		// convert jwt.Number to uint
		v, ok := claims["user_id"].(float64)

		if !ok {
			c.JSON(401, gin.H{

				"message": "Invalid user_id claim",
			})
			return
		}

		// set the user_id claim to the context
		userId := uint(v)

		// check if the user exists
		user, err := h.db.GetUser(userId)
		if err != nil {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}

		// set the user to the context
		c.Set("user", user)

		// continue
		f(c)
	}
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=32,alphanum"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if _, err := h.db.CreateUser(req.Email, req.Username, req.Password); err != nil {
		c.JSON(400, gin.H{
			"message": "Username or email already exists",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Successfully registered",
	})
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	token, err := h.db.Login(req.Username, req.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
