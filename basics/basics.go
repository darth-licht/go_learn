package basics

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

func CopySlice() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	fmt.Println(scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(worst)
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
