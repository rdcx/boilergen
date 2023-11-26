package router

import (
	"lightban/api/handler"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUp(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)

	r.GET("/user", h.Auth(h.GetUser))

	r.GET("/projects", h.Auth(h.GetProjects))
	r.GET("/projects/:id", h.Auth(h.GetProject))

	r.POST("/projects", h.Auth(h.CreateProject))
	r.PUT("/projects/:id", h.Auth(h.UpdateProject))
	r.DELETE("/projects/:id", h.Auth(h.DeleteProject))

	return r
}
