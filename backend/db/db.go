package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error

    // Get database credentials from environment variables
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Connection string
    connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbUser, dbPassword, dbName)

    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("Database not reachable:", err)
    }

    fmt.Println("Connected to database successfully")

    // ✅ Ensure "tasks" table exists
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        task TEXT NOT NULL
    );`

    _, err = DB.Exec(createTableQuery)
    if err != nil {
        log.Fatal("Error creating table:", err)
    }

    fmt.Println("Table 'tasks' is ready")
}

