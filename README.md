<div align="center">
    <h1>Funky Go</h1>
    <h4>
    A Functional Programming package for Go development
    </h4>
    <p>
        <a href="https://pkg.go.dev/github.com/Sergio16T/funky_go">
            <img alt="Go Reference" src="https://pkg.go.dev/badge/github.com/gabriel-vasile/mimetype.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Sergio16T/funky_go">
            <img alt="go report A+" src="https://goreportcard.com/badge/github.com/Sergio16T/funky_go"/>
        </a>
        <a href="LICENSE">
            <img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg">
        </a>
    </p>
    <img alt="Go Pilot" src="./go-pilot.svg" width="500px" style="margin-top: -60px">
</div>


## Summary
Funky Go includes Higher Order Functions in addition to other 
Declarative utility Functions. The utils package provides utilities
for working with slices & arrays in Go.


## Installation

```
go get github.com/Sergio16T/funky_go@v0.1.7-beta
```
## Table of Contents

- [Examples](#examples)


## Examples

Reduce the given source array to a new array with duplicates removed
```go
sourceArray := []int{1, 1, 2, 3, 4, 5, 4}
var initialValue []int

reduced := utils.Reduce(sourceArray, func(previousValue []int, element int) []int {
		if !utils.Contains(previousValue, element) {
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

filtered := utils.Filter(sourceArray, func(person TestPerson, index int) bool {
    return person.age <= 30
})

// filtered ~ []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}}

```

ForEach person in the list add 5 years to their age
```go
type TestPerson struct {
    name string
    age  int
}

sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}

utils.ForEach(sampleList, func(person TestPerson, index int) {
    sampleList[index].age = person.age + 5
})

// sampleList ~ []TestPerson{{name: "Mickey", age: 35}, {name: "Minnie", age: 32}, {name: "Goofy", age: 27}, {name: "Donald", age: 37}}

```

Map over each person in the list and return a new list containing the ages of all the characters + 5
```go
type TestPerson struct {
    name string
    age  int
}

sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}


mapped := utils.Map(sampleList, func(person TestPerson, index int) int {
    agePlusFive := person.age + 5
    return agePlusFive
})

// mapped ~  []int{35, 32, 27, 37}
```

Find the first value that passes the given predicate.

Returns a pointer and the index of value in the given array
(Nil pointer, index -1 if not found).
```go
sampleList := []int{1, 2, 3, 4, 5}

found, i := utils.Find(sampleList, func(num int, index int) bool {
    return num == 2
})

// found ~ *int((*int)(0xc00001aba8))
// i ~ 1
if i == -1 {
    log.Printf("No match was found")
} else {
    log.Printf("Match %+v found at index %+v\n", *found, i)
}
```

FindIndex of the first value in the array that passes the given predicate
```go
sampleList := []int{1, 2, 3, 4, 5}

index := utils.FindIndex(sampleList, func(num int, index int) bool {
    return num == 2
})

// index ~ 1
```

Find Index Of Element in Source Array
```go
sampleList := []int{1, 2, 3, 4, 11, 5, 1, 2, 3, 2, 1, 0, 9}

index := utils.IndexOf(sampleList, 11)
// index ~ 4
```