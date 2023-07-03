package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(candidates, target, []int{}, 0, &result)
	return result
}

func backtrack(candidates []int, target int, current []int, start int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*result = append(*result, append([]int{}, current...))
		return
	}
	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], current, i, result)
		current = current[:len(current)-1]
	}
}

func main() {
	numbers := []int{2, 3, 6, 7}
	target := 7
	combinations := combinationSum(numbers, target)
	fmt.Printf("Все комбинации суммы для чисел %v и целевой суммы %d:\n", numbers, target)
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
