# Goland Data Structures

## The basic data structures of Golang
#### 1. [Array](#arrays) 
#### 2. [Slice](#slices)
#### 3. [Map](#maps) 
#### 4. [Struct](#struct)


 ### *Arrays*

This data structure is used to store fixed number of elements. So once an array is defined, elements cannot be added or removed from the array. We can set the   values for any of the index in the array.
<br>
a. Defining the array to have package scope. Package scope it can be used different files that have a particular package name

``` var myArray [size]type  //general syntax ```

Examples:

```
var integerArray [5]int   // An integer array with five elements
var stringArray [4]string // A string array with four elements
```
b. Defining the array inside either the main function or any other function in your code. Such array is only accessible in the function that it is defined.

```
a := [5]int{10,20,30,40,50} //Defining and populating an array
b := [4]string{"first", "second", "third", "fourth"}
```
The ``` “:=” ``` is a short declaration operator

 ### *Slices*

Slices give a more robust interface to sequences compared to arrays. Unlike an array, no need to specify the length of the slice when defining it.
<br> <br>
We can create a slice in three way:
<br>
- Using var: ``` var s []int ```
- Using Shorthand Notation: ``` s := []int{1,2,3,4,5} ```
- Using make() : ``` s:=make([]int m, n) n ``` n is the number of elements

<br>

### We can [Add](#add-to-a-slice), [Remove](#delete-from-a-slice), [Replace](#replace-an-element-with-another) elements in a slice:

**a. Package scope slices:**

``` var slice_name []type //Syntax ```
<br>
Observe that we didn’t define “size”
<br>

#Examples:

```
package main

import "fmt"

//Package scope array definition
var integerSlice []int

var stringSlice []string

func main() {

	integerSlice = []int{10, 20, 30, 40}

	fmt.Println("This is the integer Slice: ", integerSlice)

	stringSlice = []string{"first", "second", "third"}

	fmt.Println("This is the string slice: ", stringSlice)

	printInteger()

}

//This function can access the integerSlice because integerSlice is package scoped
func printInteger() {
	fmt.Println("Print the integers: ", integerSlice)
}
```
<br>

**b. Function scope slice:**
Here slices are defined inside functions:

```

package main

import "fmt"

func main() {

	//Function scope slice:

	//EXAMPLE 1:
	integerSlice := []int{10, 20, 30, 40, 50}
	stringSlice := []string{"first", "second", "third", "fourth"}
	fmt.Println("This is the integer slice: ", integerSlice)
	fmt.Println("This is the string slice: ", stringSlice)

	// EXAMPLE 2:
	// We can also define slice using the "make" builtin function
	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	fmt.Println("Slice created with 'make': ", s)
}
```

**c. Slice Manipulations**

We would explore how we can perform a kind of “crud” operation with slice

<br>

From the above slice example:

```
s = []int{10,20,30,40}
```

#### **Add to a Slice**

We can add elements to this slice using “append” builtin function:

```
s = append(s, 50)
// After
s = []int{10,20,30,40,50}
```

We can append more than one element to the slice:

```
s = append(s, 60, 70)
// After
s = []int{10,20,30,40,50,60,70}
```

### **Delete from a slice:**

If “a” is a slice and we want to remove element at the “i” index, we do:

```
a = append(a[:i], a[i+1]...) //... is a variadic argument in Go
```

Say we want to remove “30” which is the third element at index “2”

```
s = append(s[:2], s[2+1:]...) 
// After
s = []int{10,20,40,50,60,70}
```
### **Replace an element with another**

If “a” is a slice and we want to replace element at the “i” index, with the last element:
```
a[i] = a[len(a)-1]
```

Say i want replace the third element now “40” (index of 2) with the last element “70”:
```
s[2] = s[len(s)-1]
//after
s =[]int{10,20,70,50,60,70}
```
**Get particular elements from the slice**

If “a” is a slice and we want to get particular elements from the slice.

```
a = []int{i, j, k, l, m, n}
```
To get elements j to m , we do:

```
a = a[j:n] //j is included, but n is not
```
We now have the slice: s =[]int{10,20,70,50,60,70}

To get the 2nd(index 1) to the 4th(index 3) element, we do:

