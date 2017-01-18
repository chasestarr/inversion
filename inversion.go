package inversion

import "fmt"

// Count returns the number of inversions found in array of ints
// Given a pair (i, j), the pair is an inversion if i < j
func Count(data []int) int {
	_, count := sortAndCount(data, 0)
	fmt.Println(count)
	return count
}

func sortAndCount(data []int, count int) ([]int, int) {
	// base case
	if len(data) <= 1 {
		return data, count
	}

	mid := len(data) / 2
	left, leftc := sortAndCount(data[:mid], count)
	right, rightc := sortAndCount(data[mid:], count)

	return mergeAndCount(left, leftc, right, rightc)
}

func mergeAndCount(a []int, ac int, b []int, bc int) ([]int, int) {
	result := []int{}
	count := ac + bc
	var i int
	var j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			// split inversion increment count
			count += len(a) - i
			result = append(result, b[j])
			j++
		}
	}

	if i < len(a) {
		result = append(result, a[i:]...)
	}
	if j < len(b) {
		result = append(result, b[j:]...)
	}
	return result, count
}
