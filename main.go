package main

import (
	"fmt"
)

// , all variables are initialized  to  their default   values
func main() {
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// multiple variable declaration
var (
	firstName string = "John"
	lastName  string = "Doe"
	age       int    = 30
)

// constants
const Pi = 3.14
// The value of a constant must be assigned when you declare it.

//mu;ltiple constant declaration
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusPending  = "pending"
)

// iota is a special identifier that can be used to create enumerated constants
const (
	Red = iota
	Green
	Blue
)
//\n creates new lines
// output:
Print() // prints its arguments with their default format.
println() // prints a new line
printf() // formatted output, similar to fmt.Printf but without the need to import fmt package
// %v is used to print the value of the arguments
// %T is used to print the type of the arguments

var i string = "Hello"
var j int = 15

fmt.Printf("i has value: %v and type: %T\n", i, i)
fmt.Printf("j has value: %v and type: %T", j, j)

//array declaration - Arrays are used to store multiple values of the same type in a single variable,
var arr [5]int // array of 5 integers, initialized to zero
arr[0] = 10 // assigning value to the first element of the array

var arr1 = [3]int{1,2,3} // array of 3 integers, initialized with values 1, 2, and 3
// array literal - a concise way to declare and initialize an array in a single lin
arr2 := [5]int{4,5,6,7,8}

// decale an array with inferred length
arr3 := [...]int{9,10,11} // array of 3 integers, length is inferred from the number of elements	

// accessing array elements
fmt.Println(arr1[0]) // prints the first element of arr1, which is 1
fmt.Println(arr2[2]) // prints the third element of arr2, which is 6

//chnge the value of an array element
arr1[1] = 20 // changes the second element of arr1 to 20
fmt.Println(arr1[1]) // prints the updated second element of arr1, which is 20	
//initialize an array with default values
var arr4 [5]int // array of 5 integers, initialized to zero
fmt.Println(arr4) // prints [0 0 0 0 0]
arr2 := [5]int{1,2} //partially initialized , prints [1 2 0 0 0]
arr3 := [5]int{1,2,3,4,5} //fully initialized , prints [1 2 3 4 5]

// Initialize Only Specific Elements
arr5 := [5]int{0: 10, 2: 30} // initializes the first and third elements, prints [10 0 30 0 0]


// multi-dimensional arrays
var matrix [2][3]int // 2 rows and 3 columns, initialized to zero
matrix[0][0] = 1 // assigning value to the first element of the first row
matrix[0][1] = 2 // assigning value to the second element of the first row
matrix[0][2] = 3	// assigning value to the third element of the first row

//i