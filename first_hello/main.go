package main

// NOTE: go mod init yourmodule(Fixes the red lines on the package main)
// git config --global init.defaultBranch <name>

import (
	"fmt"

	"project_1.com/greetings"

	"rsc.io/quote"
)



func main(){
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)


	fmt.Println(quote.Go())
	fmt.Println("Hello, World! This is my very first golang program")
	calculateSomething()
	loopy()
	listed()
}

func calculateSomething(){
	// Addition
	result_a := 5+2 
	fmt.Println("The answer is =", result_a)
	
	// Subtraction
	result_b := 5-2
	fmt.Println("The answer is =", result_b)
	
	// Multiplication
	var result_c int = 5*2
	fmt.Println("The answer is =", result_c)
	
	//Division
	var result_d int = 5/2
	fmt.Println("The answer is =", result_d)
	
	//Modulus
	var result_e int = 5%2
	fmt.Println("The answer is =", result_e)

    var age int
    age = 25

    name := "Alice"

    fmt.Printf("%s is %d years old.\n", name, age)

}


func loopy(){
    // Using a for loop to print numbers from 1 to 5
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
	}
    age := 18

    if age < 18 {
        fmt.Println("You're a minor.")
    } else if age == 18 {
        fmt.Println("You just turned 18!")
    } else {
        fmt.Println("You're an adult.")
    }

}

func listed(){
    // Array with a fixed size
    var fruitArray [3]string
    fruitArray[0] = "Apple"
    fruitArray[1] = "Banana"
    fruitArray[2] = "Cherry"

    // Slice (dynamic array)
    fruitSlice := []string{"Apple", "Banana", "Cherry"}

    fmt.Println(fruitArray)
    fmt.Println(fruitSlice)

}
























































































































































































// go install golang.org/x/website/tour@latest