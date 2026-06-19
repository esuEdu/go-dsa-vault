# Go Cheat Sheet for LeetCode

**Tags:** #go #syntax #fundamentals

## Maps (Hash Maps)

```go
// Create
m := make(map[int]int)
m := map[string]bool{"a": true}

// Insert / update
m[key] = value

// Check existence (comma-ok idiom)
val, ok := m[key]
if !ok { /* key missing */ }

// Delete
delete(m, key)

// Iterate
for k, v := range m { }
```

## Slices (Dynamic Arrays)

```go
s := make([]int, 0)          // empty, len=0
s := make([]int, n)          // zero-filled, len=n
s := make([]int, 0, n)       // pre-allocated capacity n

s = append(s, 1, 2, 3)
s = append(s, other...)      // spread another slice

// Sub-slice (shares backing array!)
sub := s[1:4]                // indices 1,2,3

// Copy to avoid aliasing
dst := make([]int, len(s))
copy(dst, s)
```

## Sorting

```go
import "sort"

sort.Ints(nums)               // ascending in-place
sort.Sort(sort.Reverse(sort.IntSlice(nums)))  // descending

// Custom comparator
sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})
```

## Strings

```go
import "strings"

s := "hello"
r := []rune(s)               // iterate over Unicode chars safely
b := []byte(s)               // iterate over bytes

strings.Contains(s, "ell")
strings.Split(s, ",")
strings.Join(parts, "-")
strconv.Itoa(42)             // int → string
strconv.Atoi("42")           // string → int, returns (int, error)
```

## Math helpers

```go
import "math"

math.MaxInt   // max int value (use for initializing min trackers)
math.MinInt   // min int value
math.Abs(x)   // only works on float64 — for int: use if x < 0 { x = -x }
math.Max(a, b)  // float64 only — write your own for int
```

## Pointers

```go
x := 42
p := &x       // pointer to x
*p = 99       // dereference to mutate

// Linked list node pattern
type ListNode struct {
    Val  int
    Next *ListNode
}
```
