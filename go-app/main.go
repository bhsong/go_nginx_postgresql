package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" 	// postgres 드라이버
)

var db *sql.DB

func main() {
	// 1. DB 연결 정보 (환경변수에서 읽어옴)
	connStr := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// 2. DB 연결 시도
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			fmt.Println("DB 연결 성공!")
			break
		}
		fmt.Println("DB 연결 대기중...", err)
		time.Sleep(2 * time.Second)
	}

	// 테이블 생성 (없으면 생성)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS visits (id SERIAL PRIMARY KEY, visited_at TIMESTAMP)")
	if err != nil {
		log.Fatal("테이블 생성 실패:", err)
	}

	// 3. 웹 서버 핸들러
	http.HandleFunc("/api", handleRequest)

	// 4. 서버 시작
	fmt.Println("GO 서버가 8080 포트에서 시작됩니다")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 방문 기록 저장 (INSERT)
	_, err := db.Exec("INSERT INTO visits (visited_at) VALUES ($1)", time.Now())
	if err != nil {
		http.Error(w, "DB INSERT Error", 500)
		return
	}

	// 총 방문자 수 조회 (SELECT)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM visits").Scan(&count)
	if err != nil {
		http.Error(w, "DB Select Error", 500)
		return
	}

	// 결과 출력
	fmt.Fprintf(w, "Welcome! You are visitor number: %d\nDB Status: Data Saved!!!", count)
}
