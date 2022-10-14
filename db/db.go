package db

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	LinkCnt      int       `json:"linkCnt"`
	Comment      string    `json:"comment"`
	VideoID      string    `json:"videoid"`
	UserName     string    `json:"username"`
	Updated_time time.Time `json:"updated_time"`
	Month_time   int       `json:"month_time"`
	WeekDay_time string    `json:"weekday_time"`
	Hour_time    int       `json:"hour_time"`
}

func NewMysqlConnect(dsn string) *sqlx.DB {
	DB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return DB

}
func DbxSelect(dbx *sqlx.DB, query string) []Comment {
	result := []Comment{}
	//"select * from tests"
	dbx.Select(&result, query)

	return result

}

func DbxInsert(dbx *sqlx.DB, comments []Comment) {

	//jsonbody := `{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"}`
	//jsonarraybody := `[{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"},{ "linkCnt":2,"comment":"nice","time": "2022-06-25T00:22:50Z"}]`
	tx := dbx.MustBegin()
	var tablename = "jarujaruch"

	// var comments []Comment
	// err := json.Unmarshal([]byte(jsonbody), &comments)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	for _, item := range comments {
		// query := fmt.Sprintf("insert into %s(comment,likecnt,updated_tiime) values(%s,%d,%s)", tablename, item.Comment, item.LinkCnt, item.Updated_time)
		query := fmt.Sprintf("insert into %s (likecnt,comment, videoid,username,updated_time,month_time,weekday_time,hour_time) values (%d,\"%s\",\"%s\",\"%s\",\"%s\",%d,\"%s\",%d)", tablename, item.LinkCnt, item.Comment, item.VideoID, item.UserName, item.Updated_time.Format("2006-01-02 15:04:05"), item.Month_time, item.WeekDay_time, item.Hour_time)
		tx.MustExec(query)
	}
	tx.Commit()

	// return commnet

}

// tx := dbx.MustBegin()
// var name = "funase"
// for id := 5; id < 10; id++ {
// 	tx.MustExec("insert into tests(name) values(?)", name+strconv.Itoa(id))
// }
// // for id := 5; id < 10; id++ {
// // 	tx.MustExec("insert into tests(name) values(?)", name+strconv.Itoa(id))
// // }
// tx.Commit()
