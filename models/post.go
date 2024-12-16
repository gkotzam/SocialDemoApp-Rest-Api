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
