package pangram

// FindShortest находит в строке str панграмму минимальной длины для заданного алфавита alphabet.
// Панграмма — это подстрока, содержащая все символы из алфавита.
// Если панграмма не найдена, возвращает пустую строку.
//
// Идея решения — использовать метод скользящего окна. Окно расширяется вправо,
// пока не будет содержать все символы алфавита. Как только панграмма найдена,
// окно сжимается слева, на каждом шаге обновляя найденную панграмму минимальной длины.
//
// Сложность по времени: O(N + M), где N — длина алфавита, а M — длина строки.
// Сложность по памяти: O(N), так как для хранения символов алфавита
// используются хеш-таблицы, размер которых не превышает N.
func FindShortest(alphabet string, str string) string {
	if len(alphabet) == 0 || len(str) == 0 {
		return ""
	}

	alphabetSet := make(map[rune]bool)
	for _, char := range alphabet {
		alphabetSet[char] = true
	}

	runes := []rune(str)
	left, right := 0, 0

	windowCounts := make(map[rune]int)
	required := len(alphabetSet)
	formed := 0

	minLength := -1
	resultStart, resultEnd := 0, 0

	for right < len(runes) {
		char := runes[right]

		if alphabetSet[char] {
			windowCounts[char]++
			if windowCounts[char] == 1 {
				formed++
			}
		}

		for left <= right && formed == required {
			currentLength := right - left + 1

			if minLength == -1 || currentLength < minLength {
				minLength = currentLength
				resultStart = left
				resultEnd = right
			}

			leftChar := runes[left]
			if alphabetSet[leftChar] {
				windowCounts[leftChar]--
				if windowCounts[leftChar] == 0 {
					formed--
				}
			}
			left++
		}

		right++
	}

	if minLength == -1 {
		return ""
	}

	return string(runes[resultStart : resultEnd+1])
}
