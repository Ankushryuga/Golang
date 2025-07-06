package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name    string
	Hobbies []string
}

func main() {
	// ----- 1. Shallow Copy -----
	fmt.Println("🔹 Shallow Copy:")
	p1 := Person{
		Name:    "Ankush",
		Hobbies: []string{"Coding", "Gaming"},
	}
	p2 := p1 // Shallow copy

	p2.Hobbies[0] = "Sleeping"
	fmt.Println("p1.Hobbies:", p1.Hobbies) // [Sleeping Gaming] ❗
	fmt.Println("p2.Hobbies:", p2.Hobbies)

	// ----- 2. Deep Copy of Struct with Slice -----
	fmt.Println("\n🔹 Deep Copy Struct with Slice:")
	p3 := deepCopyPerson(p1)
	p3.Hobbies[0] = "Reading"
	fmt.Println("p1.Hobbies:", p1.Hobbies) // [Sleeping Gaming]
	fmt.Println("p3.Hobbies:", p3.Hobbies) // [Reading Gaming] ✅

	// ----- 3. Deep Copy of Slice -----
	fmt.Println("\n🔹 Deep Copy of Slice:")
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	s2[0] = 99
	fmt.Println("s1:", s1) // [1 2 3]
	fmt.Println("s2:", s2) // [99 2 3]

	// ----- 4. Deep Copy of Map -----
	fmt.Println("\n🔹 Deep Copy of Map:")
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[k] = v
	}
	m2["a"] = 100
	fmt.Println("m1:", m1) // map[a:1 b:2]
	fmt.Println("m2:", m2) // map[a:100 b:2]

	// ----- 5. Deep Copy Using encoding/gob -----
	fmt.Println("\n🔹 Deep Copy using gob (generic):")
	original := Person{
		Name:    "GobUser",
		Hobbies: []string{"Drawing", "Writing"},
	}
	cloned := deepCopyGeneric(original)
	cloned.Hobbies[0] = "Dancing"
	fmt.Println("original.Hobbies:", original.Hobbies) // [Drawing Writing]
	fmt.Println("cloned.Hobbies:", cloned.Hobbies)     // [Dancing Writing]
}

// Deep copy of a struct with slice (manual)
func deepCopyPerson(p Person) Person {
	newHobbies := make([]string, len(p.Hobbies))
	copy(newHobbies, p.Hobbies)
	return Person{
		Name:    p.Name,
		Hobbies: newHobbies,
	}
}

// Generic deep copy using gob
func deepCopyGeneric[T any](src T) T {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	_ = enc.Encode(src)
	var dst T
	_ = dec.Decode(&dst)
	return dst
}
