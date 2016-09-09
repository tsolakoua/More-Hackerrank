package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, sum int
    fmt.Scan(&n)
    sliceA := make([]int, n)
    
    for i := 0; i< n; i++ {
        var input int
		fmt.Scan(&input)
		sliceA[i] = input
    }
    
    for i := 0; i< n; i++ {
        sum = sum + sliceA[i]
    }
    
    fmt.Println(sum)
    
    
}