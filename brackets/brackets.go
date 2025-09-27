package brackets

// findSingleInvalidIndex находит индекс одной скобки, которую нужно поменять,
// чтобы сделать последовательность валидной. Функция возвращает -1, если
// исправление невозможно, ошибок нет или их больше одной.
//
// Идея решения заключается в анализе строки с использованием стека, где хранится символы и их индексы.
// При встрече закрывающей скобки проверяется соответствие с вершиной стека
// или если стек пуст, тогда она меняется на откривающий и добавляется в стек.
// Все ошибки (несоответствия, лишние закрывающие) фиксируются, и если в итоге найдена
// ровно одна такая ошибка и не осталось незакрытых скобок, возвращается её индекс.
// Или если ошибок не найдена но в стеке ровно два открывающих скобки, например ((
//
// Сложность по времени: O(N), где N - длина входной строки,
// так как каждый символ обрабатывается один раз.
//
// Сложность по памяти: O(N) в худшем случае, когда вся строка
// состоит из открывающих скобок, которые помещаются в стек.
func findSingleInvalidIndex(str string) int {
	openingPairs := map[rune]rune{'{': '}', '(': ')', '[': ']'}
	closingPairs := map[rune]rune{'}': '{', ')': '(', ']': '['}
	allBrackets := map[rune]bool{
		'(': true, ')': true,
		'[': true, ']': true,
		'{': true, '}': true,
	}

	var stack []bracketInfo
	var problemIndices []int

	for i, char := range str {
		if _, isBracket := allBrackets[char]; !isBracket {
			continue
		}

		if _, ok := openingPairs[char]; ok {
			stack = append(stack, bracketInfo{symbol: char, index: i})
			continue
		}

		if len(stack) == 0 {
			// Если встречаем ), то в стек добавляем ( и его индекс добавим в проблемные
			stack = append(stack, bracketInfo{symbol: closingPairs[char], index: i})
			problemIndices = append(problemIndices, i)
		} else {
			lastOpen := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if openingPairs[lastOpen.symbol] != char {
				problemIndices = append(problemIndices, i)
			}
		}

		if len(problemIndices) > 1 {
			return -1
		}
	}

	if len(problemIndices) == 1 {
		if len(stack) == 0 {
			return problemIndices[0]
		} else {
			return -1
		}
	}

	// Рассмотрим только случай '((', т.к. случай '))' в основном цикле поменяться на ()
	if len(stack) == 2 && stack[0].symbol == stack[1].symbol {
		return stack[1].index
	}

	return -1
}

type bracketInfo struct {
	symbol rune
	index  int
}
