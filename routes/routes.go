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
	// GET => /comments/:id
	server.GET("/comments/:id", getComment)

	// ---- POST ----
	// POST => /signup
	server.POST("/signup", signup)
	// POST => /login
	server.POST("/login", login)
	// POST => /posts
	server.POST("/posts", createPost)
	// POST => /posts/:postId/comments
	server.POST("/posts/:postId/comments", createComment)

	// ---- PUT ----
	// PUT => /posts/:id
	server.PUT("/posts/:id", updatePost)
	// PUT => /comments/:id
	server.PUT("/comments/:id", updateComment)

	// ---- DELETE ----
	// DELETE => /posts
	server.DELETE("/posts/:id", deletePost)
	// DELETE => /comments
	server.DELETE("/comments/:id", deleteComment)

}
