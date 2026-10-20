package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach_ExecutesSuccessfully_EmptySlice(test *testing.T) {
	input := Slice[int]{}
	callback := func(elem int) {}

	input.forEach(callback)

	assert.NotPanics(test, func() { input.forEach(callback) })
}

func TestForEach_ExecutesSuccessfully_NonEmptySlice_VaryingTypes(test *testing.T) {
	stringSlice := Slice[string]{"Hello", "World"}
	stringCallback := func(elem string) {}

	stringSlice.forEach(stringCallback)

	assert.NotPanics(test, func() { stringSlice.forEach(stringCallback) })

	intSlice := Slice[int]{1, 2, 3}
	intCallback := func(elem int) {}

	intSlice.forEach(intCallback)

	assert.NotPanics(test, func() { intSlice.forEach(intCallback) })

	floatSlice := Slice[float32]{0.9, 1.344, 4234, 22}
	floatCallback := func(elem float32) {}

	floatSlice.forEach(floatCallback)

	assert.NotPanics(test, func() { floatSlice.forEach(floatCallback) })

	boolSlice := Slice[bool]{true, false}
	boolCallback := func(elem bool) {}

	boolSlice.forEach(boolCallback)

	assert.NotPanics(test, func() { boolSlice.forEach(boolCallback) })
}

func TestForEach_InvokesCallbackFunction_ForCorrectNumOfSliceElementTimes(test *testing.T) {
	input := Slice[string]{"Hello", "World"}
	count := 0
	callback := func(elem string) {
		count++
	}

	expectedCount := 2

	input.forEach(callback)

	assert.Equal(test, expectedCount, count)
}
