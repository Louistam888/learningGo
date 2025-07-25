package main

import "fmt"

var age int = 29
var name = "syntax fuel"

var isOnline bool = true
var price float64 = 99.99

const Pi = 3.1415

// func greet(name string) {
// 	fmt.Println("hello,", name)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	fmt.Println(add(1, 2))
// 	fmt.Println(divide(2, 4))
// }

// func main() {
// 	score := 85
// 	if score >= 90 {
// 		fmt.Println("Excellent")
// 	} else if score >= 75 {
// 		fmt.Println("Good")
// 	} else {
// 		fmt.Println("kKeep Practicing")
// 	}
// }

// func main() {
// 	day := "monday"
// 	switch day {
// 	case "monday":
// 		fmt.Println("start of the week")
// 	case "Friday":
// 		fmt.Println("Almost weekend")
// 	default:
// 		fmt.Println("just another day")
// 	}
// }

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// func main() {
// 	i := 1
// 	for i <= 5 {
// 		fmt.Println(i)
// 		i++
// 	}

// }

// func main() {
// 	nums := []int{10, 20, 30}
// 	for index, value := range nums {
// 		fmt.Printf("index: %d, Value: %d\n", index, value)
// 	}
// }

func main() {
	marks := []int{95, 82, 60, 40}

	for _, mark := range marks {
		switch {
		case mark >= 90:
			fmt.Println("Topper")
		case mark >= 75:
			fmt.Println("Good")
		default:
			fmt.Println("Needs Improvement")
		}
	}
}
