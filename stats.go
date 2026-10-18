package main

import (
	"errors"
	"slices"
	"strings"
)

// Сентинал ошибки
var ErrEmptyString = errors.New("пустая строка")
var ErrQuantityNWords = errors.New("возвращаемое количество слов в слайсе равно нулю или меньше нуля")

// CountLines считает количество строк в тексте.
func CountLines(text string) (int, error) {
	
	if len(text) == 0 {
		return 0, ErrEmptyString
	}
	return strings.Count(text, "\n"), nil
}

// CountWords считает количество слов в тексте.
func CountWords(text string) (int, error) {
	if len(text) == 0 {
		return 0, ErrEmptyString
	}
	return len(strings.Fields(text)), nil
}

// CountChars считает количество символов (рун) в тексте.
func CountChars(text string) (int, error) {
	if len(text) == 0 {
		return 0, ErrEmptyString
	}
	result := 0
	for index := 0; index < len([]rune(text)); index++ {
		result++
	}
	return result, nil
}

// WordFrequency возвращает карту слово→количество.
// Все слова приведены к нижнему регистру.
func WordFrequency(text string) (map[string]int, error) {
	if len(text) == 0 {
		return map[string]int{}, ErrEmptyString
	}
	countWords := map[string]int{}
	wordsSlice := strings.Fields(strings.ToLower(text))
	for _, v := range wordsSlice {
		if _, ok := countWords[v]; !ok {
			countWords[v] = 1
		} else {
			countWords[v] += 1
		}
	}
	return countWords, nil
}

// WordCount — пара слово + количество для сортировки.
type WordCount struct {
	Word  string
	Count int
}

// TopWords возвращает топ-N слов по частоте.
// Результат — слайс пар (слово, количество), отсортированный по убыванию.
func TopWords(freq map[string]int, n int) ([]WordCount, error) {
	result := []WordCount{}
	if n <= 0 {
		return result, ErrQuantityNWords
	}
	for key, value := range freq {
		result = append(result, WordCount{Word: key, Count: value})
	}
	slices.SortFunc(result, func(a, b WordCount) int {
		return b.Count - a.Count
	})
	if len(result) < n {
		return result[:len(result)], nil
	} else {
		return result[:n], nil
	}
}
