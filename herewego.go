package main

import "fmt"

func main() {

  fmt.Println("Hello World")

// Types
  // var types
  // uint8 : unsigned 8-bit int (0-255)
  // uint16 : unsigned 16-bit int (0-65535)
  // uint32 : unsigned 32-bit int (0-4294967295)
  // uint64 : unsigned 64-bit int (0-18446744073709)
  // int8 : signed 8-bit integers (-128 - 127)
  // int16 : signed 16-bit integers (-32768 - 32767)
  // int32 : signed 32-bit integers (-2147483648 - 2147483647)
  // int64 : signed 64-bit integers (-9223372036854775808 - 9223372036854775807)

  var age int = 40

  var favNum float64 = 1.6180339

  //var isOver40 bool = false

  // Go will automatically put a space so the code below is unnecessary
  // fmt.Println(age, " ", favNum)
  fmt.Println(age, favNum)

  // Floats dont always give exact results
  var numOne = 1.000
  var num99 = .9999

  fmt.Println(numOne - num99)

// Arithmetic
  fmt.Println("6 + 4 =", 6 + 4)
  fmt.Println("6 - 4 =", 6 - 4)
  fmt.Println("6 * 4 =", 6 * 4)
  fmt.Println("6 / 4 =", 6 / 4)
  fmt.Println("6 % 4 =", 6 % 4)

// COnstants
  const pi float64 = 3.14159265

  // Decalring variables
  //var(
  //  varA = 2
  //  varB = 3
  //  )

// Strings
  var myName string = "Sion Williams"

  fmt.Println(len(myName))
  fmt.Println(myName + " is a robot")
  fmt.Println("I like \n\n")
  fmt.Println("newlines!!!")

  // Types of prints
  fmt.Printf("%f \n", pi)
  fmt.Printf("%.3f \n", pi) // .3 precision
  fmt.Printf("%T \n", pi) // print data type
  fmt.Printf("%d \n", 100) // print decimal
  fmt.Printf("%b \n", 100) // print binary
  fmt.Printf("%c \n", 100) // print characters
  fmt.Printf("%x \n", 17) // print hex
  fmt.Printf("%e \n", pi) // print scientific notation

// Conditions
  fmt.Println("true && false = ", true && false)
  fmt.Println("true || false = ", true || false)
  fmt.Println("!false = ", !false)

// relational operators
  // == equals
  // != not equals
  // < less than
  // > greater than
  // <= less than or equal to
  // >= more than or equal to

// loops
  i := 1

  for i <= 10 {
    fmt.Println(i)
    i++
  }

  for j := 0; j < 5; j++ {
    fmt.Println(j)
  }

}
