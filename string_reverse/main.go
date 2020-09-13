package main

import (
	"fmt"
	"unicode/utf8"
)

// My implementation
func Reverse(str string) string {
	length := len(str)
	runes := []rune(str)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

// implementation from the internet
func Reverse2(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func main() {
	str := "Original string which should be reversed"

	fmt.Println(Reverse(str))
	fmt.Println(Reverse2(str))
}
