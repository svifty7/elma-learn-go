package main

import (
	"fmt"
)

func main() {
	res, err := getSingleElement([]int{1, 1, 2, 1, 3, 3, 4, 7, 6, 7, 4, 6})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func getSingleElement(A []int) (int, error) {
	numbers := map[int]uint8{}

	for _, val := range A {
		if numbers[val] > 2 {
			continue
		}

		numbers[val]++
	}

	var res []int

	for key, val := range numbers {
		if val == 1 {
			res = append(res, key)
		}
	}

	if len(res) > 1 {
		return 0, fmt.Errorf("во входных данных есть несколько повторяющихся элементов")
	}

	if len(res) == 0 {
		return 0, fmt.Errorf("во входных данных нет одиночных элементов")
	}

	return res[0], nil
}
