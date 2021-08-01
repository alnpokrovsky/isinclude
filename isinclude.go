/*
 * Дан упорядоченный массив чисел размером N
 * Нужно реализовать алгоритм поиска вхождения упорядоченного
 * подмассива размера M, где M << N
 */
package isinclude

// IsInclude проверяет, что array включает subarray
func IsInclude(array []int, subarray []int) bool {
	if len(subarray) == 0 {
		return true
	}

	firstPossible := findPosBin(array, subarray[0])
	if firstPossible == -1 {
		return false
	}

	cf1 := countDuplicatesOfFirst(array[firstPossible:])
	cf2 := countDuplicatesOfFirst(subarray)
	if cf1 < cf2 {
		return false
	} else {
		firstPossible += (cf1 - cf2)
	}

	return hasEqual(array[firstPossible:], subarray)
}

// findPosBin выполняет бинарный поиск в отсортированном массиве
// если элемент не найден возвращает -1, иначе позицию первого вхождения
func findPosBin(array []int, need int) int {
	lowPos := 0               // первый индекс
	highPos := len(array) - 1 // последний индекс

	if (array[lowPos] > need) || (array[highPos] < need) {
		return -1 // нужное значение не в диапазоне данных
	}

	for lowPos < highPos {
		mid := (lowPos + highPos) / 2

		if array[mid] < need {
			lowPos = mid + 1
		} else {
			highPos = mid
		}
	}

	if array[lowPos] == need {
		return lowPos
	} else {
		return -1
	}
}

// countDuplicatesOfFirst подсчитывает количество повторений
// первого элемента в упорядоченном массиве
func countDuplicatesOfFirst(array []int) int {
	count := 1
	for array[0] == array[count] {
		count++
	}
	return count
}

// hasEqual проверяет на равенство subarray и начало array
// (array может быть больше subarray в отличие от isEqual)
func hasEqual(array []int, subarray []int) bool {
	if len(array) < len(subarray) {
		return false
	}

	for i, v := range subarray {
		if v != array[i] {
			return false
		}
	}
	return true
}
