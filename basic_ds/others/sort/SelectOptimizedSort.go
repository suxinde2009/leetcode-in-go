package main

import (
    "fmt"
)

func SelectOptimizedSort(list []int) {
    n := len(list)

    for i := 0; i < n/2; i++ {
        minIdx := i
        maxIdx := i


        for j := i + 1; j < n - i; j++ {
            if list[j] > list[maxIdx] {
                maxIdx = j
                continue
            }

            if list[j] < list[minIdx] {
                minIdx = j
            }
        }
        
        if maxIdx == i && minIdx != n-i-1 {
            
            list[n-i-1], list[maxIdx] = list[maxIdx], list[n-i-1]
            list[i], list[minIdx] = list[minIdx], list[i]

        } else if maxIdx == i && minIdx == n-i-1 {
            list[minIdx], list[maxIdx] = list[maxIdx], list[minIdx]
    
        } else {
            list[i], list[minIdx] = list[minIdx], list[i]
            list[n-i-1], list[maxIdx] = list[maxIdx], list[n-i-1]
        }
    }
}


func main() {

    list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
    SelectOptimizedSort(list3)
    fmt.Println(list3)

    list4 := []int{5,9,2,6,8,14,6,49,25,4,6}
    SelectOptimizedSort(list4)
    fmt.Println(list4)

}

