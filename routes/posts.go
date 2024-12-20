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

	post.UserId = context.GetInt64("userId")
	post.CreatedAt = time.Now().UTC()

	err = post.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save post."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post Created!"})
}

func createComment(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("postId"), 10, 64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse post id."})
		return
	}

	_, err = models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post. Try again later."})
		return
	}

	// post exists -> save comment on that post

	var comment models.Comment
	err = context.ShouldBindJSON(&comment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	comment.UserId = context.GetInt64("userId")
	comment.PostId = postId
	comment.CreatedAt = time.Now().UTC()

	err = comment.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save comment."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Comment Created!"})

}

func deletePost(context *gin.Context) {
	PostId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse Post id."})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostById(PostId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Post. Try again later."})
		return
	}

	if post.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete post."})
		return
	}

	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete Post."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post Deleted!"})

}

func deleteComment(context *gin.Context) {
	commentId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse comment id."})
		return
	}

	userId := context.GetInt64("userId")
	comment, err := models.GetCommentById(commentId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch comment. Try again later."})
		return
	}

	if comment.UserId != userId && comment.PostId != userId {
		// user is not comment creator or post creator
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete comment."})
		return
	}

	err = comment.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete Comment."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Comment Deleted!"})

}

func updatePost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse post id."})
		return
	}

	userId := context.GetInt64("userId")
	post, err := models.GetPostById(postId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post. Try again later."})
		return
	}

	if post.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update post."})
		return
	}

	var updatedPost models.Post
	err = context.ShouldBindJSON(&updatedPost)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedPost.ID = postId
	err = updatedPost.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update post."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Post Updated!"})

}

func updateComment(context *gin.Context) {
	commentId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse comment id."})
		return
	}

	userId := context.GetInt64("userId")
	comment, err := models.GetCommentById(commentId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch comment. Try again later."})
		return
	}

	if comment.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update comment."})
		return
	}

	var updateComment models.Comment
	err = context.ShouldBindJSON(&updateComment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updateComment.ID = commentId
	err = updateComment.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update comment."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Comment Updated!"})

}

func getPost(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse post id."})
		return
	}

	post, err := models.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch post. Try again later."})
		return
	}

	context.JSON(http.StatusOK, post)

}

func getComment(context *gin.Context) {
	commentId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could parse comment id."})
		return
	}

	comment, err := models.GetCommentById(commentId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch comment. Try again later."})
		return
	}

	context.JSON(http.StatusOK, comment)

}

func getPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch posts. Try again later."})
		return
	}
	context.JSON(http.StatusOK, posts)
}
