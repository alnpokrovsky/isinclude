# isinclude

`XIX` Задача о поиске элемента

Дан упорядоченный массив чисел размером `N`
Нужно реализовать алгоритм поиска вхождения упорядоченного подмассива размера `M`, где `M << N`

```
func isInclude(array int[], subarray []int) bool

assert(isInclude([1, 2, 3, 5, 7, 9, 11], []) == true) 
assert(isInclude([1, 2, 3, 5, 7, 9, 11], [3, 5, 7]) == true) 
assert(isInclude([1, 2, 3, 5, 7, 9, 11], [4, 5, 7]) == false) 
``` 
