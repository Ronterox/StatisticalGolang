package main

import (
	"fmt"
)

func sum(arr []int) int {
    total := 0
    for _, n := range arr  {
        total += n
    }
    return total
}

func main() {
    var nums []int
    for i := range 10 {
        nums = append(nums, i + 1)
    }

    fmt.Printf("Hi mom ! %v\n", sum(nums))
}
