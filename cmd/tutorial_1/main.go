package main
import (
	"fmt"
	"errors"
	"unicode/utf8"
)

func main(){
	printValue := "Happy Birthday Adrian"
	printMe(printValue)

	numerator := 27
	denominator := 3
	var result, remainder, err = intDivision(numerator, denominator)
	if err!=nil{
		fmt.Println(err.Error())
	} else if remainder==0{
		fmt.Printf("The result of the division is %v", result)
	} else {
		fmt.Printf("The result of the division is %v and the remainder is %v", result, remainder)
	}

	
	fmt.Println(result, remainder)
	fmt.Println(result)
	fmt.Println(remainder, remainder)
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
	fmt.Println(myConst)

	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator==0 {
		err = errors.New("Cannot Divide by Zero")
	}
	
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}