package mvc_model

import (
	"fmt"
	"log"
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
)

type user_info struct {
	Idx int `json:"idx"`
	Username string `json:"username"`
}

func Login(input_username string,input_password string)(int,string){
	db, err := sql.Open(dbconn())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var idx int
	var username string
	var queryUserData user_info
	// 복수 Row를 갖는 SQL 쿼리
	rows, err := db.Query("select idx,username from users WHERE username=? AND password=? limit 1",input_username,input_password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	
	for rows.Next() {
		err := rows.Scan(&idx, &username)
		queryUserData = user_info{idx,username}
		fmt.Println(queryUserData)
		if err != nil {
			log.Fatal(err)
		}
		break
	}
	return idx,username
}