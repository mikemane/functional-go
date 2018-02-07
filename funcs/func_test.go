package funcs

import (
	"log"
	"testing"
)

// TestMapFn map function
func TestMapFn(t *testing.T) {
	numbers := createArray(10)
	result := MapFn(numbers, double)

	actual := func() []interface{} {
		var r []interface{}
		for _, v := range numbers {
			r = append(r, double(v))
		}
		return r
	}()

	assertArray(t, result, actual)

}

func double(item interface{}) interface{} {
	return item.(int) * 2
}

func square(item interface{}) interface{} {
	i, ok := item.(int)
	if !ok {
		log.Fatal("Could not convert value to int")
	}
	return i * i
}

func assertArray(t *testing.T, expected, received []interface{}) {
	expectedLength := len(expected)
	receivedLength := len(received)
	if expectedLength != receivedLength {
		t.Fatalf(
			"Length of values are not equal\nExp (len): %d\nReceived (Len): %d",
			expectedLength, receivedLength,
		)
	}
	for i := 0; i < expectedLength; i++ {
		if expected[i] != received[i] {
			msg := "Values at index %d are not equal\nExp: %v\nReceived::%v"
			t.Fatalf(msg, i, expected[i], received[i])
		}
	}
}

// TestFilter function
func TestFilter(t *testing.T) {
	array := createArray(10)
	filterEven := func(v interface{}) bool {
		return v.(int)%2 == 0
	}
	result := Filter(array, filterEven)

	actual := func(arr []interface{}) []interface{} {
		var a []interface{}
		for _, v := range arr {
			if filterEven(v) {
				a = append(a, v)
			}
		}
		return a
	}(array)

	assertArray(t, actual, result)
}

// TestReduceFn :)
func TestReduceFn(t *testing.T) {

	var a []interface{}
	a = createArray(100)

	add := func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	}

	res := Reduce(a, add, nil)
	actual := func() interface{} {
		var result interface{}
		result = 0
		for _, i := range a {
			result = result.(int) + i.(int)
		}
		return result
	}()

	if res.(int) != actual.(int) {
		t.Errorf(`Expected %d, got %d`, res.(int), actual.(int))
	}

	res = res.(int) + 1

	if res.(int) == actual.(int) {
		t.Errorf(`Expected %d, got %d`, res.(int), actual.(int))
	}
}

func TestComposeFn(t *testing.T) {
	f := double
	g := square

	arr := createArray(1000)
	comp := Compose(f, g)
	result := MapFn(arr, comp)

	actual := func() []interface{} {
		var r []interface{}
		for _, val := range arr {
			r = append(r, double(square(val)))
		}
		return r
	}()

	assertArray(t, actual, result)

}

func TestParMapFn(t *testing.T) {
	f := double
	g := square

	arr := createArray(1000)

	comp := Pipe(f, g)
	threads := 4

	result := ParMap(arr, comp, threads)
	actual := func() []interface{} {
		var r []interface{}
		for _, val := range arr {
			r = append(r, comp(val))
		}
		return r
	}()

	assertArray(t, actual, result)
}

func TestPipe(t *testing.T) {
	f := double
	g := square

	arr := createArray(1000)
	pipe := Pipe(f, g)

	result := MapFn(arr, pipe)

	actual := func() []interface{} {
		var r []interface{}
		for _, val := range arr {
			r = append(r, square(double(val)))
		}
		return r
	}()

	assertArray(t, actual, result)
}

func TestReverse(t *testing.T) {
	var arr []interface{}
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	actual := func() []interface{} {
		var values []interface{}
		for i := 9; i >= 0; i-- {
			values = append(values, i)
		}
		return values
	}()

	reverse(arr)

	assertArray(t, actual, arr)
}

func createArray(size int) []interface{} {
	var result []interface{}
	for i := 0; i < size; i++ {
		result = append(result, i)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
