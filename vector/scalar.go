package vector

import "errors"

var ErrMismatchedLengths = errors.New("скалярное произведение не определено для векторов разной длины")

// CompressedPair представляет пару (значение, количество) в сжатом векторе.
type CompressedPair struct {
	Value int
	Count int
}

// DotProduct вычисляет скалярное произведение двух векторов, заданных в сжатой форме.
//
// Функция принимает два вектора в формате среза структур CompressedPair, где
// каждая пара {Value, Count} представляет серию одинаковых чисел. Если
// разжатые длины векторов не совпадают, скалярное произведение не определено,
// и функция возвращает ненулевую ошибку.
//
// Идея решения — использовать два указателя для одновременного прохода по обоим
// сжатым векторам. На каждом шаге алгоритм обрабатывает минимальный общий "отрезок"
// и сдвигает тот указатель, чей отрезок был полностью исчерпан.
//
// Сложность по времени: O(L1 + L2), где L1 и L2 — длины сжатых векторов.
// Сложность по памяти: O(1).
func DotProduct(vec1, vec2 []CompressedPair) (int, error) {
	if len(vec1) == 0 && len(vec2) == 0 {
		return 0, nil
	}

	if uncompressedLength(vec1) != uncompressedLength(vec2) {
		return 0, ErrMismatchedLengths
	}

	var (
		product    = 0
		p1, p2     = 0, 0
		rem1, rem2 = 0, 0
	)

	for p1 < len(vec1) && p2 < len(vec2) {
		if rem1 == 0 {
			rem1 = vec1[p1].Count
		}
		if rem2 == 0 {
			rem2 = vec2[p2].Count
		}

		chunkSize := min(rem1, rem2)

		product += chunkSize * vec1[p1].Value * vec2[p2].Value

		rem1 -= chunkSize
		rem2 -= chunkSize

		if rem1 == 0 {
			p1++
		}
		if rem2 == 0 {
			p2++
		}
	}

	return product, nil
}

func uncompressedLength(vector []CompressedPair) int {
	sum := 0
	for _, cp := range vector {
		sum += cp.Count
	}
	return sum
}
