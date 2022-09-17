<div align="center">
    <h1>Funky Go</h1>
    <h4>
    A Functional Programming package for Go development
    </h4>
    <img src="./go-pilot.svg" width="500px">
</div>


## Summary
Funky Go includes Higher Order Functions in addition to other 
Declarative utility Functions.


## Installation

```
go get github.com/Sergio16T/funky_go v0.1.5-beta
```
## Table of Contents

- [Examples](#examples)


## Examples

Using reduce to remove duplicates
```go
sourceArray := []int{1, 1, 2, 3, 4, 5, 4}
var initialValue []int

reduced := reduce(sourceArray, func(previousValue []int, element int) []int {
		if !contains(previousValue, element) {
			previousValue = append(previousValue, element)
		}
		return previousValue
	}, initialValue)

// reduced ~ [1, 2, 3, 4, 5]
```

Filter list of Disney characters to characters 30 or younger
```go
type TestPerson struct {
    name string
    age  int
}

sourceArray := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}

filtered := filter(sourceArray, func(person TestPerson, index int) bool {
    return person.age <= 30
})

// filtered ~ []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}}

```