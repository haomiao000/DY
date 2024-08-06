package pkg

import "fmt"

func WithUserID(userID int64) string {
	return fmt.Sprintf(" user_id = %d ", userID)
}

func WithVideoID(videoID int64) string {
	return fmt.Sprintf(" video_id = %d ", videoID)
}

func assembleSQL(args ...string) string {
	res := ""
	for idx, v := range args {
		if idx == 0 {
			res += v
		} else {
			res += (" and " + v)
		}
	}
	return res
}
