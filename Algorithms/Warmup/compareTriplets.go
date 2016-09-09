package main
import "fmt"

func main() {

		n := 3
		alice := make([]int, n)
		bob := make([]int, n)
		var aliceTotal, bobTotal int

    	for i:=0; i< len(alice); i++ {
	        fmt.Scanf("%d", &alice[i])
	    }

	    for i:=0; i< len(bob); i++ {
	        fmt.Scanf("%d", &bob[i])
	    }

		for i:=0; i< len(alice); i++ {
			if alice[i] > bob [i] {
				aliceTotal = aliceTotal + 1

			} else if alice[i] < bob [i] {
				bobTotal = bobTotal + 1
			}
		}

		fmt.Print(aliceTotal, bobTotal)
}
