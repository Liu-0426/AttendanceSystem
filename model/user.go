package model

import (
    "database/sql"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	
)

var db *sql.DB

type User struct {
	ID       int
	Username string
	Password string
	Role    string
}

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Database connection established")
}
func GetUserByUsername(username string) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func GetAllUsers() ([]User, error) {
    var users []User
    rows, err := db.Query("SELECT id, username FROM users where role = 'employee'")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Username)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}
func SaveUserToDB(user User) error {
	// 插入用戶資料
	_, err := db.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}

	return nil
}
func BindLineUserID(username, lineUserID string) error {
    _, err := db.Exec("UPDATE users SET line_user_id = ? WHERE username = ?", lineUserID, username)
    return err
}
func GetUserIDByUsername(username string) (int, error) {
    var userID int
    err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
    if err != nil {
        return 0, err
    }
    return userID, nil
}