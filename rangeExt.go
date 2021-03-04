package main

import (
    "bytes"
    "fmt"
    "strconv"
)

// Could not find a one-liner that does this :(.
func arrayToString(A []int, delim string) string {

    var buffer bytes.Buffer
    for i := 0; i < len(A); i++ {
        buffer.WriteString(strconv.Itoa(A[i]))
        if i != len(A)-1 {
            buffer.WriteString(delim)
        }
    }

    return buffer.String()
}

func main() {
    A := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
    fmt.Println(arrayToString(A, ", "))
}