package brackets

// FindSingleInvalidIndexSimple находит индекс одной скобки, которую нужно поменять,
// чтобы сделать последовательность валидной. Функция возвращает -1, если
// исправление невозможно, ошибок нет или их больше одной.
//
// Алгоритм проходит по строке слева направо, чтобы найти одну лишнюю закрывающую скобку ).
// Если такая ошибка не найдена, но в строке остались незакрытые (, запускается второй проход.
// Этот второй проход идёт уже справа налево, чтобы точно найти ту самую открывающую скобку (, которой не хватило пары.
//
// Сложность по времени: O(N)
//
// Сложность по памяти: O(1)
func FindSingleInvalidIndexSimple(str string) int {
	if len(str)%2 == 1 {
		return -1
	}

	openBrackets := 0
	problemIndex := -1

	for i, char := range str {
		if char == '(' {
			openBrackets++
		} else if openBrackets > 0 {
			openBrackets--
		} else {
			if problemIndex != -1 {
				return -1
			}
			problemIndex = i
			openBrackets++
		}
	}

	if openBrackets == 0 {
		return problemIndex
	}

	if openBrackets == 2 && problemIndex == -1 {
		return findUnmatchedOpenerIndex(str)
	}

	return -1
}

func findUnmatchedOpenerIndex(str string) int {
	closeCount := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ')' {
			closeCount++
		} else if closeCount > 0 {
			closeCount--
		} else {
			return i
		}
	}
	return -1
}
