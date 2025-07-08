package main

import "fmt"

func main() {
	first := []int{5, 1, 2, 5}
	second := []int{4, 2, 5, 1, 1, 2}

	fmt.Println(twoUnique(first, second))
	fmt.Println(intersection(first, second))
	fmt.Println(uniqueUnion(first, second))

}

func unique(slice []int) []int {
	result := make([]int, 0)

	for _, value := range slice {
		isUnique := true
		for _, resultValue := range result {
			if resultValue == value {
				isUnique = false
				break
			}
		}
		if isUnique {
			result = append(result, value)
		}
	}
	return result
}

func twoUnique(firstSlice []int, secondSlice []int) ([]int, []int) {
	return unique(firstSlice), unique(secondSlice)
}

func intersection(firstSlice []int, secondSlice []int) []int {
	result := make([]int, 0)
	firstSlice = unique(firstSlice)
	secondSlice = unique(secondSlice)

	for _, firstValue := range firstSlice {
		for _, secondValue := range secondSlice {
			if firstValue == secondValue {
				result = append(result, secondValue)
			}
		}
	}
	return result
}

func uniqueUnion(firstSlice []int, secondSlice []int) []int {
	result := append(firstSlice, secondSlice...)
	return unique(result)
}
