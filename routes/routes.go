package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// ---- GET ----
	// GET => /users
	server.GET("/users", getUsers)
	// GET => /posts
	server.GET("/posts", getPosts)
	// GET => /posts/:id
	server.GET("/posts/:id", getPost)

	// ---- POST ----
	// POST => /signup
	server.POST("/signup", signup)
	// POST => /login
	// server.POST("/login")
	// POST => /posts
	server.POST("/posts", createPost)
	// POST => /posts/:postId/comments
	server.POST("/posts/:postId/comments", createComment)

	// ---- DELETE ----
	// DELETE -> /posts

}
