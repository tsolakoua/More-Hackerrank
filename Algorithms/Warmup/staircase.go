package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	myMatrix := make([][]string, n)
	for k := 0; k < len(myMatrix); k++ {
		myMatrix[k] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i + j) > (n - 1) {
				myMatrix[i][j] = " "
			} else {
				myMatrix[i][j] = "#"
			}
		}
	}

	for i := len(myMatrix) - 1; i >= 0; i-- {
		for j := len(myMatrix) - 1; j >= 0; j-- {
			fmt.Print(myMatrix[i][j])
		}

		fmt.Println()
	}
}
