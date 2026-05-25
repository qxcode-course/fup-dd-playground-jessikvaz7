package main

import "fmt"

func menorindice (arr ) {
    
    arr := make([]int, 5)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
}
     fmt.Println(arr)
    imin := 0
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr [imin] {
            imin = i
        }
    }