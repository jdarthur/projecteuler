package main
import "fmt"

func main() {
	slice := fibo_list(4000000)
	fmt.Println("slice: ", slice)
	i := 0
	sum := 0
	for i < len(slice) {
		if slice[i] % 2 == 0 {
			sum  = sum + slice[i]
		}
		i = i + 1
	}
	fmt.Println("sum: ", sum)
}

func fibo_list(max_val int) []int {
	/* 
	Return list of fibonacci numbers less than max_val
	*/
	slice := []int{1, 1}
	i := 2
	for {
		last := slice[i - 1]
		next_to_last := slice[i - 2]
		slice = append(slice, last + next_to_last)
		if slice[i] > max_val {
			break
		}
		i = i + 1
	}
	return slice[1 : len(slice) - 1]
}