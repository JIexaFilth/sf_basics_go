package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c float32
	var d string
	fmt.Println("Input first number")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("Input operation")
	fmt.Fscan(os.Stdin, &d)
	fmt.Println("Input second number")
	fmt.Fscan(os.Stdin, &b)

switch d {
  case "*": c = a*b
             fmt.Println (a, d, b, "=", c)
  case "/": c = a/b
             if b !=0 {
				     fmt.Println(a,d, "=", c)
			 }else{
				     fmt.Println("You cannot divide by 0!")
			 }
  case "+": c = a+b
             fmt.Println(a, d, b, "=", c)
  case "-": c = a-b
             fmt.Println(a, d, b, "=", c)  
  default:
             fmt.Println("Error, wrong operation")
 }
}

