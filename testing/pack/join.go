package pack

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1] // if the slice has just two items, just join them with "and"
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ") // Join every phrase except the last with commas
		result += ", and "                                     //Insert the word "and" before the last phrase
		result += phrases[len(phrases)-1]                      // add the last phrase
		return result
	}
}
