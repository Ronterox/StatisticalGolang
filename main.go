package main

import (
	"fmt"
)

func matmul[T int8](a [][]T, b [][]T) ([][]T, error) {
	arow, acol := len(a), len(a[0])
	brow, bcol := len(b), len(b[0])

	if acol != brow {
		return nil, fmt.Errorf("Matrix multiplication not possible")
	}

	result := make([][]T, arow)
	for i := range arow {
		result[i] = make([]T, bcol)
	}

	result[0][0] = a[0][0]*b[0][0] + a[0][1]*b[1][0]
	result[0][1] = a[0][0]*b[0][1] + a[0][1]*b[1][1]
	result[1][0] = a[1][0]*b[0][0] + a[1][1]*b[1][0]
	result[1][1] = a[1][0]*b[0][1] + a[1][1]*b[1][1]

	return result, nil
}

func main() {
	t1 := [][]int8{{1, 2}, {1, 2}}
	t2 := [][]int8{{1, 2}, {1, 2}, {1, 2}}

    r1, err := matmul(t1, t1)
    if err != nil {
        fmt.Printf("\nerr: %v\n", err)
        return
    }

	fmt.Printf("%v times %v is %v", t1, t1, r1)

    r2, err := matmul(t1, t2)
    if err != nil {
        fmt.Printf("\nerr: %v\n", err)
        return
    }

	fmt.Printf("%v", r2)
}
