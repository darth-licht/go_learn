package main

import (
	"fmt"
	"github.com/darth-licht/go_learn/basics"
	"github.com/darth-licht/go_learn/concurrency"
	"github.com/darth-licht/go_learn/shopping"
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
