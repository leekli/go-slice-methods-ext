# Go Slice Extension Methods

## Overview

This package implements extended methods to be used on Go Slices of varying types. The aim being to replicate how 'JavaScript Array Methods' work, see [MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array).

---

## Methods available and Usage

1. `ForEach`

- **What:** Executes a provided function once for each slice element.
- **How:**

```go
  stringSlice := []string{"Hello", "World"}
  stringCallback := func(elem string, index int) {
  	fmt.Println(elem)
  }

  ForEach(stringSlice, stringCallback)

  // Prints: "Hello"
  // Prints: "World"
```

2. `Map`

- **What:** Creates a new slice populated with the results of calling a provided function on every element in the calling array.
- **How:**

```go
	input := []string{"Apple", "Banana", "Cherry"}

	output := Map(input, func(elem string, index int) string {
		newString := elem + "-x"

		return newString
	})

	// output is: {"Apple-x", "Banana-x", "Cherry-x"}
```

3. `Filter`

- **What:** Creates a new copy of a portion of a given slice, filtered down to just the elements from the given slice that pass the test implemented by the provided function (a predicate function).
- **How:**

```go
	input := []string{"ABC", "ABCDE"}

	output := Filter(input, func(elem string, index int) bool {
		return len(elem) <= 3
	})

	// output is: {"ABC"}
```
