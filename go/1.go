package main

import "fmt"

func main() {

	i := 0
	sum := 0
	for i < 1000 {
		if i % 3 == 0 {
			//fmt.Println("divisible by 3: ", i)
			sum = sum + i
		} else if i % 5 == 0 {
			//fmt.Println("divisible by 5: ", i)
			sum = sum + i
		}  
		i = i + 1 
	}
	fmt.Println("sum: ", sum)
}