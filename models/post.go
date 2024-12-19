package models

import (
	"time"

	"github.com/gkotzam/SocialDemoApp-Rest-Api/db"
)

type Post struct {
	ID        int64
	Title     string `binding:"required"`
	PostText  string `binding:"required"`
	UserId    int64
	CreatedAt time.Time
	Comments  []Comment
}

type Comment struct {
	ID          int64
	CommentText string `binding:"required"`
	PostId      int64
	UserId      int64
	CreatedAt   time.Time
}

func (p *Post) Save() error {
	query := `INSERT INTO posts(title, postText, createdAt, userId)
	VALUES (?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.PostText, p.CreatedAt, p.UserId)

	return err
}

func (c *Comment) Save() error {
	query := `INSERT INTO comments(commentText, postId, userId, createdAt)
	VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(c.CommentText, c.PostId, c.UserId, c.CreatedAt)

	return err

}

func GetPostById(id int64) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.PostText, &post.UserId, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	post.Comments, err = GetCommentsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetCommentById(id int64) (*Comment, error) {
	query := "SELECT * FROM comments WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var comment Comment
	err := row.Scan(&comment.ID, &comment.CommentText, &comment.PostId, &comment.UserId, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func GetAllPosts() ([]Post, error) {
	query := "SELECT * FROM posts"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.PostText, &post.UserId, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		post.Comments, err = GetCommentsByPostId(post.ID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}

func GetCommentsByPostId(id int64) ([]Comment, error) {
	query := "SELECT * FROM comments WHERE postId = ?"
	rows, err := db.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.CommentText, &comment.PostId, &comment.UserId, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil

}
