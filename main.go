package main

import (
	"fmt"
	"statlearn/matrix"
)

func main() {
	m1 := matrix.New([][]float32{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}})

	r1, err := m1.Matmul(m1)
	if err != nil {
		fmt.Printf("\nerr: %v\n", err)
		return
	}

	fmt.Printf("%v x %v = %v", m1, m1, r1)
}
