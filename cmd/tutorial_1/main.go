package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
	"math/rand"
	"sync"
)

type gasEngine struct{
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type owner struct{
	name string
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("\nYou can make it there!")
	}else{
		fmt.Println("\nNeed to fuel up first!")
	}
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

var wg = sync.WaitGroup{}
var results = []string{}

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
	sliceRace()
	stringPrac()
	var myEngine gasEngine = gasEngine{25,15, owner{"Logurt"}}
	fmt.Printf("\nThe engine's original mpg was %v", myEngine.mpg)
	myEngine.mpg = 30
	fmt.Printf("\nThe engine's updated mpg was %v", myEngine.mpg)
	fmt.Printf("\nThe owner of the engine is %v", myEngine.ownerInfo)
	canMakeIt(myEngine, 50)

	var p *int32 = new(int32)
	var i int32
	fmt.Printf("\nThe value p points to is: %v. The address p points to is %v", *p, p)
	fmt.Printf("\nThe value of i is %v", i)
	p = &i
	// this makes p point to the value of i. 
	// this is not quite the same as setting p = i, because changing i after setting p = *i will ALSO change the value of p. 

	var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		dbCall(i,dbData)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))

	// with waitgroup(concurrency)
	t1 := time.Now()
	m := sync.Mutex{}
	
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCallConcurrent(i,dbData, m)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(t1))
	fmt.Printf("\nThe results are %v\n", results)
	channelPrac()
	genericPrac()
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

func sliceRace(){
	var n int = 1000000
	var slowSlice = []int{}
	var fastSlice = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(slowSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(fastSlice, n))
}

func timeLoop(slice[]int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice,1)
	}
	return time.Since(t0)
}
func stringPrac(){
var strSlice = []string{"L","o","g","u","r","t"}
var strBuilder strings.Builder
for i := range strSlice{
	strBuilder.WriteString(strSlice[i])
}
var catStr = strBuilder.String()
fmt.Printf("\n%v", catStr)
}

func dbCall(i int, dbData []string){
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from teh database is :", dbData[i])
	results = append(results, dbData[i])
}

func dbCallConcurrent(i int, dbData []string, m sync.Mutex){
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from teh database is :", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	// the line above modifies the data, and so must be locked to prevent the concurrency from tripping over itself
	m.Unlock()
	wg.Done()
}

func channelPrac(){
	var c = make(chan int, 5)
	go process(c)
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	defer close(c)
	for i:=0; i<5; i++{
		c <- i
	}
}

func genericPrac(){
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))
	var emptySlice = []bool{}
	fmt.Println(isEmpty(emptySlice))
	fmt.Println(isEmpty(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}