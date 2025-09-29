package vector

import (
	"errors"
	"testing"
)

type TestCase struct {
	name     string
	vec1     []CompressedPair
	vec2     []CompressedPair
	expected int
	err      error
}

func TestDotProduct(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "Простой случай (1-в-1)",
			vec1:     []CompressedPair{{3, 2}, {7, 1}},
			vec2:     []CompressedPair{{5, 2}, {10, 1}},
			expected: 100, // (3*5)*2 + (7*10)*1 = 30 + 70
			err:      nil,
		},
		{
			name:     "Сложный случай (пересечение)",
			vec1:     []CompressedPair{{1, 2}, {2, 3}, {3, 1}},
			vec2:     []CompressedPair{{4, 3}, {5, 3}},
			expected: 51, // (1*4)*2 + (2*4)*1 + (2*5)*2 + (3*5)*1 = 8 + 8 + 20 + 15
			err:      nil,
		},
		{
			name:     "Случай 'многие-к-одному'",
			vec1:     []CompressedPair{{10, 1}, {20, 2}},
			vec2:     []CompressedPair{{5, 3}},
			expected: 250, // (10*5)*1 + (20*5)*2 = 50 + 200
			err:      nil,
		},
		{
			name:     "Ошибка: вектора разной длины",
			vec1:     []CompressedPair{{1, 4}},
			vec2:     []CompressedPair{{1, 2}, {10, 1}},
			expected: 0,
			err:      ErrMismatchedLengths,
		},
		{
			name:     "Крайний случай (пустые векторы)",
			vec1:     []CompressedPair{},
			vec2:     []CompressedPair{},
			expected: 0,
			err:      nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := DotProduct(testCase.vec1, testCase.vec2)

			if !errors.Is(err, testCase.err) {
				t.Errorf("Получена ошибка '%v', ожидалась '%v'", err, testCase.err)
				return
			}

			if got != testCase.expected {
				t.Errorf("Получено: %d, ожидалось: %d", got, testCase.expected)
			}
		})
	}
}
