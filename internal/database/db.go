package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func NewDBConnection() *sql.DB {
	dbPath, err := filepath.Abs("./database.db")
	if err != nil {
		log.Fatalf("Failed to get absolute path: %s", err.Error())
	}

	_, err = os.Stat(dbPath)
	if err != nil || os.IsNotExist(err) {
		log.Fatalf("Database file does not exist: %s", err.Error())
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to open DB: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %s", err.Error())
	}

	fmt.Println("Connected to database:", dbPath)
	return db
}

func CreateMockDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open mock database: %s", err.Error())
	}

	createDocumentsTable := `
		CREATE TABLE IF NOT EXISTS documents (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT,
			locked BOOLEAN DEFAULT 0
		);
	`
	_, err = db.Exec(createDocumentsTable)
	if err != nil {
		log.Fatalf("Failed to create documents table: %s", err.Error())
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			role TEXT CHECK(role IN ('admin', 'employee', 'guest')) NOT NULL,
			last_login_date DATETIME
		);
	`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %s", err.Error())
	}

	createSessionsTable := `
		CREATE TABLE IF NOT EXISTS sessions (
			key TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
		);
	`
	_, err = db.Exec(createSessionsTable)
	if err != nil {
		log.Fatalf("Failed to create sessions table: %s", err.Error())
	}

	populateUsers := `
		INSERT INTO users (username, password, role, last_login_date) VALUES
		('funguy123', 'admin12345', 'admin', NULL),
		('bossman', 'emp12345', 'employee', NULL),
		('grumpy', 'guest12345', 'guest', NULL),
		('jpearson', 'guest12345', 'guest', NULL),
		('fredrick5', 'guest12345', 'guest', NULL),
		('ballhoggary', 'emp12345', 'employee', NULL),
		('erick', 'admin12345', 'admin', NULL),
		('barrylarry', 'emp12345', 'employee', NULL),
		('twotthree', 'guest12345', 'guest', NULL),
		('yap', 'guest12345', 'guest', NULL),
		('boardman45', 'guest12345', 'guest', NULL),
		('1819twenty', 'emp12345', 'employee', NULL),
		('opi', 'guest12345', 'guest', NULL),
		('patrick', 'guest12345', 'guest', NULL),
		('fred111', 'guest12345', 'guest', NULL),
		('secure21', 'guest12345', 'guest', NULL);
	`
	_, err = db.Exec(populateUsers)
	if err != nil {
		log.Fatalf("failed to insert users data: %s", err.Error())
	}

	populateDocuments := `
		INSERT INTO documents (title, content, locked) VALUES
		('Onboarding Document', 'Welcome to the company.', 1),
		('First Doc Ever', 'Everyone can see this document. Your welcome!', 0),
		('Top Secret Case Study', 'All contents of this document should be kept a secret. Only admins should have this document. Do NOT share with anyone!', 1);
	`
	_, err = db.Exec(populateDocuments)
	if err != nil {
		log.Fatalf("failed to insert documents data: %s", err.Error())
	}

	return db
}
