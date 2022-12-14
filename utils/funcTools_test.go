package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestPerson struct {
	name string
	age  int
}

func TestReducer(t *testing.T) {
	// Adder
	testA := []int{1, 5, 4}
	assert.Equal(t, 10, Reduce(testA, func(previousValue int, element int) int {
		return previousValue + element
	}, 0))

	// Remove Duplicates
	testB := []int{1, 1, 2, 3, 4, 5, 4}
	expectedB := []int{1, 2, 3, 4, 5}
	var initialValue []int

	assert.Equal(t, expectedB, Reduce(testB, func(previousValue []int, element int) []int {
		if !Contains(previousValue, element) {
			previousValue = append(previousValue, element)
		}
		return previousValue
	}, initialValue))

	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expectedC := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}}
	var initialValueB []TestPerson

	assert.Equal(t, expectedC, Reduce(sampleList, func(previousValue []TestPerson, element TestPerson) []TestPerson {
		if element.name == "Mickey" || element.name == "Minnie" {
			previousValue = append(previousValue, element)
		}
		return previousValue
	}, initialValueB))

}

func TestFilter(t *testing.T) {

	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}}

	assert.Equal(t, expected, Filter(sampleList, func(person TestPerson, index int) bool {
		return person.age <= 30
	}))
}

func TestMap(t *testing.T) {
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := []TestPerson{{name: "Mickey", age: 35}, {name: "Minnie", age: 32}, {name: "Goofy", age: 27}, {name: "Donald", age: 37}}

	assert.Equal(t, expected, Map(sampleList, func(person TestPerson, index int) TestPerson {
		person.age += 5
		return person
	}))

	expectedB := []int{35, 32, 27, 37}
	assert.Equal(t, expectedB, Map(sampleList, func(person TestPerson, index int) int {
		plusFive := person.age + 5
		return plusFive
	}))
}

func TestForEach(t *testing.T) {
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := []TestPerson{{name: "Mickey", age: 35}, {name: "Minnie", age: 32}, {name: "Goofy", age: 27}, {name: "Donald", age: 37}}

	ForEach(sampleList, func(person TestPerson, index int) {
		sampleList[index].age = person.age + 5
	})
	assert.Equal(t, expected, sampleList)
}

func TestEvery(t *testing.T) {
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Matt", age: 22}}

	firstInitialIsM := Every(sampleList, func(person TestPerson, index int) bool {
		char := string(person.name[0])
		return char == "M"
	})

	firstInitialIsB := Every(sampleList, func(person TestPerson, index int) bool {
		char := string(person.name[0])
		return char == "B"
	})

	youngerThan30 := Every(sampleList, func(person TestPerson, index int) bool {
		return person.age < 30
	})

	thirtyAndYounger := Every(sampleList, func(person TestPerson, index int) bool {
		return person.age <= 30
	})

	assert.Equal(t, firstInitialIsM, true)
	assert.Equal(t, firstInitialIsB, false)
	assert.Equal(t, youngerThan30, false)
	assert.Equal(t, thirtyAndYounger, true)
}

func TestSome(t *testing.T) {
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Matt", age: 22}}

	atLeastOneIs30 := Some(sampleList, func(person TestPerson, index int) bool {
		return person.age == 22
	})

	atLeastOneIs31 := Some(sampleList, func(person TestPerson, index int) bool {
		return person.age == 31
	})

	assert.Equal(t, atLeastOneIs30, true)
	assert.Equal(t, atLeastOneIs31, false)
}

func TestFind(t *testing.T) {
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := &TestPerson{name: "Goofy", age: 22}

	foundGoofy, index := Find(sampleList, func(person TestPerson, index int) bool {
		return person.name == "Goofy" && person.age == 22
	})
	assert.Equal(t, index, 2)
	assert.Equal(t, expected, foundGoofy)

	characterFound, indexAt := Find(sampleList, func(person TestPerson, index int) bool {
		return person.name == "Shrek" && person.age == 22
	})
	assert.Equal(t, indexAt, -1)
	assert.Equal(t, true, characterFound == nil)

	sampleList2 := []int{1, 2, 3, 4, 5}
	expected2 := 2

	found, i := Find(sampleList2, func(num int, index int) bool {
		return num == 2
	})
	assert.Equal(t, i, 1)
	assert.Equal(t, expected2, *found)

}

func TestFindIndex(t *testing.T) {
	//FindIndex
	sampleList := []TestPerson{{name: "Mickey", age: 30}, {name: "Minnie", age: 27}, {name: "Goofy", age: 22}, {name: "Donald", age: 32}}
	expected := 2

	index := FindIndex(sampleList, func(person TestPerson, index int) bool {
		return person.name == "Goofy" && person.age == 22
	})

	assert.Equal(t, expected, index)

	indexNotFound := FindIndex(sampleList, func(person TestPerson, index int) bool {
		return person.name == "Goofy" && person.age == 28
	})

	assert.Equal(t, -1, indexNotFound)
}

func TestIndexOf(t *testing.T) {
	sampleList := []int{1, 2, 3, 4, 11, 5, 1, 2, 3, 2, 1, 0, 9}
	expected := 4

	index := IndexOf(sampleList, 11)
	assert.Equal(t, expected, index)

	notFound := IndexOf(sampleList, 12)
	assert.Equal(t, -1, notFound)
}

//func TestRemoveDuplicates(t *testing.T) {
//	test := []int{1, 1, 2, 3, 4, 5, 4}
//	expected := []int{1, 2, 3, 4, 5}
//
//	assert.Equal(t, expected, RemoveDuplicates(test))
//}
