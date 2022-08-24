package main
import "fmt"

func main(){
	resultDouble := double(5)
	fmt.Println(resultDouble)

	var resultDivide = divide(float32(resultDouble))
	fmt.Println(resultDivide)

	anotherVariable := add(double(8))
	fmt.Println("Nested functions that is equal to =", anotherVariable)

	greet()
}

func double(x int) int{
	return x * 2
}

func divide(y float32) float32{
	return y / 2
}

func add(a int) int{
	return a + 2
}

func greet (){
	fmt.Println("Hello Chossen one!")
}