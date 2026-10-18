package main

import (
	"testing"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"пустая строка", "", 0},
		{"одна строка", "awdawd", 0},
		{"несколько строк", "awddd \n dawd \n dawd", 2},
		{"строка с \\n в конце", "dawdaw\n", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := CountLines(tt.input)
			if result != tt.expected {
				t.Errorf("CountLines(%s) = %d, хотели %d\n", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"пустая строка", "", 0},
		{"одно слово", "привет", 1},
		{"несколько слов", "привет мир", 2},
		{"лишние пробелы", "  привет   мир  ", 2},
		{"только пробелы", "   ", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := CountWords(tt.input)
			if result != tt.expected {
				t.Errorf("CountWords(%s) = %d, хотели %d\n", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"ASCII", "hello", 5},
		{"кириллица", "привет", 6},
		{"эмоджи", "привет 🚀", 8},
		{"пустая строка", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := CountChars(tt.input)
			if result != tt.expected {
				t.Errorf("CountChars(%s) = %d, хотели %d\n", tt.input, result, tt.expected)
			}
		})
	}
}

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			"пустой текст",
			"",
			map[string]int{},
		},
		{
			"обычный текст",
			"go is great and go is fast",
			map[string]int{"go": 2, "is": 2, "great": 1, "and": 1, "fast": 1},
		},
		{
			"повторы",
			"ok ok ok ok",
			map[string]int{"ok": 4},
		},
		{
			"разный регистр",
			"Ok Ok Kk Ok",
			map[string]int{"ok": 3, "kk": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := WordFrequency(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("WordFrequency(%s). Несовпадение длин слайсов\n", tt.input)
			}
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("WordFrequency(%s). Значения слайсов не совпадают по ключу\n", tt.input)
				}
			}
		})
	}
}

func TestTopWords(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		n        int
		expected []WordCount
	}{
		{
			"топ-2",
			map[string]int{"go": 5, "is": 3, "and": 1},
			2,
			[]WordCount{{"go", 5}, {"is", 3}},
		},
		{
			"n больше слов",
			map[string]int{"go": 5, "is": 3},
			10,
			[]WordCount{{"go", 5}, {"is", 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := TopWords(tt.input, tt.n)
			if len(result) != len(tt.expected) {
				t.Errorf("TopWords(%v, %d). Несовпадение длин слайсов\n", tt.input, tt.n)
				return
			}
			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("TopWords(%v, %d). Значения слайсов не совпадают по ключу\n", tt.input, tt.n)
				}
			}
		})
	}
}
