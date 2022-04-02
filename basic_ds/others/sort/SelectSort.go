// Package main provides ...
package main 

import (
    "fmt"
)

func SelectSort(list []int) {
    n := len(list)

    for i := 0; i < n -1; i++ {
        min := list[i]

        minIdx := i

        for j := i+1; j < n; j++ {
            if list[j] < min {
                min = list[j]
                minIdx = j
            }
        }

        if i != minIdx {
            list[i], list[minIdx] = list[minIdx], list[i]
        }
    }
}


func main() {
    list := []int{5,9,1,6,8,14,6,49, 25,4,6,3}
    SelectSort(list)
    fmt.Println(list)
}
