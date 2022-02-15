package main

import (
	"fmt"
)

func main() {
	solution, err := Solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(solution)
}

func Solution(A []int, K int) ([]int, error) {
	err := check(A, K)

	if err != nil {
		return A, err
	}

	length := len(A)

	if length <= 1 {
		return A, err
	}

	A = append(make([]int, 0, length), A...)
	K = K % length

	A = append(A[length-K:], A[:length-K]...)

	return A, err
}

func check(A []int, K int) error {
	if K < 0 || K > 100 {
		err := fmt.Errorf("переменная K должна быть в диапазоне от 0 до 100")

		if err != nil {
			return err
		}
	}

	if len(A) > 100 {
		err := fmt.Errorf("длина массива A должна быть в диапазоне от 0 до 100")

		if err != nil {
			return err
		}
	}

	var err error

	for i := 0; err == nil && i < len(A); i++ {
		if A[i] < -1000 || A[i] > 1000 {
			err = fmt.Errorf("элемент массива с индексом \"%d\" должен быть в диапазоне от -1000 до 1000", i)
		}
	}

	return err
}
