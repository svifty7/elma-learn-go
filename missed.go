package main

import (
	"fmt"
	"sort"
)

func main() {
	res, err := GetMissedElement([]int{1, 2, 3, 4, 6, 7})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func GetMissedElement(A []int) (int, error) {
	err := checkMissed(A)

	if err != nil {
		return 0, err
	}

	if len(A) == 1 {
		return 1, nil
	}

	sort.Ints(A)

	for i := 0; i < len(A)-1; i++ {
		if A[i]+1 != A[i+1] {
			return A[i] + 1, nil
		}
	}

	return A[len(A)-1] + 1, nil
}

func checkMissed(A []int) error {
	if len(A) == 0 {
		return fmt.Errorf("пришел пустой массив")
	}

	if len(A) > 100000 {
		return fmt.Errorf("массив выходит за рамки диапазона [1..100 000]")
	}

	return nil
}
