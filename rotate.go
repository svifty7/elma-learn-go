package main

import (
	"fmt"
)

func main() {
	res, err := getRotatedArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func getRotatedArray(A []int, K int) ([]int, error) {
	err := checkRotate(A, K)

	if err != nil {
		return A, err
	}

	if K == 0 {
		return A, err
	}

	length := len(A)

	if length <= 1 {
		return A, err
	}

	K = K % length
	isKNegative := K < 0

	if isKNegative {
		A = reverse(A)
		K = -K
	}

	A = append(A[length-K:], A[:length-K]...)

	if isKNegative {
		return reverse(A), err
	}

	return A, err
}

func checkRotate(A []int, K int) error {
	if K < -100 || K > 100 {
		err := fmt.Errorf("переменная K должна быть в диапазоне от -100 до 100")

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

func reverse(A []int) []int {
	res := make([]int, len(A))

	for i, val := range A {
		res[len(A)-i-1] = val
	}

	return res
}
