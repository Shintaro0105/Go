package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var dbConn *pgx.Conn

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// .env 読み込み
	err := godotenv.Load(".devcontainer/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// DB接続文字列
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	// DB接続
	dbConn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbConn.Close(context.Background())

	fmt.Println("✅ Connected to database!")

	// ハンドラ登録
	http.HandleFunc("/users", getUsers)

	// サーバー起動
	fmt.Println("🚀 Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := dbConn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "DB query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			http.Error(w, "DB scan failed", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
