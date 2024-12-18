package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "social-demo-app.db")

	if err != nil {
		panic("Could not connect to Database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		createdAt DATETIME NOT NULL
		);
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	createPostsTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		postText TEXT NOT NULL,
		userId INTEGER NOT NULL,
		createdAt DATETIME NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
		);
	`
	_, err = DB.Exec(createPostsTable)

	if err != nil {
		panic("Could not create posts table.")
	}

	createCommentsTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		commentText TEXT NOT NULL,
		postId INTEGER NOT NUL,
		userId INTEGER NOT NULL,
		createdAt DATETIME NOT NULL,
		FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE,
  		FOREIGN KEY (postId) REFERENCES posts(id) ON DELETE CASCADE
		);
	`

	_, err = DB.Exec(createCommentsTable)

	if err != nil {
		panic("Could not create comments table.")
	}

}
