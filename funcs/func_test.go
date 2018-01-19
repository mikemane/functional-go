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

	if !(assertArray(result, actual)) {
		t.Errorf(`Testing failed values are not equal.`)
	}

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

func assertArray(a, b []interface{}) bool {
	n := len(a)
	if n != len(b) {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i].(int) != b[i].(int) {
			// fmt.Printf("%v != %v at index %v\n", a[i], b[i], i)
			return false
		}
	}
	return true
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
	}

	if !assertArray(result, actual(array)) {
		t.Error(`Arrays are not equal`)
	}
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

	if !assertArray(result, actual) {
		t.Errorf(
			`Values are not equal %v, %v `, result[:10], actual[:10],
		)
	}

}

type Dummy []interface{}

func (d Dummy) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Dummy) Len() int {
	return len(d)
}

func (d Dummy) Less(i, j int) bool {
	return d[i].(int) < d[j].(int)
}

func TestParMapFn(t *testing.T) {
	f := double
	g := square

	arr := createArray(1000)

	comp := Pipe(f, g)

	result := MapFn(arr, comp)

	actual := func() []interface{} {
		var r []interface{}
		for _, val := range arr {
			r = append(r, double(square(val)))
		}
		return r
	}()

	if !assertArray(result, actual) {
		t.Error(`Values are not equal`)
	}

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

	if !assertArray(result, actual) {
		t.Errorf("Expected %v got %v", result, actual)
	}
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

	if !assertArray(actual, arr) {
		t.Errorf("Expected %v got %v", actual, arr)
	}
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
