package main

import "fmt"

func main() {
    var n, countPositive, countNegative, countZero int
    fmt.Scan(&n)

    slice := make([]int, n)
    
    for i := 0; i< n; i++ {
	fmt.Scanf("%d", &slice[i])
    }

    for i := 0; i< n; i++ {
       if slice[i] > 0 {
       	countPositive++
       } else if slice[i] < 0 {
       	countNegative++
       } else {
       	countZero++
       }
    }

    fmt.Println(float64(countPositive)/float64(n))
    fmt.Println(float64(countNegative)/float64(n))
    fmt.Println(float64(countZero)/float64(n))

}
