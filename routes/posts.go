package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gkotzam/SocialDemoApp-Rest-Api/models"
)

func createPost(context *gin.Context) {
	var post models.Post

	err := context.ShouldBindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// TODO: after login get userID
	post.UserId = 1 // temp
	post.CreatedAt = time.Now().UTC()

	err = post.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save post."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post Created!"})
}
