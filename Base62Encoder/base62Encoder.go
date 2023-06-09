package base62Encoder

import "strings"

// Map to store 62 possible characters
var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var charList = []rune(str)

func Encode(num int64) string {
	var sb strings.Builder
	for i := num; i > 0; {
		sb.WriteString(string(charList[i%62]))
		i = i / 62
	}
	return Reverse(sb.String())
}

func Decode(input string) int64 {
	var num int64 = 0
	var base int64 = 62
	for _, r := range input {
		num = num*base + int64(strings.IndexRune(str, r))
	}
	return num
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
