package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 사용자 정보를 나타내는 구조체
type User struct {
	ID       int
	Username string
	Password string
}

// 로그인 요청에 대한 핸들러
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 사용자가 POST 요청을 보내야 함
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 요청의 바디를 JSON으로 디코딩하여 사용자 정보를 추출
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 사용자 정보를 데이터베이스에서 조회
	db, err := sql.Open("mysql", "username:password@tcp(database-host:port)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 사용자 정보를 데이터베이스에서 조회
	query := "SELECT * FROM users WHERE username = ? AND password = ?"
	row := db.QueryRow(query, user.Username, user.Password)

	// 사용자가 존재하는지 확인
	var dbUser User
	err = row.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// 로그인 성공 시 세션을 생성하여 쿠키로 반환
	sessionID := generateSessionID()
	expiration := time.Now().Add(24 * time.Hour) // 쿠키 만료 시간
	cookie := http.Cookie{Name: "sessionID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie)

	// 로그인 성공 메시지 반환
	response := map[string]interface{}{
		"message": "Login successful",
		"user":    dbUser,
	}
	json.NewEncoder(w).Encode(response)
}

// 세션 ID 생성 함수 (임시 예시)
func generateSessionID() string {
	return "random-session-id"
}

func main() {
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
