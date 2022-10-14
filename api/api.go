package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"main/db"
	"main/util"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func getCredential() string {
	b, err := ioutil.ReadFile("client_secrets.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config.ClientID
}

func createService(developerKey string) (service *youtube.Service) {
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
	util.HandleError(err, "")
	// body, _ := response.MarshalJSON()

	nextpagetoken = response.NextPageToken

	jst, _ := time.LoadLocation("Asia/Tokyo")
	for _, item := range response.Items {

		commentInstance := db.Comment{}
		commentInstance.Comment = strings.Replace(item.Snippet.TopLevelComment.Snippet.TextDisplay, "\"", "", -1)
		commentInstance.LikeCnt = int(item.Snippet.TopLevelComment.Snippet.LikeCount)
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

func ApiClientCall() []db.Comment {
	clientID := getCredential()
	videoid := "HK8CzJm8gdM"
	service := createService(clientID)
	nextPageToken := ""
	var responseBodyArray []db.Comment
	for {
		responseBodyArray, nextPageToken = getCommentThread(nextPageToken, videoid, service, responseBodyArray)
		fmt.Println(len(responseBodyArray))
		if nextPageToken == "" {
			break
		}
	}
	return responseBodyArray

}
