package main

import (
	"fmt"
	
)

func main() {
//prompt user to enter two numbers +
// operation they want to perform 	
var num1, num2 float64
var operator string 
fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
fmt.Println("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

// print result 

var result float64
switch operator {
case "+":
	result = num1 + num2
case "-":
	result = num1 - num2
case "*":
	result = num1 * num2 
case "/":
	result = num1 / num2 
default:
fmt.Println("Invalid operator")	
return
}
fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)


}
