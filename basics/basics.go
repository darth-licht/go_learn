package basics

// structs, pointers, array, slices, maps
import (
	"fmt"
	"math/rand"
	"sort"
)

// RemoveAtIndex won't preserve order
func RemoveAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	// swap last value and value to remove
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func ExtractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}

func ExtractPowers2(saiyans []*Saiyan) []int {
	powers := make([]int, 0, len(saiyans))
	for _, saiyan := range saiyans {
		powers = append(powers, saiyan.Power)
	}
	return powers
}

type Saiyan struct {
	*Person
	Power   int
	Father  *Saiyan
	Friends map[string]*Saiyan
}

func (s *Saiyan) Super() {
	s.Power += 1000
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hello, I'm %s\n", s.Name)
}

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func ShowcaseBasics() {
	fmt.Println(">>>>Begin basics<<<<")
	vegeta := new(Saiyan) // create new *Saiyan
	vegeta.Person = new(Person)
	vegeta.Name = "Vegeta"
	vegeta.Power = 9000
	fmt.Printf("%s - ", vegeta.Name)
	fmt.Println(vegeta)
	gohan := &Saiyan{
		Person: &Person{Name: "Gohan"},
		Power:  1000,
		Father: &Saiyan{
			Person: &Person{Name: "Goku"},
			Power:  9001,
			Father: nil,
		},
		Friends: make(map[string]*Saiyan),
		//Friends: map[string]*Saiyan{},
	}
	//gohan.Introduce()
	//gohan.Father.Person.Introduce()
	saiyans := [2]*Saiyan{gohan, gohan.Father} // array
	fmt.Printf("Array size = %d\n", len(saiyans))
	for index, saiyan := range saiyans {
		fmt.Printf("index = %d, name = %s\n", index, saiyan.Name)
	}

	superSaiyans := []*Saiyan{gohan, gohan.Father} // slice
	fmt.Printf("Slice size = %d\n", len(superSaiyans))
	for index, saiyan := range superSaiyans {
		fmt.Printf("index = %d, name = %s\n", index, saiyan.Name)
	}
	fmt.Printf("Powers = %d\n", ExtractPowers(superSaiyans))
	fmt.Printf("Powers2 = %d\n", ExtractPowers2(superSaiyans))

	smartSaiyans := make([]*Saiyan, 0, 10)
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
	var sai []Saiyan // other slice init
	fmt.Printf("Slice size = %d\n", len(sai))

	scores := []int{1, 2, 3, 4, 5}
	scores = RemoveAtIndex(scores, 2)
	fmt.Println(scores)

	// copy slice
	scoress := make([]int, 100)
	for i := 0; i < 100; i++ {
		scoress[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scoress)
	fmt.Println(scoress)
	worst := make([]int, 5)
	copy(worst, scoress[:5])
	fmt.Println(worst)
	// end copy slice

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
	fmt.Println(">>>>Finish basics<<<<")
}
