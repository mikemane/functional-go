package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/mikemane/higherorder/funcs"
)

const MIN = 1000
const MAX = 10000

// In this program we would be comparing the time difference
// Between our functions, and the normal methodd.
func main() {
	limit := 1000000

	fmt.Println("Generate random numbers")
	randomNumbers := generateRandomNumbers(limit)

	fmt.Println("Clone random numbers to generic interface type")
	genericRandom := cloneNumbers(randomNumbers)

	start := time.Now()
	// Perform normal method
	var resultA []int

	// var resultB []interface{}

	for _, i := range randomNumbers {
		resultA = append(resultA, sumNumbersUpTo(i))
	}

	end := time.Since(start)
	fmt.Println("Naive Time taken ", end)
	fmt.Println("Naive Sum is ", sumOfNumbers(resultA))
	fmt.Println()

	start = time.Now()
	// Perform map func method
	parIntNumbers := parInt(randomNumbers, sumNumbersUpTo, 6)
	end = time.Since(start)
	fmt.Println("ParInt: Time taken is", end)
	start = time.Now()
	fmt.Println("ParallelInt sum of ints is ", sumOfNumbers(parIntNumbers))
	fmt.Println()

	// Perform map func method
	mapNumbers := funcs.MapFn(genericRandom, sumUpToGeneric)
	end = time.Since(start)
	fmt.Println("Map: Time taken is", end)
	fmt.Println("MapSum: is", sumNumbersUsingGeneric(mapNumbers))
	fmt.Println()

	start = time.Now()
	// Perform map func method
	parrallelMapNumbers := funcs.ParMap(genericRandom, sumUpToGeneric, 6)
	end = time.Since(start)
	fmt.Println("ParMap: Time taken is", end)
	fmt.Println("ParMapSum is: ", sumNumbersUsingGeneric(parrallelMapNumbers))
	fmt.Println()
}

func generateRandomNumbers(limit int) []int {
	var result []int
	for i := 0; i < limit; i++ {
		result = append(result, randomNumberBetween(MIN, MAX))
	}
	return result
}

func parGenRandom(limit int) []int {
	threads := 1000
	part := limit / threads

	var wg sync.WaitGroup
	values := make([]int, limit)
	start := 0

	for i := 0; i < threads; i++ {
		end := start + part
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				values[i] = randomNumberBetween(1, 100)
			}
		}(start, end)
		start = end
	}

	wg.Wait()
	return values
}

func cloneNumbers(A []int) []interface{} {
	var res []interface{}
	for _, i := range A {
		res = append(res, i)
	}
	return res
}

func randomNumberBetween(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func sumNumbersUpTo(number int) int {
	sum := 0
	for i := 0; i < number; i++ {
		sum += i
	}
	return sum
}

func sumUpToGeneric(value interface{}) interface{} {
	var result interface{}
	result = sumNumbersUpTo(value.(int))
	return result
}

func doubleNative(a int) int {
	// return a * 2
	if a > 0 {
		return int(math.Pow(math.Exp((math.Sqrt(float64(a)))), 5))
	}
	return 0
}

func doubleGeneric(a interface{}) interface{} {
	// return a.(int) * 2
	b := a.(int)
	if b > 0 {
		return int(math.Pow(math.Exp((math.Sqrt(float64(b)))), 5))
	}
	return 0
}

func parInt(A []int, fn func(int) int, threads int) []int {
	var results []int

	done := make(chan interface{})
	valueChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, i := range A {
			valueChan <- i
		}
		close(valueChan)
	}()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range valueChan {
				resultChan <- fn(i)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case <-done:
			return results
		case r, ok := <-resultChan:
			if ok {
				results = append(results, r)
			}
		default:
		}
	}
}

func sumOfNumbers(values []int) int {
	sum := 0
	for _, item := range values {
		sum += item
	}
	return sum
}

func sumNumbersUsingGeneric(numbers []interface{}) interface{} {
	sum := 0
	for _, number := range numbers {
		sum += number.(int)
	}
	return sum
}
