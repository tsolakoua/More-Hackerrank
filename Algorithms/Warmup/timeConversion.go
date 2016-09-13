package main 
import "fmt"

func main() {

	var hour, minutes, seconds int
	var pm_am string

	fmt.Scanf("%d:%d:%d%s", &hour, &minutes, &seconds, &pm_am)

	if pm_am == "PM" && hour != 12 {
		hour = hour+ 12
	} else if pm_am == "AM" {
		if hour == 12 {
			hour = 0
		}
	}

	fmt.Printf("%.2d:%.2d:%.2d", hour, minutes, seconds)	
}
