package brackets

import "testing"

var moreTestCases = []TestCase{
	{
		name:     "Две одинаковые закрывающие скобки",
		input:    "{}]()]",
		expected: 2,
	},
	{
		name:     "Две одинаковые открывающие скобки",
		input:    "([]{}([]",
		expected: 5,
	},
	{
		name:     "Две разные открывающие скобки",
		input:    "()(()[()",
		expected: 5,
	},
	{
		name:     "Две разные закрывающие скобки",
		input:    "])",
		expected: 0,
	},
	{
		name:     "Много проблемных индексов",
		input:    ")(",
		expected: -1,
	},
}

func TestFindSingleInvalidIndex_More(t *testing.T) {
	for _, testCase := range moreTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := FindSingleInvalidIndex(testCase.input)
			if got != testCase.expected {
				t.Errorf("получено: %d, ожидалось: %d", got, testCase.expected)
			}
		})
	}
}
