package main

import (
	"fmt"
	"strings"
	"unicode"
)

func cleanWord(word string) string{
	var clean strings.Builder
	for _, letter := range word{
		if unicode.IsDigit(letter) || unicode.IsLetter(letter){
			clean.WriteRune(letter)
		}
	}
	return clean.String()
}

func getTopWords(wordMap map[string]int, n int) []string{
	var result []string
	copy_map := make(map[string]int)
	for k, v := range wordMap{
		copy_map[k] = v
	}
	for i := 0; i < n; i++{
		max_value := 0
		var topWord string
		for key, value := range copy_map{
			if value > max_value{
				max_value = value
				topWord = key
			}
		}
		result = append(result, topWord)
		delete(copy_map, topWord)
	}
	return result
}

func AnalyzeText(text string){
	words := strings.Fields(text)
	c_words := len(words)
	words_count := make(map[string]int)
	for _, word := range words{
		word = cleanWord(strings.ToLower(word))
		words_count[word]++
	}
	max_count := 0
	var most_Frequent string
	for word, value := range words_count{
		if value > max_count{
			max_count = value
			most_Frequent = word
		}
	}
	fmt.Printf("Количество слов: %d \n", c_words)
	fmt.Printf("Количество уникальных слов: %d \n", len(words_count))
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз) \n", most_Frequent, max_count)
	topWords := getTopWords(words_count, 5)
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range topWords{
		fmt.Printf("\"%s\": %d раз \n", word, words_count[word])
	}
}

func main() {
	AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю. Море! Я море!!! ЛЮБЛЮ МОРЕ")
}