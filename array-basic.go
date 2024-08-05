package main

import "fmt"

func main() {
    // Declare and initialize an array
    arr := [5]int{10, 20, 30, 40, 50}

    // Access and modify an element
    fmt.Println("Original:", arr[2])  // Output: 30
    arr[2] = 35
    fmt.Println("Modified:", arr[2])  // Output: 35

    // Print the length of the array
    fmt.Println("Length:", len(arr))  // Output: 5

    // Iterate over the array using a for loop
    fmt.Println("Array elements:")
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }

    // Iterate over the array using range
    fmt.Println("Using range:")
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Multidimensional array example
    matrix := [2][2]int{
        {1, 2},
        {3, 4},
    }
    fmt.Println("2D Array Element [1][1]:", matrix[1][1])  // Output: 4
}
