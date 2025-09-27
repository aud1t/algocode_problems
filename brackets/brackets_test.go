package brackets

import "testing"

type TestCase struct {
	name     string
	input    string
	expected int
}

var testCases = []TestCase{
	{
		name:     "Только одна невалидная скобка",
		input:    "((){}[))",
		expected: 6,
	},
	{
		name:     "Нет невилидных скобок",
		input:    "{}()[](())",
		expected: -1,
	},
	{
		name:     "Несколько невалидных скобок",
		input:    "[(){)()[)]",
		expected: -1,
	},
	{
		name:     "Пустая строка",
		input:    "",
		expected: -1,
	},
	{
		name: "Нечетное количество скобок",
		input: "(()",
		expected: -1,
	},
}

func TestFindSingleInvalidIndex(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := FindSingleInvalidIndex(testCase.input)
			if got != testCase.expected {
				t.Errorf("получено: %d, ожидалось: %d", got, testCase.expected)
			}
		})
	}
}
