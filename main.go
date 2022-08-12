package main
import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "fmt"
)


var Db *sqlx.DB
func init()  {
    db, err := sqlx.Open("mysql", "user:Password@tcp(ip:3306)/Users")
    if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }
    Db = db
    fmt.Println("hello db")

}
func main()  {
    result, err := Db.Exec("INSERT INTO userinfo (username, password,email) VALUES (?, ?, ?)","li","li444","li@gmail.com")
    if err != nil{
        fmt.Println("insert failed,error： ", err)
        return
    }
    id,_ := result.LastInsertId()
    fmt.Println("insert id is :",id)
}