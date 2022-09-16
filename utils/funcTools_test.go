package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReducer(t *testing.T) {
	// Adder
	testA := []int{1, 5, 4}
	assert.Equal(t, 10, reduce(testA, func(previousValue int, element int) int {
		return previousValue + element
	}, 0))

	// Remove Duplicates
	testB := []int{1, 1, 2, 3, 4, 5, 4}
	expectedB := []int{1, 2, 3, 4, 5}
	var initialValue []int

	assert.Equal(t, expectedB, reduce(testB, func(previousValue []int, element int) []int {
		if !contains(previousValue, element) {
			previousValue = append(previousValue, element)
		}
		return previousValue
	}, initialValue))
}

func TestRemoveDuplicates(t *testing.T) {
	test := []int{1, 1, 2, 3, 4, 5, 4}
	expected := []int{1, 2, 3, 4, 5}

	assert.Equal(t, expected, removeDuplicates(test))
}

type TestPerson struct {
	name string
	age  int
}

func TestFilter(t *testing.T) {

	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}}

	assert.Equal(t, expected, filter(sampleList, func(person TestPerson, index int) bool {
		return person.age <= 30
	}))
}
