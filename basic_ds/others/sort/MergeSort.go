package main 

import (
    "fmt"
)


func merge(array []int, begin int, mid int, end int) {
    
    leftSize := mid - begin
    rightSize := end - mid

    newSize := leftSize + rightSize
    result := make([]int, 0, newSize)

    l, r := 0, 0

    for l < leftSize && r < rightSize {

        lValue := array[begin+l]
        rValue := array[mid+r]

        if lValue < rValue {
            result = append(result, lValue)
            l++
        } else {
            result = append(result, rValue)
            r++
        }
    }

    result = append(result, array[begin+l:mid]...)
    result = append(result, array[mid+r:end]...)

    for i := 0; i < newSize; i++ {
        array[begin+i] = result[i]
    }
    return
}


func MergeSort(array []int, begin int, end int) {
    
    if end-begin > 1 {
        mid := begin + (end-begin+1)/2

        MergeSort(array, begin, mid)

        MergeSort(array, mid, end)

        merge(array, begin, mid, end)
    }

}



func main() {
    list := []int{ 5, 56, 3, 5, 2, 2, 1, 44, 566, 99, 100, 7 }
    MergeSort(list, 0, len(list))
    fmt.Println(list)
}
