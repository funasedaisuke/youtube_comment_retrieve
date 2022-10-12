package main

import (
	"context"
	"fmt"
	"main/db"
	"main/util"
	"strconv"
	"strings"
	"time"

	//"flag"

	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

//やること
//1動画のコメントを全て取得できるようにする。
//-ちょうどいい動画を探す

//動画 (ダイアン)
//HK8CzJm8gdM
//

const developerKey = ""

func createService() (service *youtube.Service) {
	ctx := context.Background()
	//service, err := youtube.NewService(ctx)
	service, err := youtube.NewService(ctx, option.WithAPIKey(developerKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	return service
}

func getCommentThread(nextPageToken, videoID string, service *youtube.Service, responseBodyArray []db.Comment) (responseBody []db.Comment, nextpagetoken string) {
	// Make the API call to YouTube.
	var call *youtube.CommentThreadsListCall
	if nextPageToken == "" {
		call = service.CommentThreads.List([]string{"snippet"}).
			VideoId(videoID).MaxResults(50)
	} else {
		call = service.CommentThreads.List([]string{"snippet"}).
			VideoId(videoID).PageToken(nextPageToken).MaxResults(50)
	}
	response, err := call.Do()
	handleError(err, "")
	// body, _ := response.MarshalJSON()

	nextpagetoken = response.NextPageToken
	// fmt.Println("---------------")
	// fmt.Println(string(body))
	// fmt.Println("---------------")
	// index := 0
	jst, _ := time.LoadLocation("Asia/Tokyo")
	for _, item := range response.Items {
		//時間を曜日で分ける
		//月で分ける
		//時間で分ける

		commentInstance := db.Comment{}
		commentInstance.Comment = strings.Replace(item.Snippet.TopLevelComment.Snippet.TextDisplay, "\"", "", -1)
		commentInstance.LinkCnt = int(item.Snippet.TopLevelComment.Snippet.LikeCount)
		commentInstance.UserName = item.Snippet.TopLevelComment.Snippet.AuthorDisplayName
		commentInstance.Updated_time = util.TimeToJapan(item.Snippet.TopLevelComment.Snippet.PublishedAt, jst)
		commentInstance.Month_time, _ = strconv.Atoi(commentInstance.Updated_time.Month().String())
		commentInstance.WeekDay_time = commentInstance.Updated_time.Weekday().String()
		commentInstance.Hour_time = commentInstance.Updated_time.Hour()

		commentInstance.VideoID = item.Snippet.TopLevelComment.Snippet.VideoId
		responseBodyArray = append(responseBodyArray, commentInstance)
		// fmt.Printf("comment: %v\n", commentInstance.Comment)
		// fmt.Printf("likeCnt: %v\n", commentInstance.LinkCnt)
		// fmt.Printf("time: %v\n", commentInstance.Updated_time)
		// fmt.Printf("videoID: %v\n", commentInstance.VideoID)

	}
	// fmt.Println(index)
	return responseBodyArray, nextpagetoken
}

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

// func getVideoDatas(videoIDs []string, service *youtube.Service) (videoDatas []VideoData) {
// 	for _, videoID := range videoIDs {
// 		call := service.Videos.List([]string{"snippet", "statistics"}).Id(videoID)
// 		response, err := call.Do()
// 		if err != nil {
// 			log.Fatalf("Error call YouTube API: %v", err)
// 		}

// 		item := response.Items[0]
// 		videoData := VideoData{
// 			VideoID:       videoID,
// 			Title:         item.Snippet.Title,
// 			ViewCount:     item.Statistics.ViewCount,
// 			LikeCount:     item.Statistics.LikeCount,
// 			DislikeCount:  item.Statistics.DislikeCount, // ※oauth認証じゃないと取得できない
// 			FavoriteCount: item.Statistics.FavoriteCount,
// 			CommentCount:  item.Statistics.CommentCount,
// 			PublishedAt:   item.Snippet.PublishedAt,
// 		}
// 		videoDatas = append(videoDatas, videoData)
// 	}
// 	return videoDatas
// }

func main() {
	// flag.Parse()
	videoid := "HK8CzJm8gdM"
	service := createService()
	nextPageToken := ""
	var responseBodyArray []db.Comment
	for {
		responseBodyArray, nextPageToken = getCommentThread(nextPageToken, videoid, service, responseBodyArray)
		fmt.Println(len(responseBodyArray))
		if nextPageToken == "" {
			break
		}
	}
	fmt.Println(len(responseBodyArray))
	fmt.Println(responseBodyArray)
	//dbに保存
	dsn := "root:root@tcp(localhost:3306)/jarujaru"

	dbx, dberr := db.NewMysqlConnect(dsn)
	if dberr != nil {
		fmt.Println("err")
		panic(dberr)
	}
	defer dbx.Close()
	dbx.SetConnMaxLifetime(time.Minute * 3)
	dbx.SetMaxOpenConns(10)
	dbx.SetMaxIdleConns(10)

	db.DbxInsert(dbx, responseBodyArray)

	query := "select * from jarujaru"
	db.DbxSelect(dbx, query)

	// call := service.CommentThreads.List([]string{"snippet"}).
	// 	VideoId(videoid).MaxResults(50)
	// response, err := call.Do()
	// handleError(err, "")
	// fmt.Println(response.Items[0])
	// if response.NextPageToken != "" {
	// 	nextPageToken := response.NextPageToken
	// 	fmt.Printf("nextPageToken: %v\n", nextPageToken)
	// index := 1
	// for {
	// 	call := service.CommentThreads.List([]string{"snippet"}).
	// 		VideoId(videoid).
	// 		PageToken(nextPageToken)
	// 	response, err := call.Do()
	// 	handleError(err, "")
	// 	// fmt.Println(response.Items[0])

	// 	// for _, item := range response.Items {
	// 	// 	text := item.Snippet.TopLevelComment.Snippet.TextDisplay
	// 	// 	likeCnt := item.Snippet.TopLevelComment.Snippet.LikeCount
	// 	// 	time := item.Snippet.TopLevelComment.Snippet.PublishedAt
	// 	// 	fmt.Printf("Text: %v\n", text)
	// 	// 	fmt.Printf("likeCnt: %v\n", likeCnt)
	// 	// 	fmt.Printf("time: %v\n", time)
	// 	// }
	// 	nextPageToken := response.NextPageToken
	// 	if nextPageToken == "" {
	// 		break
	// 	}
	// 	fmt.Printf("nextPageToken: %v\n", nextPageToken)
	// 	index += 1
	// 	if index == 3 {
	// 		break
	// 	}

	//	}

	// }

}

//linkカウントも取得する ok
//綺麗にする　ok

//動画のURLを書く
//取得したコメントをデータベースに保存する
//DBに保存する

//時間を曜日で分ける
//月で分ける
//時間で分ける

//より簡単にする

//typeを分類する方法を考える
//デザインを考える。
//動画を毎日保存する。

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
