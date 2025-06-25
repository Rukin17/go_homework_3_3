package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()
	words := strings.Fields(input)
	result := checkWordCount(words)
	showResult(result)

}

func checkWordCount(s []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range s {
		word = strings.ToLower(word)
		wordCount[word]++
	}
	return wordCount
}

func showResult(wordCount map[string]int) {
	for key, value := range wordCount {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите строку: ")
	input, _ := reader.ReadString('\n')
	return input
}
