package main

import (
	"main/db"
	"main/util"
	"net/http"
	"time"

	//"flag"

	"log"

	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

// func getVideoIDs(uploadsID string, service *youtube.Service) (videoIDs []string) {
// 	playlistsCall := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(uploadsID).MaxResults(50)

// 	playlistsResponse, err := playlistsCall.Do()
// 	if err != nil {
// 		log.Fatalf("Error call YouTube API: %v", err)
// 	}

// 	for _, item := range playlistsResponse.Items {
// 		videoIDs = append(videoIDs, item.ContentDetails.VideoId)
// 	}

// 	if playlistsResponse.NextPageToken != "" {
// 		nextPageToken := playlistsResponse.NextPageToken
// 		for {
// 			nextCall := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(uploadsID).PageToken(nextPageToken).MaxResults(50)
// 			nextResponse, err := nextCall.Do()
// 			if err != nil {
// 				log.Fatalf("Error call YouTube API: %v", err)
// 			}
// 			for _, nextItem := range nextResponse.Items {
// 				videoIDs = append(videoIDs, nextItem.ContentDetails.VideoId)
// 			}
// 			nextPageToken = nextResponse.NextPageToken
// 			if nextPageToken == "" {
// 				break
// 			}
// 		}
// 	}

// 	return videoIDs
// }

type VideoData struct {
	VideoID       string
	Title         string
	ViewCount     uint64
	LikeCount     uint64
	DislikeCount  uint64
	FavoriteCount uint64
	CommentCount  uint64
	PublishedAt   string
}

func main() {

	//log
	util.LoggingSetting("system.log")
	log.Println("Logging start")

	//	database
	dsn := "root:root@tcp(localhost:3306)/jarujaru?parseTime=true"
	dbx = db.NewMysqlConnect(dsn)
	log.Println("DB connect success")

	defer dbx.Close()
	dbx.SetConnMaxLifetime(time.Minute * 3)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)

	//router
	router := GetRouter()
	http.ListenAndServe(":8080", router)
	log.Println("router success")

	// fmt.Println(len(responseBodyArray))
	// fmt.Println(responseBodyArray)
	// //dbに保存
	// dsn := "root:root@tcp(localhost:3306)/jarujaru"

	// dbx, dberr := db.NewMysqlConnect(dsn)
	// if dberr != nil {
	// 	fmt.Println("err")
	// 	panic(dberr)
	// }
	// defer dbx.Close()
	// dbx.SetConnMaxLifetime(time.Minute * 3)
	// dbx.SetMaxOpenConns(10)
	// dbx.SetMaxIdleConns(10)

	// db.DbxInsert(dbx, responseBodyArray)

	// query := "select * from jarujaru"
	// db.DbxSelect(dbx, query)
	//------------------------------------------------------------------------

}

//linkカウントも取得する ok
//綺麗にする　ok

//動画のURLを書く
//取得したコメントをデータベースに保存する
//DBに保存する

//時間を曜日で分ける
//月で分ける
//時間で分ける
//user

//より簡単にする
// - git ignore
// - 設定ファイルから環境変数を取得
// - ログを書き込む
// - 配列のアドレス

//typeを分類する方法を考える
//デザインを考える。
//動画を毎日保存する。

//　今はやらない
// - viideoIDを三日分とかとる

// 	// Group video, channel, and playlist results in separate lists.
// 	videos := make(map[string]string)
// 	channels := make(map[string]string)
// 	playlists := make(map[string]string)

// 	// Iterate through each item and add it to the correct list.
// 	for _, item := range response.Items {
// 		switch item.Id.Kind {
// 		case "youtube#video":
// 			videos[item.Id.VideoId] = item.Snippet.Title
// 		case "youtube#channel":
// 			channels[item.Id.ChannelId] = item.Snippet.Title
// 		case "youtube#playlist":
// 			playlists[item.Id.PlaylistId] = item.Snippet.Title
// 		}
// 	}

// 	printIDs("Videos", videos)
// 	printIDs("Channels", channels)
// 	printIDs("Playlists", playlists)
// }

// // Print the ID and title of each result in a list as well as a name that
// // identifies the list. For example, print the word section name "Videos"
// // above a list of video search results, followed by the video ID and title
// // of each matching video.
// func printIDs(sectionName string, matches map[string]string) {
// 	fmt.Printf("%v:\n", sectionName)
// 	for id, title := range matches {
// 		fmt.Printf("[%v] %v\n", id, title)
// 	}
// 	fmt.Printf("\n\n")
// }
