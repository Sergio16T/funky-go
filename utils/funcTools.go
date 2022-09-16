package utils

// RemoveDuplicates Instantiates new array and executes a shallow comparison internally to determine if the new array contains the
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

func Contains[T comparable](list []T, element T) bool {
	for i := range list {
		s := list[i]
		if s == element {
			return true
		}
	}
	return false
}

// Reduce Executes a user-supplied "reducer" callback function on each element of the source array, in order,
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
// Param f - filter function is a predicate, to test each element of the array and return a bool
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
