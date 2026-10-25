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

func TestFilter_LargerCallbackFunction_MoreRoutes(test *testing.T) {
	input := []int{-3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	callback := func(elem int, index int) bool {
		if elem < 2 {
			return false
		}

		if elem%2 == 0 {
			return elem == 2
		}

		for factor := 3; factor*factor <= elem; factor += 2 {
			if elem%factor == 0 {
				return false
			}
		}
		return true
	}

	output := Filter(input, callback)

	assert.Equal(test, []int{2, 3, 5, 7, 11, 13}, output)
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

func TestFilter_StructValueType(t *testing.T) {
	type Product struct {
		Name  string
		Price float64
		Stock int
	}

	input := []Product{
		{Name: "A", Price: 10.0, Stock: 5},
		{Name: "B", Price: 25.0, Stock: 0},
		{Name: "C", Price: 15.0, Stock: 3},
		{Name: "D", Price: 5.0, Stock: 10},
	}

	output := Filter(input, func(p Product, i int) bool {
		return p.Price >= 10.0 && p.Stock > 0
	})

	assert.Equal(t, 2, len(output))
	assert.Equal(t, "A", output[0].Name)
	assert.Equal(t, "C", output[1].Name)
}

func TestFilter_StructPointerType(t *testing.T) {
	type Item struct {
		ID    int
		Active bool
	}

	input := []*Item{
		{ID: 1, Active: true},
		{ID: 2, Active: false},
		{ID: 3, Active: true},
		nil,
	}

	output := Filter(input, func(it *Item, i int) bool {
		return it != nil && it.Active
	})

	assert.Equal(t, 2, len(output))
	assert.Equal(t, 1, output[0].ID)
	assert.Equal(t, 3, output[1].ID)
}

func TestFilter_FromFixedArray(t *testing.T) {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	slice := arr[:]

	output := Filter(slice, func(v int, i int) bool {
		return v%2 == 0 && i < 5
	})

	assert.Equal(t, []int{2, 4}, output)
	assert.Equal(t, [6]int{1, 2, 3, 4, 5, 6}, arr)
}

func TestFilter_NestedSlices(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4},
		{},
		{5, 6, 7, 8},
		{9},
	}

	output := Filter(input, func(s []int, i int) bool {
		return len(s) >= 2
	})

	assert.Equal(t, 2, len(output))
	assert.Equal(t, []int{1, 2, 3}, output[0])
	assert.Equal(t, []int{5, 6, 7, 8}, output[1])
}

func TestFilter_CustomTypeAndComplex(t *testing.T) {
	type Score float64

	input := []Score{1.5, 3.0, 0.5, 4.2, 2.8}

	output := Filter(input, func(s Score, i int) bool {
		return s >= 2.5
	})

	assert.Equal(t, []Score{3.0, 4.2, 2.8}, output)

	inputC := []complex64{1 + 1i, 2 + 0i, 0 + 3i, 4 + 4i}

	outC := Filter(inputC, func(c complex64, i int) bool {
		return real(c) > 1
	})

	assert.Equal(t, []complex64{2 + 0i, 4 + 4i}, outC)
}

func TestFilter_InterfaceSlice(t *testing.T) {
	input := []any{nil, 42, "hello", false, 3.14, struct{}{}}

	output := Filter(input, func(v any, i int) bool {
		switch v.(type) {
		case int, float64:
			return true
		default:
			return false
		}
	})

	assert.Equal(t, []any{42, 3.14}, output)
}

func TestFilter_MapAsElement(t *testing.T) {
	input := []map[string]bool{
		{"ok": true},
		{"ok": false},
		{},
		{"ok": true, "extra": false},
	}

	output := Filter(input, func(m map[string]bool, i int) bool {
		return m["ok"]
	})

	assert.Equal(t, 2, len(output))
	assert.True(t, output[0]["ok"])
	assert.True(t, output[1]["ok"])
}

func TestFilter_AllPassAllFailAndIndex(t *testing.T) {
	input := []string{"a", "bb", "ccc", "dddd"}

	all := Filter(input, func(_ string, _ int) bool { return true })
	assert.Equal(t, input, all)

	none := Filter(input, func(_ string, _ int) bool { return false })
	assert.Equal(t, 0, len(none))

	evenIdx := Filter(input, func(_ string, i int) bool { return i % 2 == 0 })
	assert.Equal(t, []string{"a", "ccc"}, evenIdx)
}

func TestFilter_LargeSlice(t *testing.T) {
	const n = 20_000
	input := make([]int, n)
	for i := range input {
		input[i] = i
	}

	output := Filter(input, func(v int, _ int) bool {
		return v%1000 == 0
	})

	assert.Equal(t, 20, len(output))
	assert.Equal(t, 0, output[0])
	assert.Equal(t, 19000, output[19])
}

func TestFilter_NilElementsAndZeroValues(t *testing.T) {
	input := []*string{nil, new(string), nil}
	*input[1] = "keep"

	output := Filter(input, func(s *string, _ int) bool {
		return s != nil
	})

	assert.Equal(t, 1, len(output))
	assert.Equal(t, "keep", *output[0])
}
