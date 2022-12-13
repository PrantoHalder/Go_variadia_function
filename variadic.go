package main

import "fmt"

func variadic (){
   fmt.Println("print the adddition of numbers " ,sum(1,2,3,4,5,6))
}

func sum (num ...int)int {
	res := 0
	for _,value := range num{
		res = res + value
	}
	return res
}