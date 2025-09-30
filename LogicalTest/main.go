package main

import "fmt"

func main () {
	fmt.Println("Hello, Logical Test!")
	fmt.Println(reverseText("olleh, dllrow"))
	fmt.Println("FizzBuzz up to 100:")
	fizbuzz(100)
	fmt.Println("Fibonacci for 0, 9:", Fibonacci(1, 9))
	fmt.Println("Get numbers from slice:", GetNumberOfSlice([]any{2, "h", 6, "u","y","t", 7, "j","y","h", 8}))
	fmt.Println("Max profit from stock prices:", maxProfit([]int{7,8,3,10,8}))
}


// test point 1
func reverseText(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
// test point 2
func fizbuzz(number int){
	for i := 1; i <= number; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
// test point 3
func Fibonacci(startNumber int, size int) []int {
	if size <= 0 {
		return []int{}
	}
	result := make([]int, size)
	result[0] = startNumber
	if size > 1 {
		result[1] = 1
	}
	for i := 2; i < size; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

// test point 4
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		profit := prices[i] - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// test point 5
func GetNumberOfSlice(slice []any) int {
	var numbers []int
	for _, v := range slice {
		switch value := v.(type) {
		case int:
			numbers = append(numbers, value)
		case float64:
			numbers = append(numbers, int(value))
		case float32:
			numbers = append(numbers, int(value))
		case string:
			var num int
			_, err := fmt.Sscanf(value, "%d", &num)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
	}
	return len(numbers)
}