```
s = s[1:4]
//after
 s =[]int{20,70,50}
```
**v. Length of Slice:**

The “crud” code:

```
package main

import "fmt"

func main() {

	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	fmt.Println("Slice created with 'make': ", s)

	//a. We can add elements to this slice using "append" builtin function:
	s = append(s, 50)
	fmt.Println("Added one element to slice: ", s)

	//b. We can append more than one element to the slice:
	s = append(s, 60, 70)
	fmt.Println("Added two elements to slice: ", s)

	//c. We can remove from that slice:
	//Say we want to remove "30" which is the third element at index "2"
	s = append(s[:2], s[2+1:]...) //... is a variadic argument in Go
	fmt.Println("Deleted one element from slice: ", s)

	//d. Replace an element with another
	// Say i want replace the third element now "40" with the fifth element "60":
	s[2] = s[len(s)-2]
	fmt.Println("Slice with element replaced: ", s)

	// Say i want replace the third element now "60" with the last element "70":
	s[2] = s[len(s)-1]
	fmt.Println("Updated Slice with element replaced: ", s)

	//e. Get particular elements from the slice:
	//We now have this slice: [10 20 70 50 60 70]
	//To get the 2nd(index 1) to the 4th(index 3) element, we do:
	s = s[1:4]
	fmt.Println("Slice with second to fourth element: ", s)

	//f. Get the length of the current slice:
	fmt.Println("Length: ", len(s))

	//g. Get the capacity of the current slice:
	fmt.Println("Capacity: ", cap(s)) //this is give "7"

	//h. Copy one slice to another:
	//Make a slice with the same length as "s":
	d := make([]int, len(s))
	copy(d, s)
	fmt.Println("This is the new slice: ", d)
}
```
Output:
```
Slice created with 'make':  [10 20 30 40]
Added one element to slice:  [10 20 30 40 50]
Added two elements to slice:  [10 20 30 40 50 60 70]
Deleted one element from slice:  [10 20 40 50 60 70]
Slice with element replaced:  [10 20 60 50 60 70]
Updated Slice with element replaced:  [10 20 70 50 60 70]
Slice with second to fourth element:  [20 70 50]
Length:  3
Capacity:  7
This is the new slice:  [20 70 50]
```

**d. Loop through a Slice**

We can do this in two ways:

1. Using range
2. Using for loop

```
package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40}

	// We can loop through this slice in two ways:

	// 1. using "range"
	for key, value := range s {
		fmt.Println(key, value)
	}
	//If we dont want the key, we do, replace "key" with "_":
	for _, value := range s {
		fmt.Println(value)
	}

	// 2. Using traditional forloop:
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //get the value at index "i"
	}
}
```
**e. Nested Slice:**

A slice can be nested:

```
nestedInt := make([][]int, 0) //slice of a slice of int
nestedString := [][]string{slice1, slice2}//slice of slice of string
```

Example

```
package main

import "fmt"

func main() {

  	//Example 1:
	nested := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			out = append(out, j)
		}
		nested = append(nested, out)
	}
	fmt.Println(nested)

	//Example 2:
	appleLaptops := []string{"MacbookPro", "MacbookAir"}
	hpLaptops := []string{"hp650", "hpEliteBook"}
	laptops := [][]string{appleLaptops, hpLaptops}
	for i, v := range laptops {
		fmt.Println("Record: ", i)
		for _, name := range v {
			fmt.Printf("\t Laptop name: %v \n", name)
		}
	}
}
```

Output:
```
[[0 1 2 3] [0 1 2 3] [0 1 2 3]]
Record:  0
         Laptop name: MacbookPro 
         Laptop name: MacbookAir 
Record:  1
         Laptop name: hp650 
         Laptop name: hpEliteBook
```

### *Maps*
It is possible for you to use key value pairs to refer to an element. This is accomplish using map. If you coming from a javascript, php background, see map as your object .


General syntax:
```
map[keyType]valueType
Map can be defined in three ways:
```

**i. Using var:**
```
var sampleMap = map[string]int
```

The above is the general syntax that applies both in package and function scope the key is of type string, while the value of the map is of type int.

**ii. Using Shorthand Notation:**
```
sampleMap := map[string]int
```

