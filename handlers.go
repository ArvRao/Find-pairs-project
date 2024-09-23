package main

import "reflect"

type RequestBody struct {
	Numbers []int `json: "numbers"`
	Target  uint  `json: "target"`
}

func validateArrayContainsIntegers(numbers []int) bool {
	for _, v := range numbers {
		if reflect.TypeOf(v).Kind() != reflect.Int {
			return false
		}
	}
	return true
}

func findAllPairs(numbers []int, target int) [][]int {
	allPairs := [][]int{}
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target && i != j {
				allPairs = append(allPairs, []int{i, j})
			}
		}
	}
	return allPairs
}
