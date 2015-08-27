package main
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"beelike/models"
)


func main() {
	//db, err := sql.Open("mysql",  "sa:123456@tcp(localhost:3306)/beelikeDB?charset=utf8")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		models.DB_USER, models.DB_PWD, models.DB_HOST, models.DB_PORT, models.DB_NAME))
	defer db.Close()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select 1 as nick")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var nick string
		err = rows.Scan(&nick)
		if err != nil {
			panic(err)
		}

		fmt.Println(nick,"sss")
	}
}
