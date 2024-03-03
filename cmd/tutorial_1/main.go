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
	arrayPrac()
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

func arrayPrac(){
	var intArr [3]int32
	// default values are 0, so current array is int32 [0,0,0]
	// int32 is 4 bytes, so this array is 12 bytes 
	// 0 indexed
	intArr[1] = 5
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	intSlice := []int32{1,2,3}
	fmt.Printf("The original length of the slice was %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nAfter appending, the length was %v with capacity %v", len(intSlice), cap(intSlice))

	for i:=0; i<=10; i++{
		fmt.Println(i)
	}
}

func mapPrac(){
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var map2 = map[string]uint8{"Jim":23, "Carlos":47}
	fmt.Println(map2["Adam"])
	// maps always return something even if the key doesn't exist

	var jasonAge, ok = map2["Jason"]
	if ok{
		fmt.Printf("Jason's age is %v", jasonAge)
	}else{
		fmt.Printf("Invalid")
	}

	for name, age := range map2{
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}