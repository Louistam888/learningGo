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

// func main() {
// 	marks := []int{95, 82, 60, 40}

// 	for _, mark := range marks {
// 		switch {
// 		case mark >= 90:
// 			fmt.Println("Topper")
// 		case mark >= 75:
// 			fmt.Println("Good")
// 		default:
// 			fmt.Println("Needs Improvement")
// 		}
// 	}
// }

// func divide(a, b int) (int, int) {
// 	return a / b, a % b
// }

// func main() {
// 	fmt.Println(divide(2, 4))
// }

// func getUserInfo() (name string, age int) {
// 	name = "Alice"
// 	age = 30
// 	return
// }

// func main() {
// 	fmt.Println(getUserInfo())
// }

// func sum(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

// func main() {
// 	fmt.Println(sum(1, 2, 3, 4, 5))
// }

// func greet(name string) {
// 	fmt.Println("hello", name)
// }

// var sayHello = greet

// func main() {
// 	sayHello("suntax")
// // }

// func main() {
// 	func() {
// 		fmt.Println("anon function")
// 	}()
// }

// func main() {
// 	names := []string{"alice", "bob"}
// 	names = append(names, "delta")
// 	fmt.Println(names)
// }

// func main() {
// 	arr := []interface{}{1, 2, 3, 4, 5}
// 	slice := arr[1:4]
// 	fmt.Println(slice)
// 	fmt.Println(len(slice), slice)
// 	fmt.Println(cap(slice), "cap")
// }

// func main() {
// 	temps:=[]float64{22,3,4,5,6}
// 	temps = append(temps, 300)
// 	// fmt.Println(temps)
// 	for i, t:=range temps{
// 		fmt.Printf("day %d: %.1fC\n", i+1, t)
// 	}
// }

// func main() {
// 	ages := map[string]int{
// 		"Alice": 25,
// 		"Bob":   30,
// 	}
// 	fmt.Println(ages)
// 	ages["Charlie"] = 32
// 	delete(ages, "Bob")
// 	fmt.Println(ages)
// 	val, ok:=ages["Derek"]
// 	if(ok) {
// 		fmt.Println("Derek's age is", val)
// 	} else {
// 		fmt.Println("Derek not found")
// 	}
// }

// func main() {
// 	type Person struct{
// 		Name string
// 		Age int
// 	}

// 	p1:= Person{"Alice",25}
// 	p2:= Person{"Joe",35}
// 	fmt.Println(p1, p2)
// 	fmt.Println(p1.Name, "p1")

// 	type Team struct {
// 		Name string
// 		Members []Person
// 	}

// 	team:=Team{
// 		Name: "SyntaxFuel",
// 		Members: []Person{p1,p2},
// 	}
// 	fmt.Println(team.Members)
// }

func main() {
	emp:=struct{
		ID int
		Name string
	}{101, "Charlie"}

	fmt.Println(emp)
}