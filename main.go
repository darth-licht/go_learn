package main

import (
	"fmt"
	"github.com/darth-licht/go_learn/basics"
	"github.com/darth-licht/go_learn/shopping"
)

func main() {
	vegeta := new(basics.Saiyan) // create new *Saiyan
	vegeta.Person = new(basics.Person)
	vegeta.Name = "Vegeta"
	vegeta.Power = 9000
	fmt.Printf("%s - ", vegeta.Name)
	fmt.Println(vegeta)
	gohan := &basics.Saiyan{
		Person: &basics.Person{Name: "Gohan"},
		Power:  1000,
		Father: &basics.Saiyan{
			Person: &basics.Person{Name: "Goku"},
			Power:  9001,
			Father: nil,
		},
		Friends: make(map[string]*basics.Saiyan),
		//Friends: map[string]*Saiyan{},
	}
	//gohan.Introduce()
	//gohan.Father.Person.Introduce()
	saiyans := [2]*basics.Saiyan{gohan, gohan.Father} // array
	fmt.Printf("Array size = %d\n", len(saiyans))
	for index, saiyan := range saiyans {
		fmt.Printf("index = %d, name = %s\n", index, saiyan.Name)
	}

	superSaiyans := []*basics.Saiyan{gohan, gohan.Father} // slice
	fmt.Printf("Slice size = %d\n", len(superSaiyans))
	for index, saiyan := range superSaiyans {
		fmt.Printf("index = %d, name = %s\n", index, saiyan.Name)
	}
	fmt.Printf("Powers = %d\n", basics.ExtractPowers(superSaiyans))
	fmt.Printf("Powers2 = %d\n", basics.ExtractPowers2(superSaiyans))

	smartSaiyans := make([]*basics.Saiyan, 0, 10)
	fmt.Printf("Slice array capacity = %d\n", cap(smartSaiyans))
	//fmt.Println(smartSaiyans[1]) // panic, size is 0
	//smartSaiyans = smartSaiyans[0:8] // resize slice (only upto capacity, 10 in this case)
	//smartSaiyans = smartSaiyans[0:8] // takes slice from index 0 to 8-1
	//smartSaiyans = smartSaiyans[0:] // takes slice from index 0 to end
	//smartSaiyans = smartSaiyans[:8] // takes slice from index start to 8-1
	//fmt.Println(smartSaiyans[1]) // no panic, size is 8
	fmt.Printf("Slice size = %d\n", len(smartSaiyans))
	smartSaiyans = append(smartSaiyans, gohan) // auto resize backing array if needed
	fmt.Printf("Slice size = %d\n", len(smartSaiyans))
	var sai []basics.Saiyan // other slice init
	fmt.Printf("Slice size = %d\n", len(sai))

	scores := []int{1, 2, 3, 4, 5}
	scores = basics.RemoveAtIndex(scores, 2)
	fmt.Println(scores)

	basics.CopySlice()

	//lookup := make(map[string]int)
	// lookup := make(map[string]int, 100) // if you know the size prior
	lookup := map[string]int{ // composite literal
		"gohan": 2000,
	}
	lookup["goku"] = 9001
	power, exists := lookup["vegeta"]
	// prints 0 false. def value of int is 0
	fmt.Println(power, exists)
	fmt.Println("number of keys", len(lookup))
	delete(lookup, "goku")
	fmt.Println("number of keys", len(lookup))

	for key, value := range lookup {
		fmt.Println(key, value)
	}

	fmt.Println(shopping.PriceCheck(4343))
}