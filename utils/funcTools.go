package utils

// Utilities for Arrays & Slices

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
// @Param f - filter function is a predicate, to test each element of the array and return a bool.
//
// Return value - Array of elements from the given array that passed the test provided by the test function.
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
// the function to a new array.
//
// Return value - array of results populated by executing the callBack fn on each element of the given array.
func Map[T, U any, I int](source []T, f func(T, I) U) []U {
	var mapped []U
	for i, v := range source {
		returnVal := f(v, I(i))
		mapped = append(mapped, returnVal)
	}
	return mapped
}

// ForEach iterates over the given array and executes a given callBack on each element of the source array.
//
// @Param source - source array upon which forEach will traverse.
//
// @Param f - callback fn.
func ForEach[T any, I int](source []T, f func(T, I)) {
	for i, v := range source {
		f(v, I(i))
	}
}

// Every iterates through the provided array and determines whether all elements in the array
// pass the test implemented by the provided function
//
// @Param source - source array upon which Every will traverse.
//
// @Param f - callback fn is a predicate
//
// Return value - boolean
func Every[T any, I int, B bool](source []T, f func(T, I) B) B {
	for i, v := range source {
		passed := f(v, I(i))
		if !passed {
			return false
		}
	}
	return true
}

// Some iterates through the provided array and determines whether 1 of the elements in the array
// passes the test implemented by the provided function
//
// @Param source - source array upon which Some will traverse.
//
// @Param f - callback fn is a predicate
//
// Return value - boolean
func Some[T any, I int, B bool](source []T, f func(T, I) B) B {
	for i, v := range source {
		passed := f(v, I(i))
		if passed {
			return true
		}
	}
	return false
}

// Find returns a pointer and the index of first value in the given array that satisfies the provided predicate
// If none of the elements satisfy the provided testing function, a nil pointer is returned with -1 for the index
//
// @Param source - source array upon which Find will traverse.
//
// @Param f - callBack fn(element, index) is a predicate to execute on each value in the array.
//
// Return value - Returns a pointer and the index of first value in the given array that satisfies the provided predicate
// (Nil pointer, index -1 if not found).
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

// FindIndex returns the index of the first element in an array that satisfies the provided testing function.
// If no elements satisfy the testing function, -1 is returned.
//
// @Param source - source array upon which FindIndex will traverse.
//
// @Param f - callBack fn(element, index) is a predicate to execute on each value in the array.
//
// Return value - index of first value in the given array that satisfies the provided predicate or -1 if not found.
func FindIndex[T any, B bool, I int](source []T, f func(T, I) B) I {
	for i, v := range source {
		pass := f(v, I(i))
		if pass {
			return I(i)
		}
	}
	return -1
}

// IndexOf returns the first index at which a given element can be found in the array, or -1 if it is not present.
//
// @Param source - source array to traverse.
//
// @Param element - element to search for.
//
// Return value - first index at which the element can be found. -1 if not present.
func IndexOf[T comparable, I int](source []T, element T) I {
	for i, v := range source {
		if element == v {
			return I(i)
		}
	}
	return -1
}

// Contains iterates through the provided array and does a shallow comparison to determine if the array contains
// the element.
func Contains[T comparable](list []T, element T) bool {
	return IndexOf(list, element) >= 0
}

// RemoveDuplicates instantiates a new array and executes a shallow comparison internally to determine if the new array contains the
// element and appends the item if not.
// Returns a new array with duplicates removed.
//func RemoveDuplicates[T comparable](list []T) []T {
//	var uniqueList []T
//	for i := range list {
//		s := list[i]
//		if !Contains(uniqueList, s) {
//			uniqueList = append(uniqueList, s)
//		}
//	}
//
//	return uniqueList
//}
