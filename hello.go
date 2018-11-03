
package main

import (
	"fmt"
	"errors"
	"math"
)
// creating classes

type person struct {
	name string
	age int
}
func main() {
	// naming variables
	var x int = 5
	y := 7
	sumu := x + y
	fmt.Println(sumu)
	// statements
	if x > 6 {
	     fmt.Println("More than 6")
	} else if x > 4 {
	fmt.Println("small")
	} else {
	fmt.Println("not it")
	}
	//arrays and append function 
	a := []int{5,4,3,2,1}
	a = append(a, 13)
	fmt.Println(a)
	// creating dictionaries map function
	vertices := make(map[string]int)
	
	vertices["triangle"] = 2
	vertices["square"] = 3
	delete(vertices,"square")
	fmt.Println(vertices)

	// loops
	for  i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// looping something inside a n array

	arr := []string{"a","b","c"}

	for index, value := range arr {
		fmt.Println("index:",index,"value:",value)
	}
	// looping in a dictionary
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:",key,"value:",value)
	}

	result := sum(3,4)
	fmt.Println(result)

	results, errors := sqrt(16)
	if errors != nil {
		fmt.Println(errors)
	} else {
		fmt.Println(results)
	}

	p := person{name:"Jake", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)

	s :=7
	inc(&s)
	fmt.Println(s)
}

// creating functions
func sum( k int, l int) int {
	return k +l
	}
// multiple return functions 
	// go does not have exception
func sqrt(m float64) (float64, error) {
	if m < 0 {
		return 0, errors.New("Undefined for negatives numbers")
	}
	return math.Sqrt(m), nil
}	

// pointers f

func inc(x *int) {
	*x++
}