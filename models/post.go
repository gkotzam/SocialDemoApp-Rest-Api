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

func GetPostById(id int64) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.PostText, &post.UserId, &post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
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
		posts = append(posts, post)
	}

	return posts, nil

}
