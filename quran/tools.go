package quran

import "unicode/utf8"

// CanTweet check a string can be tweet or not by checking string length
func CanTweet(s string) bool {
	if utf8.RuneCountInString(s) > maxTweetLen {
		return false
	}

	return true
}
