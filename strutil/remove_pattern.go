package strutil

// RemoveOccurrences удаляет все вхождения подстроки pattern из строки text,
// следуя правилу "найти самое левое вхождение подстроки и удалить его" до тех пор,
// пока вхождений не останется.
//
// Идея решения заключается в построении итоговой строки за один проход с
// использованием стека. Символы добавляются в буфер (builder), и после 
// каждого добавления проверяется, не заканчивается ли буфер на искомый шаблон.
// Если совпадение найдено, из конца буфера обрезается подстрока длиной patternLen.
//
// Сложность по времени: O(N*M), где N — длина text, а M — длина part.
// В худшем случае на каждом из N шагов будет выполняться сравнение длиной M.
//
// Сложность по памяти: O(N), в худшем случае, если вхождений
// нет, то строитель вырастет до размера исходной строки.
func RemoveOccurrences(text string, pattern string) string {
	builder := []rune{}
	patternRunes := []rune(pattern)
	patternLen := len(patternRunes)

	for _, char := range text {
		builder = append(builder, char)

		if endsWith(builder, patternRunes) {
			builderLen := len(builder)
			builder = builder[:builderLen-patternLen]
		}
	}

	return string(builder)
}

func endsWith(a []rune, b []rune) bool {
	aLen := len(a)
	bLen := len(b)

	if aLen < bLen {
		return false
	}

	for i := range bLen {
		if a[aLen-bLen+i] != b[i] {
			return false
		}
	}
	return true
}
