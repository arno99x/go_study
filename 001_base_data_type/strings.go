package main

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	fmt.Println("init strings.go")
}
func main() {
	testMaxNoRepateStrLeng()
}

func testStrForRune() {
	s := "yes我爱慕课网."
	fmt.Println(s)

	for _, v := range []byte(s) {
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ,(%d)", ch, size)
	}

	fmt.Println()
	runes := []rune(s)
	for _, v := range runes {
		fmt.Printf("%c ", v)
	}

	fmt.Println()
	for _, v := range s {
		fmt.Printf("%c ", v)
	}
}

func testMaxNoRepateStrLeng() {
	fmt.Println(maxNoRepateStrLeng("abcabcadfw"))
	fmt.Println(maxNoRepateStrLeng("aaaaaaaa"))
	fmt.Println(maxNoRepateStrLeng("twaacaawtc"))
	fmt.Println(maxNoRepateStrLeng("一二三二一"))
}

func maxNoRepateStrLeng(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range s {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
