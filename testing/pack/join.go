package pack

import "strings"

func JoinWithCommas(phrases []string) string {
	result := strings.Join(phrases[:len(phrases)-1], ", ") // Join every phrase except the last with commas
	result += ", and "                                     //Insert the word "and" before the last phrase
	result += phrases[len(phrases)-1]                      // add the last phrase
	return result
}
