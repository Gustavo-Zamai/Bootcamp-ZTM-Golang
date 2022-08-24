package main
import "fmt"

func main(){
	myName := "Gustavo"
	var myLastName = "Zamai"
	fmt.Println("My name is", myName, myLastName)

	userName := "admin"
	fmt.Println("username =", userName)

	var sum int 

	part1, other :=  2,5
	fmt.Println("Part 1 is", part1, "and other is", other)

	part2, other := 5,0
	fmt.Println("Part 2 is", part2, "and other is", other)

	sum = part1 + part2
	fmt.Println("The sum of parts is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson name is", lessonName, "the type of lesson is", lessonType)

	word1, word2, _ := "Hello", "World", "!"
	fmt.Println(word1,word2)
}