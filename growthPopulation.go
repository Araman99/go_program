package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) int {
	NoOfYear := 0
	percent = percent / 100
	for i := p0; i < p; {
		i = p0 + int(float64(p0)*percent) + aug
		p0 = i
		NoOfYear++
	}
	return NoOfYear
}
func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))

}
