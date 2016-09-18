package main
import "fmt"

func main() {

	var n, k, q, m int
	fmt.Scanf("%d %d %d", &n, &k, &q)
	k%=n

	mySlice := make([]int, n)
	for i:=0; i< len(mySlice); i++ {
	        fmt.Scanf("%d", &mySlice[i])
	}

	for i:=0; i< q; i++ {
		fmt.Scan(&m)
		fmt.Println(mySlice[(n-k+m)%n]);
	}
}
