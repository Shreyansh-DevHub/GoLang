package main

import (
	"fmt"
)

func calling(name, place string) string {
	return fmt.Sprintf("my name is %s and i live in %s", name, place)
}
func calc(a, b int, add string) int {
	switch add {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	case "^":
		return a ^ b
	default:
		fmt.Println("invalid operator")
	}
	return 0
}

func main() {

	var name string
	fmt.Println("enter your name")
	fmt.Scan(&name)
	fmt.Println("hello, ", name)
	fmt.Println(calc(3, 4, "^"))
	fmt.Println(calling("shreyansh", "gurugram"))
	/**for i := 0; i <= 10; i++ {
	fmt.Println("hello world")}**/
}
