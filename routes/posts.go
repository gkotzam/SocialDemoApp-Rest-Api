package routes

import (
	"net/http"
	"strconv"
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

func getPost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse post id."})
		return
	}

	event, err := models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post. Try again later."})
		return
	}

	context.JSON(http.StatusOK, event)

}

func getPosts(context *gin.Context) {
	events, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}
