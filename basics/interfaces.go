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
func (logger ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type Add func(a int, b int) int

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
	logger.Log(">>>>End interfaces<<<<")
}
