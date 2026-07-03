# Pack 02 — Final Questions

Use these questions to test your understanding after completing all exercises in `02-core-data-structures.md`. Answer each one directly in this file, without looking at the exercise pack first, then verify your answer by writing runnable code.

---

## Conceptual Questions

**1.** A slice header has three components. What are they, and which one determines whether `append` can grow in place versus must allocate a new backing array?

**2.** Why does writing to a nil map panic, but reading from a nil map does not? What does a read return in that case?

**3.** What does "method set" mean in Go, and why does it determine whether a value of type `T` (not `*T`) can be assigned to an interface variable when the method was declared as `func (t *T) Method()`?

**4.** What is the difference between struct embedding in Go and class inheritance in languages like Java or TypeScript? Specifically: does an embedded type know about the outer type that embeds it?

**5.** Why is interface satisfaction in Go called "implicit" or "structural"? What would you have to do differently in a language with explicit interface declarations (e.g. `class Foo implements Bar`)?

**6.** When should a method use a value receiver versus a pointer receiver? Name at least two concrete signals that should push you toward a pointer receiver.

**7.** Why is relying on `for range` order over a `map[string]int` considered a bug waiting to happen, even if it "happens to work" in local testing?

---

## Code Prediction Questions

**8.** What is printed? Explain in terms of backing arrays.

```go
a := make([]int, 3, 5)
a[0], a[1], a[2] = 1, 2, 3

b := a[:2]
b = append(b, 99)

fmt.Println(a)
fmt.Println(b)
```

**9.** Does this compile? If not, what is the category of error, and why?

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d *Dog) Speak() string { return "woof" }

func main() {
    var s Speaker = Dog{}
    fmt.Println(s.Speak())
}
```

**10.** What is printed, and why? Pay attention to whether `Base`'s method is shadowed or actually replaced at the interface-dispatch level.

```go
type Base struct{}

func (Base) Describe() string { return "base" }

type Derived struct {
    Base
}

func (Derived) Describe() string { return "derived" }

func describe(d interface{ Describe() string }) string {
    return d.Describe()
}

func main() {
    d := Derived{}
    fmt.Println(d.Describe())
    fmt.Println(describe(d))
    fmt.Println(d.Base.Describe())
}
```

**11.** Given two independently-created maps with identical key/value pairs, does `fmt.Println` ever print their contents in a different relative order between two separate program runs? What guarantee (if any) does Go make here?

**12.** What is printed? Explain what a nil pointer stored inside an interface value means for a `== nil` comparison.

```go
type Charger interface {
    Charge() error
}

type Card struct{}

func (c *Card) Charge() error { return nil }

func getCharger(valid bool) Charger {
    var c *Card
    if valid {
        c = &Card{}
    }
    return c
}

func main() {
    ch := getCharger(false)
    fmt.Println(ch == nil)
}
```

---

## Implementation Challenges

**13.** Implement a generic-free `Set` type over strings backed by `map[string]struct{}` (not `map[string]bool`) with methods `Add`, `Has`, and `Slice() []string` (sorted). Write a test explaining, in a comment, why `struct{}` is used as the value type instead of `bool`.

**14.** Given the `billing` package from Exercise 2, implement a new type `SplitCharge` that is itself a `Charger`, constructed from a list of other `Charger`s and an amount-splitting strategy, so that calling `.Charge(amount)` on it charges a fraction of `amount` to each underlying charger and rolls back (refunds) any already-succeeded partial charges if a later one fails. Write a test where the second of three chargers fails and assert the first charger's balance is restored.

**15.** Write a function `Merge(maps ...map[string]int) map[string]int` that merges any number of `map[string]int` values, summing values for keys that appear in more than one map, without mutating any of the input maps. Write a test that explicitly asserts none of the original input maps changed after the call.
