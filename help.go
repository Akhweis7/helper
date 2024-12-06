package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Amro interface {
	getHandler()
	postHandler()
	err()
	createDatabase()
}

func err(err error) {
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}
}

type PostData struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data PostData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	// Process the data (for example, just print it)
	fmt.Printf("Received post: %+v\n", data)

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post received successfully"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Example response data
	response := map[string]string{"message": "Hello, World!"}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func createDatabase(dbName string) (*sql.DB, error) {
	// Open a connection to the database file (creates it if it doesn't exist)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create a table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	fmt.Println("Database and table created successfully!")
	return db, nil
}
