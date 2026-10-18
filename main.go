package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "", "путь к файлу для анализа")
	topN := flag.Int("top", 10, "количество слов в топе")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("использование: textanalyzer -file <путь>")
		os.Exit(1)
	}

	f, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Print(err.Error()) // TODO: ЗАГУГЛИ
        os.Exit(1)
	}

	countentFile := string(f)

	lines, err := CountLines(countentFile)
	if err != nil {
		fmt.Println(err.Error())
        os.Exit(1)
	}

	words, err := CountWords(countentFile)
	if err != nil {
		fmt.Println(err.Error())
        os.Exit(1)
	}

	chars, err := CountChars(countentFile)
	if err != nil {
		fmt.Println(err.Error())
        os.Exit(1)
	}

	wordMap, err := WordFrequency(countentFile)
	if err != nil {
		fmt.Println(err.Error())
        os.Exit(1)
	}

	sliceWordCount, err := TopWords(wordMap, *topN)
	if err != nil {
		fmt.Println(err.Error())
        os.Exit(1)
	}

	fmt.Printf("Количество строк в тексте: %d\n", lines)
	fmt.Printf("Количество слов в тексте: %d\n", words)
	fmt.Printf("Количество символов рун в тексте: %d\n", chars)
	fmt.Println("Результат — слайс пар (слово, количество):", sliceWordCount)
}
