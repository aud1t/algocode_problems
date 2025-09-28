package pangram

import (
	"testing"
)

type inputData struct {
	alphabet string
	str      string
}

type TestCase struct {
	name     string
	input    inputData
	expected string
}

func TestFindShortest(t *testing.T) {
	var testCases = []TestCase{
		{
			name:     "Стандартный случай",
			input:    inputData{alphabet: "abc", str: "dfagabkaceb"},
			expected: "bkac",
		},
		{
			name:     "Алфавит из одного элемента",
			input:    inputData{alphabet: "a", str: "dfagabkaceb"},
			expected: "a",
		},
		{
			name:     "Панграмма размером со всю строку",
			input:    inputData{alphabet: "abc", str: "addcffb"},
			expected: "addcffb",
		},
		{
			name:     "Уменьшающаяся панграмма",
			input:    inputData{alphabet: "abc", str: "caaaaaaaabbbcccc"},
			expected: "abbbc",
		},
		{
			name:     "Панграмма не найдена",
			input:    inputData{alphabet: "abc", str: "dfgbkceb"},
			expected: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := FindShortest(testCase.input.alphabet, testCase.input.str)

			if testCase.expected == "" {
				if got != "" {
					t.Errorf("ожидалась пустая строка, но получено: '%s'", got)
				}
				return
			}

			if len(got) != len(testCase.expected) {
				t.Errorf("ожидалась панграмма длиной %d, но получена '%s' (длина %d)",
					len(testCase.expected), got, len(got))
				return
			}

			if !isPangram(got, testCase.input.alphabet) {
				t.Errorf("строка '%s' не является панграммой для алфавита '%s'", got, testCase.input.alphabet)
			}
		})
	}
}

func isPangram(s string, alphabet string) bool {
	alphabetSet := make(map[rune]bool)
	for _, r := range alphabet {
		alphabetSet[r] = true
	}
	foundSet := make(map[rune]bool)
	for _, r := range s {
		if alphabetSet[r] {
			foundSet[r] = true
		}
	}
	return len(foundSet) == len(alphabetSet)
}