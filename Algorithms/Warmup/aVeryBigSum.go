package main
import "fmt"

func main() {
	
	var n, sum int
	fmt.Scan(&n)
	my_slice := make([]int, n)

	for i:=0; i< len(my_slice); i++ {
		fmt.Scanf("%d", &my_slice[i])
	 }

	 for i:=0; i< len(my_slice); i++ {
	        sum = sum + my_slice[i]
	 }

	 fmt.Println(sum)
}
