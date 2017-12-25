package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mikemane/higherorder/funcs"
)

// Point test
type Point struct {
	x int
	y int
}

// NewPoint point
func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) scale(v int) Point {
	return Point{
		x: p.x * v,
		y: p.y * v,
	}
}

func main() {
	fmt.Println("Mikemane is the name")
	var numbers []interface{}
	for i := 0; i < 100000000; i++ {
		numbers = append(numbers, rand.Intn(1000))
	}

	composer := funcs.Compose(double, square)

	start := time.Now()
	// result := funcs.MapFn(numbers, composer)
	result := funcs.ParMap(numbers, composer, 10)
	end := time.Since(start)
	fmt.Println(end)
	sum := funcs.Reduce(result, add, nil)
	fmt.Println(sum)

}

func add(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func double(a interface{}) interface{} {
	return a.(int) * 2
}

func square(a interface{}) interface{} {
	return a.(int) * a.(int)
}
