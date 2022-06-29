package basics

import (
	"fmt"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

// Log function from Logger interface
// implementation is implicit
func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

// this method comes from the Stringer interface
// a Stringer can describe itself as a string
// it is akin to the toString method in java
// the print functions in fmt look for this method
func (l ConsoleLogger) String() string {
	return fmt.Sprint("I am a logger!!")
}

type MyError struct{}

// this method is from the error interface
// now MyError is implicitly and (error)
func (e MyError) Error() string {
	return fmt.Sprint("error!!")
}

type Add func(a int, b int) int // function type

func process(adder Add) int {
	return adder(1, 2)
}

//func add(a interface{}, b interface{}) interface{} {
//}

func ShowcaseInterfaces() {
	// ConsoleLogger is implicitly Logger cos it has a Log(string) function
	var logger Logger = ConsoleLogger{}
	logger.Log(">>>>Start interfaces<<<<")
	fmt.Printf("%d\n", process(func(a int, b int) int {
		return a + b
	}))
	fmt.Println(logger.(ConsoleLogger))
	logger.Log(">>>>End interfaces<<<<")
}
