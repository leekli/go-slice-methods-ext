package methods

import (
	"fmt"
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

func TestMap_StructValueType(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	input := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Carol", Age: 40},
	}

	output := Map(input, func(p Person, i int) Person {
		p.Age += 1
		p.Name = p.Name + "_senior"
		return p
	})

	assert.Equal(t, 3, len(output))
	assert.Equal(t, "Alice_senior", output[0].Name)
	assert.Equal(t, 31, output[0].Age)
	assert.Equal(t, "Bob_senior", output[1].Name)
	assert.Equal(t, 26, output[1].Age)
	assert.Equal(t, "Carol_senior", output[2].Name)
	assert.Equal(t, 41, output[2].Age)

	// Original must remain unchanged
	assert.Equal(t, "Alice", input[0].Name)
	assert.Equal(t, 30, input[0].Age)
}

func TestMap_StructPointerType(t *testing.T) {
	type Point struct {
		X, Y float64
	}

	input := []*Point{
		{X: 1.0, Y: 2.0},
		{X: 3.0, Y: 4.0},
	}

	output := Map(input, func(p *Point, i int) *Point {
		p.X *= 2
		p.Y *= 2

		return p
	})

	assert.Equal(t, 2, len(output))
	assert.Equal(t, 2.0, output[0].X)
	assert.Equal(t, 4.0, output[0].Y)
	assert.Equal(t, 6.0, output[1].X)
	assert.Equal(t, 8.0, output[1].Y)

	// Original *should* be mutated due to pointer use
	assert.Equal(t, 2.0, input[0].X)
	assert.Equal(t, 4.0, input[0].Y)
}

func TestMap_FromFixedArray(t *testing.T) {
	arr := [4]int{10, 20, 30, 40}
	slice := arr[:] // Convert array to slice

	output := Map(slice, func(v int, i int) int {
		return v + i*100
	})

	assert.Equal(t, []int{10, 120, 230, 340}, output)
	assert.Equal(t, [4]int{10, 20, 30, 40}, arr) // Original array not mutated
}

func TestMap_NestedSlices(t *testing.T) {
	input := [][]string{
		{"a", "b"},
		{"c"},
		{},
		{"d", "e", "f"},
	}

	output := Map(input, func(s []string, i int) []string {
		return append([]string{}, append(s, fmt.Sprintf("idx-%d", i))...)
	})

	assert.Equal(t, 4, len(output))
	assert.Equal(t, []string{"a", "b", "idx-0"}, output[0])
	assert.Equal(t, []string{"c", "idx-1"}, output[1])
	assert.Equal(t, []string{"idx-2"}, output[2])
	assert.Equal(t, []string{"d", "e", "f", "idx-3"}, output[3])
}

func TestMap_CustomTypeAndComplex(t *testing.T) {
	type ID int

	inputIDs := []ID{1, 2, 3}

	outIDs := Map(inputIDs, func(id ID, i int) ID {
		return id * 10
	})

	assert.Equal(t, []ID{10, 20, 30}, outIDs)

	inputC := []complex128{1 + 2i, 3 + 4i}

	outC := Map(inputC, func(c complex128, i int) complex128 {
		return c * 2
	})

	assert.Equal(t, []complex128{2 + 4i, 6 + 8i}, outC)
}

func TestMap_InterfaceSlice(t *testing.T) {
	input := []any{1, "two", 3.0, true}

	output := Map(input, func(v any, i int) any {
		switch val := v.(type) {
		case int:
			return val * 10
		case string:
			return val + "!"
		case float64:
			return val * 2
		case bool:
			return !val
		default:
			return v
		}
	})

	assert.Equal(t, []any{10, "two!", 6.0, false}, output)
}

func TestMap_MapAsElement(t *testing.T) {
	input := []map[string]int{
		{"a": 1},
		{"b": 2, "c": 3},
	}

	output := Map(input, func(m map[string]int, i int) map[string]int {
		newM := make(map[string]int, len(m)+1)
		for k, v := range m {
			newM[k] = v
		}
		newM["index"] = i

		return newM
	})

	assert.Equal(t, 2, len(output))
	assert.Equal(t, 1, output[0]["a"])
	assert.Equal(t, 0, output[0]["index"])
	assert.Equal(t, 2, output[1]["b"])
	assert.Equal(t, 3, output[1]["c"])
	assert.Equal(t, 1, output[1]["index"])
}

func TestMap_LargeSliceAndIndexUsage(t *testing.T) {
	const n = 10_000

	input := make([]int, n)

	for i := range input {
		input[i] = i
	}

	output := Map(input, func(v int, i int) int {
		return v + i
	})

	assert.Equal(t, n, len(output))
	assert.Equal(t, 0, output[0])
	assert.Equal(t, 1+1, output[1])
	assert.Equal(t, (n-1)+(n-1), output[n-1])
}

func TestMap_ZeroValuesAndNilElements(t *testing.T) {
	type Node struct {
		Val *int
	}

	input := []Node{
		{},
		{Val: new(int)},
	}

	*input[1].Val = 42

	output := Map(input, func(n Node, i int) Node {
		if n.Val == nil {
			return Node{}
		}
		v := *n.Val * 2

		return Node{Val: &v}
	})

	assert.Nil(t, output[0].Val)
	assert.NotNil(t, output[1].Val)
	assert.Equal(t, 84, *output[1].Val)
}

func TestMap_IdentityPreservesContent(t *testing.T) {
	input := []string{"x", "y", "z"}

	output := Map(input, func(s string, _ int) string { return s })
	
	assert.Equal(t, input, output)
	assert.NotSame(t, &input, &output)
}