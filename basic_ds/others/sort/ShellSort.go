// Package main provides ...
package main


import (
    "fmt"
)


func ShellSort(list []int) {
    
    n := len(list)

    for step := n / 2; step >= 1; step /= 2 {
        for i := step; i < n; i += step {
            
            for j := i - step; j >= 0; j -= step {
                if list[j+step] < list[j] {
                    list[j], list[j+step] = list[j+step], list[j]
                    continue
                }
                break
            }
        }
    }
}

func main() {
    list := []int{5,8,3,543,123,45,5,7,7,8,45,2,23,3,4,4,342,34,1,56,8,99,23}
    ShellSort(list)
    fmt.Println(list)
}
