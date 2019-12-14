package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word    string
	counter int
}
type wordsCounter []wordCount

func (a wordsCounter) Less(i, j int) bool { return a[i].counter > a[j].counter }
func (a wordsCounter) Len() int           { return len(a) }
func (a wordsCounter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func prepareStrings(sentence string) []string {
	space := regexp.MustCompile(`\s+`)
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	sentence = strings.ToLower(sentence)
	sentence = strings.TrimSpace(sentence)
	sentence = reg.ReplaceAllString(sentence, "")
	sentence = space.ReplaceAllString(sentence, " ")

	return strings.Split(sentence, " ")
}
func countWords(splitedSentence []string) []wordCount {
	allCounters := []wordCount{}
	counter := make(map[string]int)
	for _, subString := range splitedSentence {
		counter[subString]++
	}
	for word, count := range counter {
		allCounters = append(allCounters, wordCount{word: word, counter: count})
	}
	sort.Sort(wordsCounter(allCounters))
	return allCounters
}
func getTopCommon(sentence string, count int) []wordCount {
	splitedSentence := prepareStrings(sentence)
	answer := countWords(splitedSentence)
	return answer[0:count]
}

func main() {
	wordsNumber := 2
	answer := getTopCommon("test string test, b b b", wordsNumber)
	fmt.Println(answer)
}
