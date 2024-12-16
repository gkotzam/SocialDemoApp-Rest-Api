package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// ---- GET ----
	// GET => /users
	server.GET("/users", getUsers)

	// ---- POST ----
	// POST => /signup
	server.POST("/signup", signup)
}
