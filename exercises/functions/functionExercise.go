package main
import "fmt"

func main(){
	hello("Gustavo")
	fmt.Println(hi())
	any := anyNumber()
	fmt.Println(any)

	//anyONe, anyTwo := anyTwoNUmber()
	fmt.Println(anyTwoNumber())

	fmt.Println(addThree(5,6,3))

	g,h := anyTwoNumber()	
	result := addThree(anyNumber(), g, h)
	fmt.Println(result)



}

func hello(name string) {
	fmt.Println("Hello,",name,". Nice to meet you!")
}

func hi()string{
	return "Hi there!"
}

func anyNumber()int{
	return 8
}

func anyTwoNumber()(int, int){
	return  2, 5
}

func addThree(a,b,c int)int{
	return a+b+c
}