package main 
import "fmt"

func main() {

	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)
    if v1-v2!=0 && x2 >= x1 && v1 > v2  {
		j := (x2-x1)%(v1-v2)
		if j == 0 {
			fmt.Println("YES") 
		} else {
			fmt.Println("NO") 
		}

	} else {
		fmt.Println("NO") 
	}
	
}
