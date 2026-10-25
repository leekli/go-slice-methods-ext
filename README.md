# go-slice-methods-ext

A small, zero-dependency package that brings the familiar JavaScript-style Array Methods such as `ForEach`, `Map` and `Filter` to Go slices.

Inspired by the [MDN Array methods](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array).

## Overview

This package provides generic helpers that mirror the most commonly used JavaScript Array methods for the following (API reference below):

- **ForEach** – run a function once for every element
- **Map** – transform every element into a new value (same type)
- **Filter** – keep only the elements that pass a test

They work with any slice type using Go generics.

> **Important note about `Map`**  
> Unlike JavaScript’s `Array.prototype.map`, this implementation always returns a slice of the **same** element type.

## Installation

```bash
go get github.com/leekli/go-slice-methods-ext
```

Then import it:

```go
import "github.com/leekli/go-slice-methods-ext"
```

---

## Tests

To run the tests:

```bash
go test ./...
```

---

## API References

### ForEach

```go
func ForEach[T any](slice []T, callback func(elem T, index int))
```

Runs the provided callback once for every element in the slice.

The callback receives the element and its index.

**Example**

```go
names := []string{"Alice", "Bob", "Carol"}

methods.ForEach(names, func(name string, i int) {
    fmt.Printf("%d → %s\n", i, name)
})

// Output prints each index number & element value
```

### Map

```go
func Map[T any](slice []T, callback func(elem T, index int) T) []T
```

Creates a new slice by calling the callback on every element.

The new slice has the same length and the same element type.

**Example**

```go
numbers := []int{1, 2, 3, 4}

doubled := methods.Map(numbers, func(n int, _ int) int {
    return n * 2
})

// doubled is []int{2, 4, 6, 8}
```

### Filter

```go
func Filter[T any](slice []T, callback func(elem T, index int) bool) []T
```

Creates a new slice containing only the elements for which the callback returns `true` (a predicate function).

**Example**

```go
ages := []int{12, 17, 25, 9, 42}

adults := methods.Filter(ages, func(age int, _ int) bool {
    return age >= 18
})

// adults is []int{25, 42}
```

---

## Behaviour Notes

| Situation                             | What Happens                                    |
| ------------------------------------- | ----------------------------------------------- |
| `nil` slice                           | Treated as empty – no panic, empty result       |
| Empty slice                           | Callback is never called, empty result returned |
| Panic inside callback                 | The panic propagates to the caller              |
| Value types (structs, ints, strings…) | Original slice is not modified                  |
| Pointer types                         | Original elements are modified                  |
| Index                                 | Always starts at 0 and increases by 1           |

---

## Further Examples (Varying Types)

### Structs

```go
type User struct {
    Name string
    Age  int
}

users := []User{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 17},
    {Name: "Carol", Age: 25},
}

// Keep only adults
adults := methods.Filter(users, func(u User, _ int) bool {
    return u.Age >= 18
})

// Upper-case their names
upper := methods.Map(adults, func(u User, _ int) User {
    u.Name = strings.ToUpper(u.Name)
    return u
})
```

### Pointer slices

```go
type Counter struct {
    N int
}

counters := []*Counter{{N: 1}, {N: 2}, {N: 3}}

methods.ForEach(counters, func(c *Counter, _ int) {
    c.N *= 10
})

// counters now contain 10, 20, 30
```

### Converting a fixed-size array to a slice

```go
arr := [4]string{"a", "b", "c", "d"}

upper := methods.Map(arr[:], func(s string, _ int) string {
    return strings.ToUpper(s)
})

// upper is []string{"A", "B", "C", "D"}
```

### Simple chaining

```go
result := methods.Map(
    methods.Filter(numbers, func(n int, _ int) bool {
        return n%2 == 0
    }),
    func(n int, _ int) int {
        return n * 10
    },
)
```
