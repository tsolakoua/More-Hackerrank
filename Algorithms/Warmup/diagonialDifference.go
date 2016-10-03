package  main
import "fmt"
import "math"

func main() {
	
    var n, sumA, sumB int
    fmt.Scan(&n)

    myMatrix := make([][]int, n)
    for k := 0; k < len(myMatrix); k++ {
        myMatrix[k] = make([]int, n)
    }
    
    for i := 0 ; i < len(myMatrix) ; i ++  {
        for j := 0 ; j < len(myMatrix) ; j ++ {            
            fmt.Scanf("%d", &myMatrix[i][j])
        }
    }

    for i := 0 ; i < len(myMatrix) ; i ++  {
        for j := 0 ; j < len(myMatrix) ; j ++ {
                if i == j {
                     sumA = sumA + myMatrix[i][j]
                }
            if i+j == n-1 {
                sumB = sumB + myMatrix[i][j]
            }
            
        }
    }

    fmt.Println(math.Abs(float64(sumA-sumB)))
}
