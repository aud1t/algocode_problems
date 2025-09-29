package strutil

import "testing"

type TestCase struct {
	name     string
	text     string
	pattern  string
	expected string
}

func TestRemoveOccurrences(t *testing.T) {
	var testCases = []TestCase{
		{
			name:     "Множественное вложения pattern-а",
			text:     "axxxxyyyyb",
			pattern:  "xy",
			expected: "ab",
		},
		{
			name:     "Pattern нет в строке",
			text:     "qwertyu",
			pattern:  "dfg",
			expected: "qwertyu",
		},
		{
			name:     "Pattern совпадает с строкой",
			text:     "qwertyu",
			pattern:  "qwertyu",
			expected: "",
		},
		{
			name:     "Pattern длиннее, чем сама строка",
			text:     "qwerty",
			pattern:  "qwertyu",
			expected: "qwerty",
		},
		{
			name:     "Строка состоит только из pattern-a",
			text:     "aaaaabcbcbcbcbc",
			pattern:  "abc",
			expected: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := RemoveOccurrences(testCase.text, testCase.pattern)

			if got != testCase.expected {
				t.Errorf("Получено: %s, ожидалось: %s", got, testCase.expected)
			}
		})
	}
}
