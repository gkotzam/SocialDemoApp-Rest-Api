package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gkotzam/SocialDemoApp-Rest-Api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// ---- GET ----
	// GET => /users
	// server.GET("/users", middlewares.Authenticate, getUsers) TODO: add admin User
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
	server.POST("/posts", middlewares.Authenticate, createPost)
	// POST => /posts/:postId/comments
	server.POST("/posts/:postId/comments", middlewares.Authenticate, createComment)

	// ---- PUT ----
	// PUT => /posts/:id
	server.PUT("/posts/:id", middlewares.Authenticate, updatePost)
	// PUT => /comments/:id
	server.PUT("/comments/:id", middlewares.Authenticate, updateComment)

	// ---- DELETE ----
	// DELETE => /posts
	server.DELETE("/posts/:id", middlewares.Authenticate, deletePost)
	// DELETE => /comments
	server.DELETE("/comments/:id", middlewares.Authenticate, deleteComment)

}
