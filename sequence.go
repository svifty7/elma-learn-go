package main

import "fmt"

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

	length := len(A)
	isExist := false

	for i := 0; i < length; i++ {
		isExist = false

		if (i == length-1) && A[i]-1 == A[i-1] {
			return 1, nil
		}

		for j := 0; isExist != true && j < length; j++ {
			isExist = A[j] == A[i]+1 || A[j] == A[i]-1
		}

		if isExist == false {
			return 0, nil
		}
	}

	return 1, nil
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
