package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	LinkCnt      int       `json:"linkCnt"`
	Comment      string    `json:"comment"`
	Updated_time time.Time `json:"time"`
}

func NewMysqlConnect(dsn string) (*sqlx.DB, error) {
	DB, err := sqlx.Open("mysql", dsn)
	if err != nil {

		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		return nil, err
	}
	return DB, nil

}
func dbxSelect(dbx *sqlx.DB, query string) {
	result := []Comment{}
	//"select * from tests"
	dbx.Select(&result, query)
	for _, row := range result {
		fmt.Println(row)
	}
}

func dbxInsert(dbx *sqlx.DB, jsonbody string) {

	//jsonbody := `{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"}`
	//jsonarraybody := `[{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"},{ "linkCnt":2,"comment":"nice","time": "2022-06-25T00:22:50Z"}]`
	tx := dbx.MustBegin()
	var tablename = "jarujaruch"

	var comments []Comment
	err := json.Unmarshal([]byte(jsonbody), &comments)
	if err != nil {
		fmt.Print(err)
	}
	for _, item := range comments {
		// query := fmt.Sprintf("insert into %s(comment,likecnt,updated_tiime) values(%s,%d,%s)", tablename, item.Comment, item.LinkCnt, item.Updated_time)

		query := fmt.Sprintf("insert into %s (comment, likecnt,updated_time) values (\"%s\",%d,\"%s\")", tablename, item.Comment, item.LinkCnt, item.Updated_time.Format("2006-01-02 15:04:05"))

		fmt.Print(query)
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

func main() {
	dsn := "root:root@tcp(localhost:3306)/jarujaru"

	dbx, dberr := NewMysqlConnect(dsn)
	if dberr != nil {
		fmt.Println("err")
		panic(dberr)
	}
	defer dbx.Close()
	dbx.SetConnMaxLifetime(time.Minute * 3)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)
	//update

	// tx := dbx.MustBegin()
	// var name = "funase"
	// //	for id := 5;id <10;id ++{
	// //		tx.MustExec("insert into tests values(?,?)",id,name+strconv.Itoa(id))
	// //	}
	// for id := 5; id < 10; id++ {
	// 	tx.MustExec("insert into tests(name) values(?)", name+strconv.Itoa(id))
	// }
	// tx.Commit()

	// select
	// type tests struct {
	// 	Id   int
	// 	Name string
	// }

	// result := []tests{}
	// dbx.Select(&result, "select * from tests")
	// for _, row := range result {

	// 	fmt.Println(row)
	// }
	// type Results struct {
	// 	Id   int
	// 	Name string
	// }

	// rows := []Results{}
	// dbx.Select(&rows, "select id,name from tests where id <= 10")
	// for i, _ := range rows {
	// 	fmt.Printf("取得したIDは「%d」：NAMEは「%s」です。\n", rows[i].Id, rows[i].Name)

	//jsonbody := `{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"}`
	jsonbody := `[{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"},{ "linkCnt":2,"comment":"nice","time": "2022-06-25T00:22:50Z"}]`
	dbxInsert(dbx, jsonbody)
	query := "select * from jarujaru"
	dbxSelect(dbx, query)

	fmt.Println("succsess")

	//インクリメントIDを更新するときの書き方
	//jarajaruの動画コメントを保存するために必要な体制を巣作る
	//取得したデータのテスト jsonファイルを作る。
	//構造体を作る
	//保存する

}
