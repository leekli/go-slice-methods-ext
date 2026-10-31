package methods

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce_ExecutesSuccessfully_EmptySlice(test *testing.T) {
	input := []int{}
	callback := func(accumulator int, currentValue int, index int) int { return 0 }
	initialValue := 0

	Reduce(input, callback, initialValue)

	assert.NotPanics(test, func() { Reduce(input, callback, initialValue) })
}

func TestReduce_ReducesToSingleValue_SumSlice(test *testing.T) {
	input := []int{1, 2, 3, 4}
	callback := func(accumulator int, currentValue int, index int) int { return accumulator + currentValue }
	initialValue := 0

	output := Reduce(input, callback, initialValue)

	assert.NotPanics(test, func() { Reduce(input, callback, initialValue) })
	assert.Equal(test, 10, output)
}

func TestReduce_ReducesToSingleValue_MultiplyNumbers(test *testing.T) {
	input := []int{1, 2, 3, 4}
	callback := func(accumulator int, currentValue int, index int) int { return accumulator * currentValue }
	initialValue := 1

	output := Reduce(input, callback, initialValue)

	assert.NotPanics(test, func() { Reduce(input, callback, initialValue) })
	assert.Equal(test, 24, output)
}

func TestReduce_ReducesToSingleValue_FindMaxNumber(test *testing.T) {
	input := []int{10, 5, 30, 20}
	callback := func(accumulator int, currentValue int, index int) int { return slices.Max(input) }
	initialValue := 0

	output := Reduce(input, callback, initialValue)

	assert.NotPanics(test, func() { Reduce(input, callback, initialValue) })
	assert.Equal(test, 30, output)
}

func TestReduce_ReducesToSingleValue_FindMinNumber(test *testing.T) {
	input := []int{10, 5, 30, 20}
	callback := func(accumulator int, currentValue int, index int) int { return slices.Min(input) }
	initialValue := 0

	output := Reduce(input, callback, initialValue)

	assert.NotPanics(test, func() { Reduce(input, callback, initialValue) })
	assert.Equal(test, 5, output)
}