package models

import (
	"time"

	"github.com/gkotzam/SocialDemoApp-Rest-Api/db"
	"github.com/gkotzam/SocialDemoApp-Rest-Api/utils"
)

type User struct {
	ID        int64
	Username  string `binding:"required"`
	Email     string `binding:"required"`
	Password  string `binding:"required"`
	CreatedAt time.Time
}

// Saves user to Database
func (u *User) Save() error {
	query := `INSERT INTO users(username,email,password,createdAt)
	VALUES (?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Email, hashedPassword, u.CreatedAt)

	return err
}

// Returns all saved users ([]User) from Database
func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil

}
