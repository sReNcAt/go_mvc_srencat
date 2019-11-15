package mvc_model

import (
    "log"
    "database/sql" 
    _ "github.com/go-sql-driver/mysql"
)

func Model(){
    db, err := sql.Open(dbconn())
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var idx int
    var username string
    var password string
    var flags int
    var createdAt string
    var updatedAt string
    
    // 복수 Row를 갖는 SQL 쿼리
    rows, err := db.Query("select * from users limit ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() //반드시 닫는다 (지연하여 닫기)
    
    for rows.Next() {
        err := rows.Scan(&idx, &username, &password, &flags, &createdAt, &updatedAt)
        if err != nil {
            log.Fatal(err)
        }		
    }

}