package main

import (
	"fmt"
	"math"
)
func main() {

	//Prints
	fmt.Println("This is a test.")
	fmt.Println("test" + "anothertest")
	fmt.Println("2 * 4 =",2*4)
	fmt.Println("7 / 4 =",7.0/4.0)
	fmt.Println("Answer is =",true && true)
	fmt.Println("Answer is =",false && true)
	fmt.Println("Answer is =",(true&&false)||(false||true))

	//Declaring Variable without Type
	test := "Batuhan"
	fmt.Println(test)

	//Multiple Variable Declarations
	var b, c int = 5,4
	fmt.Println(c+b)
	//It doesnt matter if their types are same or nor. We can still get the sum.

	//Default Values of Types
	var X int
	var Y string
	var Z float32
	fmt.Println("int=",X,"string",Y,"float",Z)

	//Const Variables and Math Pi Value
	const T = 3.14159265
	fmt.Println(T)
	fmt.Println(math.Pi)

	//For Loop (Zero to Ten)
	fmt.Print("Values: ")
	for i:=0; i<=10; i++ {
		fmt.Print(i)
	}

	fmt.Println("-")

	//** Break will terminate loop immediatly, Continue only terminates the current iteration of loop

	//For Loop (With Continue)
	fmt.Print("Values: ")
	for i:=0; i<=10; i++ {
		if i==5 || i==6 || i==7 {
			continue
		}
		fmt.Print(i)
	}

	fmt.Println("-")

	//For Loop (With Break)
	fmt.Print("Values: ")
	for i:=0; i<=10; i++ {
		if i==5 || i==6 || i==7 {
			break
		}
		fmt.Print(i)
	}

	fmt.Println("")


	var num  = 55
	//If Else
	if num < 10 {
		fmt.Println(num, "is less than 10")
	} else if num > 10{
		fmt.Println(num,"is greater than 10")
	} else {
		fmt.Println("Error")
	}

	var num2 = 3

	switch num2{
	case 1:{
		fmt.Println("Your number is ",num2)
	}
	case 2:{
		fmt.Println("Your number + 5 =",num2+5)
	}
	case 3:{
		fmt.Println("Your number * 5 =",num2*5)
	}
	default:{
		fmt.Println("Wrong choice")
	}
	}

}
