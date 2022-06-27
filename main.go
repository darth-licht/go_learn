package main

import (
	"fmt"
	"go_learn/basics"
	"go_learn/concurrency"
	"go_learn/shopping"
)

func main() {
	basics.ShowcaseBasics()
	basics.ShowcaseInterfaces()
	fmt.Println(shopping.PriceCheck(4343))
	concurrency.ShowcaseGoroutines()
	concurrency.ShowcaseChannels()

	//var input int
	//_, err := fmt.Scan(&input)
	//if err == io.EOF {
	//	fmt.Println("no more input!")
	//}
}
