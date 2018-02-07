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

// Operation for parallel execution to ensure the order in which the operation is performed
type Operation struct {
	index int
	data  interface{}
}

// NewOperation instance
func NewOperation(index int, data interface{}) *Operation {
	return &Operation{index, data}
}

// ParMap function: Parrallel version of the map function.
func ParMap(
	A []interface{}, fn func(interface{}) interface{}, threads int,
) []interface{} {

	results := make([]interface{}, len(A))
	var wg sync.WaitGroup

	valueChan := make(chan *Operation)
	resultChan := make(chan *Operation)

	done := make(chan interface{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for index, item := range A {
			valueChan <- NewOperation(index, item)
		}
		close(valueChan)
	}()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for op := range valueChan {
				op.data = fn(op.data)
				resultChan <- op
			}
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	defer close(resultChan)

	for {
		select {
		case <-done:
			return results
		case op, ok := <-resultChan:
			if ok {
				results[op.index] = op.data
			}
		default:
		}
	}
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
	fns ...func(interface{}) interface{},
) func(interface{}) interface{} {
	return func(val interface{}) interface{} {
		result := val
		for i := len(fns) - 1; i >= 0; i-- {
			fn := fns[i]
			result = fn(result)
		}
		return result
	}
}

// Pipe performs the opposite of compose applys function from left to right.
func Pipe(
	fns ...func(interface{}) interface{},
) func(interface{}) interface{} {
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

func reverse(values []interface{}) {
	begin := 0
	end := len(values) - 1
	for begin < end {
		values[begin], values[end] = values[end], values[begin]
		begin++
		end--
	}
}
