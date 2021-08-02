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

	firstPossible := lowerBound(array, subarray[0])
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

// lowerBound выполняет бинарный поиск в отсортированном массиве
// если элемент не найден возвращает -1, иначе позицию первого вхождения
func lowerBound(array []int, need int) int {
	lowPos := 0               // первый индекс
	highPos := len(array) - 1 // последний индекс

	// определяем функцию сравнения для упорядоченного
	// по возрастанию или по убыванию массива
	var compare func(a int, b int) bool
	if array[lowPos] < array[highPos] {
		if (array[lowPos] > need) || (array[highPos] < need) {
			return -1 // нужное значение не в диапазоне данных
		}
		compare = func(a int, b int) bool { return a < b }
	} else {
		if (array[highPos] > need) || (array[lowPos] < need) {
			return -1 // нужное значение не в диапазоне данных
		}
		compare = func(a int, b int) bool { return a > b }
	}

	// бинарный поиск
	for lowPos < highPos {
		mid := (lowPos + highPos) / 2

		if compare(array[mid], need) {
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
	for count := 1; count < len(array); count++ {
		if array[0] != array[count] {
			return count
		}
	}
	return 0
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
