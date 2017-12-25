package funcs

import (
	"sync"
)

//MapFn given a list transform the values in the list and return a new list with the transform values
func MapFn(
	A []interface{}, fn func(interface{}) interface{},
) []interface{} {
	var result []interface{}
	for _, v := range A {
		result = append(result, fn(v))
	}
	return result
}

// ParMap function: Parrallel version of the map function.
func ParMap(
	A []interface{}, fn func(interface{}) interface{}, threads int,
) []interface{} {
	var wg sync.WaitGroup

	var result []interface{}
	values := make(chan []interface{})

	partition := (len(A) / threads) + 1

	startIndex := 0

	// Map values
	for i := 0; i < threads; i++ {
		endIndex := min(startIndex+partition, len(A))
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			values <- MapFn(A[start:end], fn)
		}(startIndex, endIndex)
		startIndex = endIndex
	}

	go func() {
		wg.Wait()
		close(values)
	}()

	// Reduce values
	for item := range values {
		result = append(result, item...)
	}
	return result
}

// Filter function: Based on the list and a predicate function return a new list with the values that satisfies the predicate.
func Filter(
	array []interface{}, fn func(interface{}) bool,
) []interface{} {
	var results []interface{}
	for _, v := range array {
		if fn(v) {
			results = append(results, v)
		}
	}
	return results
}

// Reduce function. Given an accumulator and a function with two arguments reduce their functionality and return a single value based on the accumulated value
func Reduce(
	array []interface{},
	fn func(a interface{}, b interface{}) interface{},
	initialValue interface{},
) interface{} {
	var accumulator interface{}

	if initialValue != nil {
		accumulator = initialValue
	} else {
		accumulator = array[0]
	}

	if initialValue == nil {
		for i := 1; i < len(array); i++ {
			accumulator = fn(accumulator, array[i])
		}
	} else {
		for _, val := range array {
			accumulator = fn(accumulator, val)
		}
	}
	return accumulator
}

// Compose function: Given a series of function, apply on after the other.
func Compose(
	fns ...func(interface{},
	) interface{}) func(interface{}) interface{} {
	return func(val interface{}) interface{} {
		result := val
		for _, fn := range fns {
			result = fn(result)
		}
		return result
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
