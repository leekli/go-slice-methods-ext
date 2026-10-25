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

func TestForEach_StructValueType(t *testing.T) {
	type Counter struct {
		Value int
	}

	input := []Counter{{1}, {2}, {3}}
	sum := 0

	ForEach(input, func(c Counter, i int) {
		sum += c.Value + i
	})

	assert.Equal(t, 1+0+2+1+3+2, sum)

	// Original input slice is not mutated
	assert.Equal(t, 1, input[0].Value)
}

func TestForEach_StructPointerType(t *testing.T) {
	type Node struct {
		Data string
	}

	input := []*Node{
		{Data: "first"},
		{Data: "second"},
		{Data: "third"},
	}

	ForEach(input, func(n *Node, i int) {
		n.Data = fmt.Sprintf("%s-%d", n.Data, i)
	})

	assert.Equal(t, "first-0", input[0].Data)
	assert.Equal(t, "second-1", input[1].Data)
	assert.Equal(t, "third-2", input[2].Data)
}

func TestForEach_FromFixedArray(t *testing.T) {
	arr := [5]string{"a", "b", "c", "d", "e"}
	slice := arr[:]
	var collected []string

	ForEach(slice, func(s string, i int) {
		collected = append(collected, fmt.Sprintf("%s%d", s, i))
	})

	assert.Equal(t, []string{"a0", "b1", "c2", "d3", "e4"}, collected)
	assert.Equal(t, [5]string{"a", "b", "c", "d", "e"}, arr)
}

func TestForEach_CustomTypeAndComplex(t *testing.T) {
	type Flag bool
	input := []Flag{true, false, true}
	countTrue := 0

	ForEach(input, func(f Flag, _ int) {
		if f {
			countTrue++
		}
	})

	assert.Equal(t, 2, countTrue)

	inputC := []complex128{1 + 1i, 2 + 2i}

	var sum complex128

	ForEach(inputC, func(c complex128, _ int) {
		sum += c
	})

	assert.Equal(t, 3+3i, sum)
}

func TestForEach_InterfaceSlice(t *testing.T) {
	input := []any{10, "hello", 3.14, true}
	var results []string

	ForEach(input, func(v any, i int) {
		results = append(results, fmt.Sprintf("%d:%v", i, v))
	})

	assert.Equal(t, []string{"0:10", "1:hello", "2:3.14", "3:true"}, results)
}

func TestForEach_MapAsElement(t *testing.T) {
	input := []map[string]int{
		{"x": 1},
		{"y": 2, "z": 3},
	}
	total := 0

	ForEach(input, func(m map[string]int, _ int) {
		for _, v := range m {
			total += v
		}
	})

	assert.Equal(t, 6, total)
}

func TestForEach_SideEffectsAndIndex(t *testing.T) {
	input := []int{5, 10, 15, 20}
	var doubled []int
	var indices []int

	ForEach(input, func(v int, i int) {
		doubled = append(doubled, v*2)
		indices = append(indices, i)
	})

	assert.Equal(t, []int{10, 20, 30, 40}, doubled)
	assert.Equal(t, []int{0, 1, 2, 3}, indices)
}

func TestForEach_LargeSlice(t *testing.T) {
	const n = 50_000
	input := make([]int, n)
	for i := range input {
		input[i] = 1
	}

	sum := 0
	ForEach(input, func(v int, _ int) {
		sum += v
	})

	assert.Equal(t, n, sum)
}

func TestForEach_NilElements(t *testing.T) {
	input := []*int{nil, new(int), nil}
	*input[1] = 99
	countNonNil := 0

	ForEach(input, func(p *int, _ int) {
		if p != nil {
			countNonNil++
			assert.Equal(t, 99, *p)
		}
	})

	assert.Equal(t, 1, countNonNil)
}
