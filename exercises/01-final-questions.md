# Pack 01 — Final Questions

Use these questions to test your understanding after completing all exercises in `01-setup-and-fundamentals.md`. Answer each one directly in this file, without looking at the exercise pack first, then verify your answer by writing runnable code.

---

## Conceptual Questions

**1.** What is the practical difference between `go vet` and `go test`? Give an example of something each one catches that the other does not.

**2.** What is a "zero value" in Go? What is the zero value of `int`, `string`, `bool`, a pointer type, and a slice?

**3.** Why does Go prefer returning `(T, error)` over throwing/catching exceptions for expected failure cases? What would `Classify(score int) (string, error)` look like if it panicked instead — what would break for callers?

**4.** What is the difference between a named return and a regular return? What danger does a named return introduce with `defer`?

**5.** What does it mean that Go passes everything "by value", including pointers? If you pass a `*int` to a function and the function reassigns the pointer itself (not `*p`), does the caller see that change?

**6.** What is a closure? Give an example where a closure captures a variable by reference in a way that surprises a beginner (e.g. capturing a loop variable).

**7.** Why does `gofmt -w .` producing no further changes matter as an acceptance criterion, instead of just "the code compiles"?

---

## Code Prediction Questions

**8.** What does this print, and why?

```go
func bump(score *int, delta int) {
    if score == nil {
        return
    }
    *score += delta
}

func main() {
    var scores []int
    for _, s := range []int{10, 20, 30} {
        scores = append(scores, s)
    }
    bump(&scores[0], 5)
    fmt.Println(scores)
}
```

**9.** What is printed by this loop, and why? (Go version matters — state which behavior you'd expect on Go 1.22+.)

```go
funcs := make([]func(), 0, 3)
for i := 0; i < 3; i++ {
    funcs = append(funcs, func() {
        fmt.Println(i)
    })
}
for _, f := range funcs {
    f()
}
```

**10.** Does this compile? If not, what is the exact compiler error category, and why does Go enforce it?

```go
func main() {
    var x int
    fmt.Println("done")
}
```

**11.** What is printed, and in what order? Explain the role of `defer` here.

```go
func process() {
    fmt.Println("start")
    defer fmt.Println("cleanup 1")
    defer fmt.Println("cleanup 2")
    fmt.Println("end")
}
```

**12.** Given:

```go
func classify(score int) (result string, err error) {
    defer func() {
        if result == "" {
            err = fmt.Errorf("no classification produced")
        }
    }()

    if score < 0 || score > 100 {
        return "", fmt.Errorf("invalid score: %d", score)
    }
    return "pass", nil
}
```

If `score` is `150`, what are the final values of `result` and `err` when the function returns? Does the deferred function's check ever actually fire in this version? Why or why not?

**13.** What is the output? What does this reveal about variadic functions and slices?

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    values := []int{1, 2, 3}
    fmt.Println(sum(values...))
    fmt.Println(sum(1, 2, 3, 4))
}
```

---

## Implementation Challenges

**14.** Write a function `SafeDivide(a, b float64) (float64, error)` that returns an error instead of producing `+Inf`/`NaN`/panicking when `b == 0`. Write a table-driven test covering: normal division, division by zero, division of zero by a non-zero number.

**15.** Write a function `Retry(attempts int, fn func() error) error` that calls `fn` up to `attempts` times, returning `nil` on the first success, or the last error seen if every attempt fails. Use a closure in your test to count how many times `fn` was actually invoked.

**16.** Write a function `ClampAll(values []int, min, max int) []int` that returns a **new** slice where every value is clamped between `min` and `max`, without mutating the input slice. Write a test that explicitly asserts the original slice is unchanged after the call — this is also a pointer/aliasing exercise, not just a clamping exercise.
