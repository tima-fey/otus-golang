package main

import (
	"fmt"
	"strings"
	"unicode"
)

func checkLetters(letter, nextLetter rune) string {
	if unicode.IsDigit(nextLetter) {
		return strings.Repeat(string(letter), int(nextLetter-'0'))
	}
	return string(letter)
}

func unpack(word string) string {
	tempWord := []rune(word)
	answer := ""
	for num := 0; num < len(tempWord); num++ {
		if unicode.IsLetter(tempWord[num]) {
			if num+1 < len(tempWord) {
				answer += checkLetters(tempWord[num], tempWord[num+1])
			} else {
				answer += string(tempWord[num])
			}
		}
		if string(tempWord[num]) == `\` && num+1 < len(tempWord) {
			num++
			if num+1 < len(tempWord) {
				answer += checkLetters(tempWord[num], tempWord[num+1])
			} else {
				answer += string(tempWord[num])
			}
		}
	}
	return answer
}
func main() {
	cases := [][]string{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}
	for i, aCase := range cases {
		if unpack(aCase[0]) != aCase[1] {
			fmt.Printf("Error %v \n", i)
		} else {
			fmt.Printf("Ok %v \n", i)
		}
	}
}