**iii. Using make( ):**

```
sampleMap := make(map[string]int)
```

**a. Defining a simple map**

```
package main

import "fmt"

var nameAgeMap map[string]int

func main() {

	nameAgeMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}
	fmt.Println("Print the age of James: ", nameAgeMap["James"])

	//We can range through the map and print each value:
	for key, value := range nameAgeMap {
		fmt.Printf("%v is %d years old\n", key, value)
	}
}
```
Output:

```
Print the age of James:  50
James is 50 years old
Ali is 39 years old
```

**b. Just like Slice, we can [Add], [Remove], [Replace] elements in a Map:**

```
package main

import "fmt"

func main() {

	currency := map[string]string{
		"AUD": "Australia Dollar",
		"GBP": "Great Britain Pound",
		"JPY": "Japan Yen",
		"CHF": "Switzerland Franc",
	}

	//a. Adding to the map:
	currency["USD"] = "USA Dollar"

	fmt.Println("Currency with USD added: ", currency)

	//b. Remove from the map:
	delete(currency, "GBP")
	fmt.Println("Currency with GBP deleted: ", currency)

	//c. Replacing one entry with another:
	currency["AUD"] = "New Zealand Dollar"
	fmt.Println("Currency with AUD value replaced with NZD: ", currency)

	//Ranging through the map:
	for key, value := range currency {
		fmt.Printf("%v might be equal to: %v\n", key, value)
	}
}
```

Output:

```
Currency with USD added:  map[AUD:Australia Dollar CHF:Switzerland Franc GBP:Great Britain Pound JPY:Japan Yen USD:USA Dollar]
Currency with GBP deleted:  map[AUD:Australia Dollar CHF:Switzerland Franc JPY:Japan Yen USD:USA Dollar]
Currency with AUD value replaced with NZD:  map[AUD:New Zealand Dollar CHF:Switzerland Franc JPY:Japan Yen USD:USA Dollar]
JPY might be equal to: Japan Yen
CHF might be equal to: Switzerland Franc
USD might be equal to: USA Dollar
AUD might be equal to: New Zealand Dollar
```

**c. Comma Ok Idiom:**

From the above example , we can check if a key exist, if they do, we the found value and “ ok” to it :
```
if value, ok := currency["USD"]; ok { //comma ok idiom
  fmt.Printf("The value %s is present\n", value) 
  fmt.Println(ok) //This will print "true"
} else {
  fmt.Println("We could not find that key")
}
```

**d. Map with a key of string/int and value of a slice of string/int:**

Example: A map with a key of string and value of a slice of string:

```
map[string][]string
```

```
package main

import "fmt"

func main() {

	//map with a key of string and value of a slice of string:

	nameAndHobby := map[string][]string{
		"Steven": []string{"Basket Ball", "Table Tennis", "Coding"},
		"Nnamdi": []string{"Sleeping", "Watching Movie", "Eating"},
	}
	//We can add someone else with their hobby:
	nameAndHobby["Timi"] = []string{"Watching Cartoon", 
					"Dreaming", 
					"Laughing", 
					"Lazing around",
				 }

	//We can delete from the map:
	delete(nameAndHobby, "Steven")

	for i, v := range nameAndHobby {
		fmt.Printf("%v likes \n", i)
		for j, value := range v {
			fmt.Printf("\t%v %v\n", j, value)
		}
	}
}
```

Output:
```
Nnamdi likes 
	0 Sleeping
	1 Watching Movie
	2 Eating
Timi likes 
	0 Watching Cartoon
	1 Dreaming
	2 Laughing
	3 Lazing around

```

**e. Nested Map: Map inside a Map**

```
package main

import "fmt"

func main() {

	currency := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	for key, value := range currency {
		fmt.Printf("Currency Name: %v\n", key)
		for k, v := range value {
			fmt.Printf("\t Currency Code: %v\n\t\t\t Ranking: %v\n\n", k, v)
		}
	}
}
```

Output:
```
Currency Name: Great Britain Pound
	 Currency Code: GBP
			 Ranking: 1

Currency Name: Euro
	 Currency Code: EUR
			 Ranking: 2

Currency Name: USA Dollar
	 Currency Code: USD
			 Ranking: 3
```


