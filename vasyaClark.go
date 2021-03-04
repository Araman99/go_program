package main

import "fmt"

func main() {
	fmt.Println(Tickets([]int{25, 25, 50}))
	fmt.Println(Tickets([]int{25, 100}))
	fmt.Println(Tickets([]int{25, 25, 50, 50, 100}))
}

func Tickets(peopleInLine []int) string {
	var a25, a50 = 0, 0
	for _, i := range peopleInLine {
		if i == 25 {
			a25 += 1
		}
		if i == 50 {
			a25 -= 1
			a50 += 1
		}
		if i == 100 {
			if a50 == 0 && a25 >= 3 {
				a25 -= 3
			} else {
				a25 -= 1
				a50 -= 1
			}
		}
		if a25 < 0 || a50 < 0 {
			return "NO"
		}
	}
	return "YES"
}
