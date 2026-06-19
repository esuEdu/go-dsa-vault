# Big O Notation

**Tags:** #complexity #fundamentals

## Time Complexity (common, fastest to slowest)

| Notation     | Name         | Example                          |
|--------------|--------------|----------------------------------|
| O(1)         | Constant     | Hash map lookup                  |
| O(log n)     | Logarithmic  | [[Binary Search]]                |
| O(n)         | Linear       | Single loop over array           |
| O(n log n)   | Linearithmic | [[Merge Sort]], [[Quick Sort]]   |
| O(n²)        | Quadratic    | Nested loops (brute force)       |
| O(2ⁿ)        | Exponential  | [[Backtracking]], subsets        |

## Space Complexity

Think about auxiliary space only (not input size).
- Recursion depth counts → each call frame is O(1) on the stack.
- A hash map that grows with input → O(n).

## Go-specific notes

- `len()` on slices/maps is O(1) — it reads a stored field, not iterating.
- Appending to a slice is **amortized O(1)** but can trigger O(n) copy on resize.
