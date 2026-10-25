package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_ExecutesSuccessfully_EmptySlice(test *testing.T) {
	input := []int{}
	callback := func(elem int, index int) bool { return false }

	output := Filter(input, callback)

	assert.NotPanics(test, func() { Filter(input, callback) })
	assert.Equal(test, 0, len(output))
}

func TestFilter_ExecutesSuccessfully_NonEmptySlice_VaryingTypes(test *testing.T) {
	stringSlice := []string{"ABC", "ABCDE"}
	stringCallback := func(elem string, index int) bool { return len(elem) <= 3 }

	output := Filter(stringSlice, stringCallback)

	assert.NotPanics(test, func() { Filter(stringSlice, stringCallback) })
	assert.Equal(test, 1, len(output))
	assert.Equal(test, "ABC", output[0])

	intSlice := []int{1, 2, 3}
	intCallback := func(elem int, index int) bool { return elem > 1 }

	intOutput := Filter(intSlice, intCallback)

	assert.NotPanics(test, func() { Filter(intSlice, intCallback) })
	assert.Equal(test, 2, len(intOutput))
	assert.Equal(test, 2, intOutput[0])
	assert.Equal(test, 3, intOutput[1])

	floatSlice := []float32{0.9, 1.344, 4234, 22.1}
	floatCallback := func(elem float32, index int) bool { return elem == float32(0.9) || elem == float32(22.1) }

	floatOutput := Filter(floatSlice, floatCallback)

	assert.NotPanics(test, func() { Filter(floatSlice, floatCallback) })
	assert.Equal(test, 2, len(floatOutput))
	assert.Equal(test, float32(0.9), floatOutput[0])
	assert.Equal(test, float32(22.1), floatOutput[1])

	boolSlice := []bool{true, false}
	boolCallback := func(elem bool, index int) bool { return elem == true }

	boolOutput := Filter(boolSlice, boolCallback)

	assert.NotPanics(test, func() { Filter(boolSlice, boolCallback) })
	assert.Equal(test, 1, len(boolOutput))
	assert.Equal(test, true, boolOutput[0])
}

func TestFilter_InvokesCallbackFunction_ForCorrectNumOfSliceElementTimes(test *testing.T) {
	input := []string{"Hello", "World"}
	count := 0
	callback := func(elem string, index int) bool {
		count++

		return false
	}

	Filter(input, callback)

	assert.Equal(test, count, len(input))
}

func TestFilter_IterationPassesInCorrectOrder(test *testing.T) {
	input := []string{"A", "AB", "ABC"}
	var collectedElements []string
	var collectedIndices []int

	callback := func(elem string, index int) bool {
		if elem == "A" || elem == "ABC" {
			collectedElements = append(collectedElements, elem)
			collectedIndices = append(collectedIndices, index)

			return true
		}

		return false
	}

	output := Filter(input, callback)

	assert.Equal(test, []string{"A", "ABC"}, collectedElements)
	assert.Equal(test, []int{0, 2}, collectedIndices)
	assert.Equal(test, 2, len(output))
}

func TestFilter_EmptyAndNilSlices(test *testing.T) {
	// Empty slice
	input := []int{}
	emptyCount := 0

	callback := func(elem int, index int) bool {
		emptyCount++

		return false
	}

	Filter(input, callback)

	assert.Equal(test, 0, emptyCount)

	// nil slice (should not panic)
	var nilSlice []int
	nilCount := 0

	assert.NotPanics(test, func() {
		Filter(nilSlice, func(elem int, index int) bool {
			nilCount++

			return false
		})
	})

	assert.Equal(test, 0, nilCount)
}

func TestFilter_TwoSliceReferencesNotTheSame(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}

	output := Filter(input, func(elem string, index int) bool {
		return true
	})

	assert.Equal(test, len(input), len(output))
	assert.NotSame(test, &input, &output)
}

func TestFilter_ShouldNotMutateOriginalInput(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}

	output := Filter(input, func(elem string, index int) bool {
		return true
	})

	assert.Equal(test, "Apple", input[0])
	assert.Equal(test, "Banana", input[1])
	assert.Equal(test, "Cherry", input[2])
	assert.Equal(test, len(input), len(output))
}

func TestFilter_PanicPropagation(test *testing.T) {
	input := []int{1, 2, 3}

	// Verify that a panic inside the callback propagates correctly
	assert.Panics(test, func() {
		Filter(input, func(elem int, index int) bool {
			if index == 1 {
				panic("callback failure")
			}

			return false
		})
	})
}
