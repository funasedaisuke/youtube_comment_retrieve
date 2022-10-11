package main

import (
	"fmt"
	//"flag"

	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// var (
// 	query      = flag.String("query", "Google", "Search term")
// 	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
// )

const developerKey = ""

func createService() (service *youtube.Service) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	return service
}

func getCommentThread(nextPageToken, videoID string, service *youtube.Service) {
	// Make the API call to YouTube.
	call := service.CommentThreads.List([]string{"snippet"}).
		VideoId(videoID)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(response.Items[0])

	for _, item := range response.Items {
		text := item.Snippet.TopLevelComment.Snippet.TextDisplay
		likeCnt := item.Snippet.TopLevelComment.Snippet.LikeCount
		time := item.Snippet.TopLevelComment.Snippet.PublishedAt
		fmt.Printf("Text: %v\n", text)
		fmt.Printf("likeCnt: %v\n", likeCnt)
		fmt.Printf("time: %v\n", time)
	}
}

func getVideoIDs(uploadsID string, service *youtube.Service) (videoIDs []string) {
	playlistsCall := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(uploadsID).MaxResults(50)

	playlistsResponse, err := playlistsCall.Do()
	if err != nil {
		log.Fatalf("Error call YouTube API: %v", err)
	}

	for _, item := range playlistsResponse.Items {
		videoIDs = append(videoIDs, item.ContentDetails.VideoId)
	}

	if playlistsResponse.NextPageToken != "" {
		nextPageToken := playlistsResponse.NextPageToken
		for {
			nextCall := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(uploadsID).PageToken(nextPageToken).MaxResults(50)
			nextResponse, err := nextCall.Do()
			if err != nil {
				log.Fatalf("Error call YouTube API: %v", err)
			}
			for _, nextItem := range nextResponse.Items {
				videoIDs = append(videoIDs, nextItem.ContentDetails.VideoId)
			}
			nextPageToken = nextResponse.NextPageToken
			if nextPageToken == "" {
				break
			}
		}
	}

	return videoIDs
}

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

func getVideoDatas(videoIDs []string, service *youtube.Service) (videoDatas []VideoData) {
	for _, videoID := range videoIDs {
		call := service.Videos.List([]string{"snippet", "statistics"}).Id(videoID)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error call YouTube API: %v", err)
		}

		item := response.Items[0]
		videoData := VideoData{
			VideoID:       videoID,
			Title:         item.Snippet.Title,
			ViewCount:     item.Statistics.ViewCount,
			LikeCount:     item.Statistics.LikeCount,
			DislikeCount:  item.Statistics.DislikeCount, // ※oauth認証じゃないと取得できない
			FavoriteCount: item.Statistics.FavoriteCount,
			CommentCount:  item.Statistics.CommentCount,
			PublishedAt:   item.Snippet.PublishedAt,
		}
		videoDatas = append(videoDatas, videoData)
	}
	return videoDatas
}

func main() {
	// flag.Parse()
	videoid := "9BBoWAANXHo"
	service := createService()
	call := service.CommentThreads.List([]string{"snippet"}).
		VideoId(videoid).MaxResults(50)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(response.Items[0])
	if response.NextPageToken != "" {
		nextPageToken := response.NextPageToken
		fmt.Printf("nextPageToken: %v\n", nextPageToken)
		index := 1
		for {
			call := service.CommentThreads.List([]string{"snippet"}).
				VideoId(videoid).
				PageToken(nextPageToken)
			response, err := call.Do()
			handleError(err, "")
			// fmt.Println(response.Items[0])

			// for _, item := range response.Items {
			// 	text := item.Snippet.TopLevelComment.Snippet.TextDisplay
			// 	likeCnt := item.Snippet.TopLevelComment.Snippet.LikeCount
			// 	time := item.Snippet.TopLevelComment.Snippet.PublishedAt
			// 	fmt.Printf("Text: %v\n", text)
			// 	fmt.Printf("likeCnt: %v\n", likeCnt)
			// 	fmt.Printf("time: %v\n", time)
			// }
			nextPageToken := response.NextPageToken
			if nextPageToken == "" {
				break
			}
			fmt.Printf("nextPageToken: %v\n", nextPageToken)
			index += 1
			if index == 3 {
				break
			}

		}

	}

}

//linkカウントも取得する
//綺麗にする

//DBに保存する
//より簡単にする

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
