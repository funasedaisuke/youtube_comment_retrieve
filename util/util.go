package util

import "time"

func TimeToJapan(timeStrinig string, jst *time.Location) time.Time {
	utctime, _ := time.Parse(time.RFC3339, timeStrinig)
	jptime := utctime.In(jst)
	return jptime
}
