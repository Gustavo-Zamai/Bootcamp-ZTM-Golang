package main

import "fmt"

const Author = "Gus Simons"

func main() {
	/*
	//var number = 3
	//var number int = 3
	a, b, c := 1, 2, "float"

	var ( //block of vars
		x int     = 55
		y float32 = 22.5
		z string  = "Praise"
	)

	/*print(a, x)
	print(b, y)
	print(c, z)*/
	fmt.Print(Author)

	result := sum(5, 3)
	fmt.Print("\n" ,result)

	f, _, g := multiReturn()
	fmt.Print("\n",f,g,)
}


func sum(num1, num2 int) int{// function to sum 2 integer numbers
	return num1 + num2
}

func multiReturn()(int,int,int){
	return 1,2,5
}
