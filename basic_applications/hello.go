package main

import ("fmt")

func add(x ,y float64) float64{
	return x+y
}
func main() {
	var num float64=5.8
	var num1 float64=5.9

	fmt.Println(add(num,num1))
}