### **Struct**
Go struct is a collection of named fields/properties. A struct can have the same or different types of fields

A struct is defined using the “type” keyword:
```
type person struct {
    firstName string
    lastName  string 
    age       int 
}
```
Above, we defined a person struct that have three fields with their datatypes.


**a. A simple Struct**
```

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	//Assigning values to the fields in the person struct:
	p1 := person{
		firstName: "Mark",
		lastName:  "Kedu",
		age:       30,
	}

	fmt.Println("The is the person: ", p1)
}
```
Output:
```
The is the person:  {Mark Kedu 30}
```

**b.Struct with a slice field**
```
type animal struct {
    name string
    characteristics []string 
}
```

We use dot(.) to accees the individual elements in a struct:

```
package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

func main() {

	animal1 := animal{
		name: "Lion",
		characteristics: []string{"Eats human",
			"Wild Animal",
			"King of the jungle",
		},
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name: ", animal1.name)
	for _, v := range animal1.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}
```
Output:
```
Lion
   Eats human
   Wild Animal
   King of the jungle

```

**c. Nested Struct:**
A struct can be used inside another struct as seen below:

```
package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

//A herbivore is an animal, so it can have the animal struct as a field
type herbivore struct {
	animal
	eatHuman bool
}

func main() {

	herb := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{"Lacks sense",
				"Lazy",
				"Eat grass",
			},
		},
		eatHuman: false, //maybe
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name:", herb.animal.name)
	fmt.Println("Eats human? ", herb.eatHuman)
	fmt.Println("Characteristics:")
	for _, v := range herb.animal.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}
```

Output:
```
Animal name: Goat
Eats human?  false
Characteristics:
	 Lacks sense
	 Lazy
	 Eats grass`
```

**d. Promotion:**

From the example above, we can call the fields in the animal struct directly like this:

```
fmt.Println("Animal name:", herb.name)
fmt.Println("Eats human? ", herb.eatHuman)
fmt.Println("Characteristics:")
for _, v := range herb.characteristics {
fmt.Printf("\t %v\n", v)
}
```
Observe that we don’t need animal to access its fields again. It means that the animal fields are now seen as herbivores fields. This is called promotion.

**e. Anonymous Struct:**

A struct can be defined together with its field values. The example below have a struct with map and slice fields
```
package main

import "fmt"

func main() {
	bio := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Steven",
		friends: map[string]int{
			"Tim":   12345678,
			"Abdul": 34789657,
			"Kally": 28993332,
		},
		favDrinks: []string{
			"Pepsi",
			"Cocacola",
		},
	}
	fmt.Println(bio.firstName)

	for k, v := range bio.friends {
		fmt.Println(k, v)
	}

	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}

}
```

Output:
```
Steven
Abdul 34789657
Kally 28993332
Tim 12345678
0 Pepsi
1 Cocacola
```

**f. Functions with Receivers (Methods)**

A function can be defined that have a receiver of type struct:

If we define a struct:

```
type animal struct {
      name            string
      characteristics []string
}
```
```
package main

import "fmt"

type animal struct {
	name            string
	characteristics []string
}

func main() {

	animal1 := animal{
		name: "Elephant",
	}

	animal1.run()
}

func (a animal) run() {
	fmt.Println(a.name, "is a lazy animal hence cannot run")
}
```
Output:
```
Elephant is a lazy animal hence cannot run
```
From the above example, you can liken the run() function to be a field of the animal struct, because it is called just the same way a field is called (using the dot notation).

Let us cement the understanding we have about “receiver” by looking at interfaces.

**g. Interfaces**

Interfaces are named collections of method signatures. An interface is used to achieve polymorphism in Go.

What do i mean? An example perhaps will clarify what is said.

General Syntax:
```
type interfaceName interface {
   ...
}

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func calArea(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("This is the area of the square: ", s.area())
	case circle:
		fmt.Println("This is the area of the circle: ", s.area())
	}
}

func main() {
	c := circle{
		radius: 2.5,
	}
	s := square{
		length: 1.5,
	}

	calArea(c)
	calArea(s)

}
```
Output:
```
This is the area of the circle:  19.634954084936208
This is the area of the square:  2.25
```