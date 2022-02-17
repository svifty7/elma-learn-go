package main

import (
	"fmt"
	"sort"
)

func main() {
	res, err := sequence([]int{10, 9, 4, 1, 2, 3, 7, 5, 6, 8})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func sequence(A []int) (int, error) {
	err := checkSequence(A)

	if err != nil {
		return 0, err
	}

	if len(A) == 1 {
		return 1, nil
	}

	sort.Ints(A)

	res := 1

	for i := 0; i < len(A)-1 && res != 0; i++ {
		if A[i]+1 != A[i+1] {
			res = 0
		}
	}

	return res, nil
}

func checkSequence(A []int) error {
	if len(A) == 0 {
		return fmt.Errorf("пришел пустой массив")
	}

	if len(A) > 100000 {
		return fmt.Errorf("массив выходит за рамки диапазона [1..100 000]")
	}

	var err error

	for i, val := range A {
		if val < 1 || val > 1000000 {
			err = fmt.Errorf("элемента массива с индексом %d выходит за рамки диапазона [1..1 000 000]", i)
		}
	}

	return err
}
