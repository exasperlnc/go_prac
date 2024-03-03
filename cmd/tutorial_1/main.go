package main
import "fmt"
import "unicode/utf8"

func main(){
	printValue := "Happy Birthday Adrian"
	printMe(printValue)
}

func printMe(printValue string){
	var intNum int16 = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 5
	var intNum2 = 4
	fmt.Println(intNum1/intNum2)

	var myString string = "Hello World"
	fmt.Println(len(myString))

	fmt.Println(utf8.RuneCountInString("x"))

	myVar := "words"
	fmt.Println(myVar)

	const myConst float64 = 3.14159265

	fmt.Println(printValue)
}