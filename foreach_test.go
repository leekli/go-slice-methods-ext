package methods

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach_ExecutesSuccessfully_EmptySlice(test *testing.T) {
	input := []int{}
	callback := func(elem int, index int) {}

	ForEach(input, callback)

	assert.NotPanics(test, func() { ForEach(input, callback) })
}

func TestForEach_ExecutesSuccessfully_NonEmptySlice_VaryingTypes(test *testing.T) {
	stringSlice := []string{"Hello", "World"}
	stringCallback := func(elem string, index int) {}

	ForEach(stringSlice, stringCallback)

	assert.NotPanics(test, func() { ForEach(stringSlice, stringCallback) })

	intSlice := []int{1, 2, 3}
	intCallback := func(elem int, index int) {}

	ForEach(intSlice, intCallback)

	assert.NotPanics(test, func() { ForEach(intSlice, intCallback) })

	floatSlice := []float32{0.9, 1.344, 4234, 22}
	floatCallback := func(elem float32, index int) {}

	ForEach(floatSlice, floatCallback)

	assert.NotPanics(test, func() { ForEach(floatSlice, floatCallback) })

	boolSlice := []bool{true, false}
	boolCallback := func(elem bool, index int) {}

	ForEach(boolSlice, boolCallback)

	assert.NotPanics(test, func() { ForEach(boolSlice, boolCallback) })
}

func TestForEach_InvokesCallbackFunction_ForCorrectNumOfSliceElementTimes(test *testing.T) {
	input := []string{"Hello", "World"}
	count := 0
	callback := func(elem string, index int) {
		count++
	}

	ForEach(input, callback)

	assert.Equal(test, count, len(input))
}

func TestForEach_IterationPassesInCorrectOrder(test *testing.T) {
	input := []string{"Apple", "Banana", "Cherry"}
	var collectedElements []string
	var collectedIndices []int

	callback := func(elem string, index int) {
		collectedElements = append(collectedElements, elem)
		collectedIndices = append(collectedIndices, index)
	}

	ForEach(input, callback)

	assert.Equal(test, input, collectedElements)
	assert.Equal(test, []int{0, 1, 2}, collectedIndices)
}

func TestForEach_EmptyAndNilSlices(test *testing.T) {
	// Empty slice
	input := []int{}
	emptyCount := 0

	callback := func(elem int, index int) {
		emptyCount++
	}

	ForEach(input, callback)

	assert.Equal(test, 0, emptyCount)

	// nil slice (should not panic)
	var nilSlice []int
	nilCount := 0

	assert.NotPanics(test, func() {
		ForEach(nilSlice, func(elem int, index int) {
			nilCount++
		})
	})

	assert.Equal(test, 0, nilCount)
}

func TestForEach_PointerMutation(test *testing.T) {
	type User struct {
		Name string
	}

	users := []*User{
		{Name: "Alice"},
		{Name: "Bob"},
	}

	// Mutate fields through pointer
	ForEach(users, func(u *User, index int) {
		u.Name = u.Name + "_updated"
	})

	assert.Equal(test, "Alice_updated", users[0].Name)
	assert.Equal(test, "Bob_updated", users[1].Name)
}

func TestForEach_PanicPropagation(test *testing.T) {
	input := []int{1, 2, 3}

	// Verify that a panic inside the callback propagates correctly
	assert.Panics(test, func() {
		ForEach(input, func(elem int, index int) {
			if index == 1 {
				panic("callback failure")
			}
		})
	})
}

func TestForEach_CanSuccessfullyCallExternalFunc_PrintLn(test *testing.T) {
	input := []string{"Apple", "Banana"}

	// Create a pipe to intercept stdout
	r, w, err := os.Pipe()

	assert.NoError(test, err)

	os.Stdout = w // Point stdout to our write pipe

	// Execute ForEach using fmt.Println in the callback wrapper
	ForEach(input, func(elem string, index int) {
		fmt.Println(elem)
	})

	// Close the writer
	_ = w.Close()

	// Read all data captured by the read pipe into a buffer
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	assert.NoError(test, err)

	output := buf.String()

	// Assert that fmt.Println was called and printed the expected lines
	assert.Contains(test, output, "Apple\n")
	assert.Contains(test, output, "Banana\n")
}
