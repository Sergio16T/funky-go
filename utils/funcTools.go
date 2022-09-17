package utils

// RemoveDuplicates instantiates a new array and executes a shallow comparison internally to determine if the new array contains the
// element and appends the item if not.
// Returns a new array with duplicates removed.
func RemoveDuplicates[T comparable](list []T) []T {
	var uniqueList []T
	for i := range list {
		s := list[i]
		if !Contains(uniqueList, s) {
			uniqueList = append(uniqueList, s)
		}
	}

	return uniqueList
}

// Contains iterates through the provided array and does a shallow comparison to determine if the array contains
// the element.
func Contains[T comparable](list []T, element T) bool {
	for i := range list {
		s := list[i]
		if s == element {
			return true
		}
	}
	return false
}

// Reduce executes a user-supplied "reducer" callback function on each element of the source array, in order,
// passing in the return value from the calculation on the preceding element.
// The final result of running the reducer across all elements of the array is a single value.
//
// @Param source - source array upon which the reducer will traverse.
//
// @Param f - reducer (accumulator/previousValue A, element T) where T is the current element
// in the iteration of source array and previous value is the value resulting from the previous call to callbackFn f.
// Previous value is equal to the initialValue on the first call.
//
// @Param initialValue - represents our accumulator. A value to which previousValue is initialized the first time
// the callback is called.
//
// Return value - The value that results from running the "reducer" callback function to completion over the entire array.
func Reduce[T, A any](source []T, f func(A, T) A, initialValue A) A {
	acc := initialValue
	for _, v := range source {
		acc = f(acc, v)
	}
	return acc
}

// Filter creates a shallow copy of a portion of a given array, filtered down to just the elements
// from the given array that pass the test implemented by the provided function.
//
// @Param source - source array upon which filter will traverse.
//
// @Param f - filter function is a predicate, to test each element of the array and return a bool
//
// Return value - Array of elements from the given array that passed the test provided by the test function
func Filter[T any, I int, B bool](source []T, f func(T, I) B) []T {
	var filtered []T
	for i, v := range source {
		pass := f(v, I(i))
		if pass {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Map creates a new array populated with the results of calling a provided function on every element in the
// source array.
//
// @Param source - source array upon which map will traverse.
//
// @Param f - callback fn that takes the element in the current iteration and the index and appends the result of
// the function to a new array
//
// Return value - array of results populated by executing the callBack fn on each element of the given array
func Map[T, U any, I int](source []T, f func(T, I) U) []U {
	var mapped []U
	for i, v := range source {
		returnVal := f(v, I(i))
		mapped = append(mapped, returnVal)
	}
	return mapped
}

// ForEach iterates over the given array and executes a given callBack on each element of the source array
//
// @Param source - source array upon which forEach will traverse.
//
// @Param f - callback fn
func ForEach[T any, I int](source []T, f func(T, I)) {
	for i, v := range source {
		f(v, I(i))
	}
}

// Find returns the memory address of the first element in the provided array that satisfies the provided testing function.
// As well as the index of the element.
// If none of the elements satisfy the provided testing function, a nil pointer is returned with -1 for the index
//
// @Param source - source array upon which Find will traverse.
//
// @Param f - callBack fn(element, index) is a predicate to execute on each value in the array
//
// Return value - 2 values are returned.
//
// The first value is the memory address of the first element that passes the provided testing fn (nil pointer if no match is found).
// Note: Dereference the pointer to retrieve the value.
//
// The Second Value is the index at which it was found (-1 if not found)
func Find[T any, B bool, I int](source []T, f func(T, I) B) (*T, I) {
	for i, v := range source {
		pass := f(v, I(i))
		if pass {
			return &v, I(i)
		}
	}
	var noMatch *T
	return noMatch, -1
}

// @Todo
//func IndexOf() {
//
//}
//
//func FindIndex() {
//
//}
//
//func Every() {
//
//}
//
//func Some() {
//
//}
