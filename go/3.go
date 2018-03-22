package main
import "fmt"
import "math"

func main() {
	fmt.Println("lpf: ", largeprimefactor(600851475143))
}

func largeprimefactor(number int) int {
	num := float64(number)

	i := 3
	highval := int(math.Sqrt(num))
	lpf := 0
	for i <= highval  {
		for number % i == 0 {
			lpf = i
			number = number / i
		}
		i = i + 2
	}

	return lpf
}