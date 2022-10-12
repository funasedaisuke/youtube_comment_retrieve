package db

import (
	// "database/sql"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	LinkCnt      int       `json:"linkCnt"`
	Comment      string    `json:"comment"`
	Updated_time time.Time `json:"time"`
	VideoID      string    `json:"videoid"`
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
func DbxSelect(dbx *sqlx.DB, query string) {
	result := []Comment{}
	//"select * from tests"
	dbx.Select(&result, query)
	for _, row := range result {
		fmt.Println(row)
	}
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

		query := fmt.Sprintf("insert into %s (comment, likecnt,updated_time,videoid) values (\"%s\",%d,\"%s\",\"%s\")", tablename, item.Comment, item.LinkCnt, item.Updated_time.Format("2006-01-02 15:04:05"), item.VideoID)

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

func Action() {
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
	// jsonbody := `[{ "linkCnt":0,"comment":"good","time": "2022-06-25T00:22:50Z"},{ "linkCnt":2,"comment":"nice","time": "2022-06-25T00:22:50Z"}]`
	// DbxInsert(dbx, jsonbody)
	// query := "select * from jarujaru"
	// DbxSelect(dbx, query)

	fmt.Println("succsess")

	//インクリメントIDを更新するときの書き方
	//jarajaruの動画コメントを保存するために必要な体制を巣作る
	//取得したデータのテスト jsonファイルを作る。
	//構造体を作る
	//保存する

}
