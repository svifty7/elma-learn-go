package main

import (
	"errors"
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		solution, err := Solution([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, i)

		if err != nil {
			panic(err)
		}

		fmt.Println(solution)
	}
}

func Solution(A []int, K int) ([]int, error) {
	length := len(A)

	if length <= 1 {
		return A, nil
	}

	err := check(A, K)

	if err != nil {
		return A, err
	}

	A = append(make([]int, 0, length), A...)
	K = K % length

	A = append(A[length-K:], A[:length-K]...)

	return A, err
}

func check(A []int, K int) error {
	if K < 0 || K > 100 {
		err := errors.New("переменная K должна быть в диапазоне от 0 до 100")

		if err != nil {
			return err
		}
	}

	if len(A) > 100 {
		err := errors.New("длина массива A должна быть в диапазоне от 0 до 100")

		if err != nil {
			return err
		}
	}

	var err error = nil

	for i := 0; err == nil && i < len(A); i++ {
		if A[i] < -1000 || A[i] > 1000 {
			err = fmt.Errorf("элемент массива A с индексом %d", i)
		}
	}

	return err
}
