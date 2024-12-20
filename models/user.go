package models

import (
	"errors"
	"time"

	"github.com/gkotzam/SocialDemoApp-Rest-Api/db"
	"github.com/gkotzam/SocialDemoApp-Rest-Api/utils"
)

type User struct {
	ID        int64
	Username  string
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

	// Default username
	if u.Username == "" {
		u.Username = "user"
	}

	_, err = stmt.Exec(u.Username, u.Email, hashedPassword, u.CreatedAt)

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var storedPassword string
	err := row.Scan(&u.ID, &storedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, storedPassword)

	if !passwordIsValid {
		return errors.New("credentials Invalid")
	}

	return nil
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
