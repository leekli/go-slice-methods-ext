package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_ExecutesSuccessfully_EmptySlice(test *testing.T) {
	input := []int{}
	callback := func(elem int, index int) int { return elem }

	output := Map(input, callback)

	assert.NotPanics(test, func() { Map(input, callback) })
	assert.Equal(test, 0, len(output))
}

func TestMap_ExecutesSuccessfully_NonEmptySlice_VaryingTypes(test *testing.T) {
	stringSlice := []string{"Hello", "World"}
	stringCallback := func(elem string, index int) string { return elem }

	output := Map(stringSlice, stringCallback)

	assert.NotPanics(test, func() { Map(stringSlice, stringCallback) })
	assert.Equal(test, 2, len(output))

	intSlice := []int{1, 2, 3}
	intCallback := func(elem int, index int) int { return elem }

	intOutput := Map(intSlice, intCallback)

	assert.NotPanics(test, func() { Map(intSlice, intCallback) })
	assert.Equal(test, 3, len(intOutput))

	floatSlice := []float32{0.9, 1.344, 4234, 22}
	floatCallback := func(elem float32, index int) float32 { return elem }

	floatOutput := Map(floatSlice, floatCallback)

	assert.NotPanics(test, func() { Map(floatSlice, floatCallback) })
	assert.Equal(test, 4, len(floatOutput))

	boolSlice := []bool{true, false}
	boolCallback := func(elem bool, index int) bool { return elem }

	boolOutput := Map(boolSlice, boolCallback)

	assert.NotPanics(test, func() { Map(boolSlice, boolCallback) })
	assert.Equal(test, 2, len(boolOutput))
}

func TestMap_InvokesCallbackFunction_ForCorrectNumOfSliceElementTimes(test *testing.T) {
	input := []string{"Hello", "World"}
	count := 0
	callback := func(elem string, index int) string {
		count++

		return elem
	}

	output := Map(input, callback)

	assert.Equal(test, count, len(output))
}

func TestMap_IterationPassesInCorrectOrder(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}
	var collectedElements []string
	var collectedIndices []int

	callback := func(elem string, index int) string {
		collectedElements = append(collectedElements, elem)
		collectedIndices = append(collectedIndices, index)

		return elem
	}

	output := Map(input, callback)

	assert.Equal(test, input, collectedElements)
	assert.Equal(test, []int{0, 1, 2}, collectedIndices)
	assert.Equal(test, len(collectedElements), len(output))
}

func TestMap_EmptyAndNilSlices(test *testing.T) {
	// Empty slice
	input := []int{}
	emptyCount := 0

	callback := func(elem int, index int) int {
		emptyCount++

		return elem
	}

	Map(input, callback)

	assert.Equal(test, 0, emptyCount)

	// nil slice (should not panic)
	var nilSlice []int
	nilCount := 0

	assert.NotPanics(test, func() {
		Map(nilSlice, func(elem int, index int) int {
			nilCount++

			return elem
		})
	})

	assert.Equal(test, 0, nilCount)
}

func TestMap_NormalMutation(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}

	output := Map(input, func(elem string, index int) string {
		newString := elem + "lay"

		return newString
	})

	assert.Equal(test, "Applelay", output[0])
	assert.Equal(test, "Bananalay", output[1])
	assert.Equal(test, "Cherrylay", output[2])
	assert.Equal(test, len(input), len(output))
}

func TestMap_PointerMutation(test *testing.T) {
	type User struct {
		Name string
	}

	users := []*User{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	// Mutate fields through pointer
	output := Map(users, func(u *User, index int) *User {
		u.Name = u.Name + "_updated"

		return u
	})

	assert.Equal(test, "Alice_updated", users[0].Name)
	assert.Equal(test, "Bob_updated", users[1].Name)
	assert.Equal(test, len(users), len(output))
}

func TestMap_ShouldNotMutateOriginalInput(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}

	// Mutate fields through pointer
	output := Map(input, func(elem string, index int) string {
		newString := elem + "lay"

		return newString
	})

	assert.Equal(test, "Apple", input[0])
	assert.Equal(test, "Banana", input[1])
	assert.Equal(test, "Cherry", input[2])
	assert.Equal(test, len(input), len(output))
}

func TestMap_TwoSliceReferencesNotTheSame(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}

	output := Map(input, func(elem string, index int) string {
		return elem
	})

	assert.Equal(test, len(input), len(output))
	assert.NotSame(test, &input, &output)
}

func TestMap_PanicPropagation(test *testing.T) {
	input := []int{1, 2, 3}

	// Verify that a panic inside the callback propagates correctly
	assert.Panics(test, func() {
		Map(input, func(elem int, index int) int {
			if index == 1 {
				panic("callback failure")
			}

			return elem
		})
	})
}
