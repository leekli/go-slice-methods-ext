package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAt_ExecutesSuccessfully(test *testing.T) {
	input := []int{1}
	index := 0

	At(input, 0)

	assert.NotPanics(test, func() { At(input, index) })
}

func TestAt_ReturnsElementAndExistence_AtGivenIndex_SingleElement(test *testing.T) {
		index := 0

		intInput := []int{2}

		intValue, intExists := At(intInput, index)
		assert.Equal(test, true, intExists)
		assert.Equal(test, 2, intValue)

		strInput := []string{"A"}

		strValue, strExists := At(strInput, index)
		assert.Equal(test, true, strExists)
		assert.Equal(test, "A", strValue)

		boolInput := []bool{true}

		boolValue, boolExists := At(boolInput, index)
		assert.Equal(test, true, boolExists)
		assert.Equal(test, true, boolValue)
}

func TestAt_ReturnsElementAndExistence_AtGivenIndex_MultipleElements(test *testing.T) {
		intInput := []int{24, 378, 221, -1, 33, 59}
		intIndex := 3

		intValue, intExists := At(intInput, intIndex)
		assert.Equal(test, true, intExists)
		assert.Equal(test, -1, intValue)

		strInput := []string{"Hello", "World", "I", "Am", "A", "String"}
		strIndex := 2

		strValue, strExists := At(strInput, strIndex)
		assert.Equal(test, true, strExists)
		assert.Equal(test, "I", strValue)


		boolInput := []bool{true, false, false, true, true, true, false, true, false, false}
		boolIndex := 8

		boolValue, boolExists := At(boolInput, boolIndex)
		assert.Equal(test, true, boolExists)
		assert.Equal(test, false, boolValue)
}

func TestAt_ReturnsFalseExistence_IfSliceIsEmpty(test *testing.T) {
	input := []int{}
	index := 0

	_, exists := At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, false, exists)
}

func TestAt_ReturnsFalseExistence_IfIndexIsLargerThanLengthOfSlice(test *testing.T) {
	input := []string{"A", "B", "C"}
	index := 3

	_, exists := At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, false, exists)
}

func TestAt_ReturnsTrueExistenceAndValue_WhenGivenANegativeIndex_CountsBackFromLastItem(test *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := -1

	value, exists := At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, true, exists)
	assert.Equal(test, 5, value)

	index = -2

	value, exists = At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, true, exists)
	assert.Equal(test, 4, value)

	index = -3

	value, exists = At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, true, exists)
	assert.Equal(test, 3, value)

	index = -4

	value, exists = At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, true, exists)
	assert.Equal(test, 2, value)

	index = -5

	value, exists = At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, true, exists)
	assert.Equal(test, 1, value)
}

func TestAt_ReturnsFalseExistence_WhenGivenANegativeIndex_AndItsBeyondSliceLen(test *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := -6

	_, exists := At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, false, exists)

	index = -99

	_, exists = At(input, index)

	assert.NotPanics(test, func() { At(input, index) })
	assert.Equal(test, false, exists)
}


