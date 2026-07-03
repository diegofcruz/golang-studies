# Exercise Pack 02 - Core Data Structures and Types

This pack covers Topic 2 from TOPICS.md.

## Exercise 1: Arrays, Slices, and Maps Through a Word Counter

### 1. Goal

Build an accurate mental model of slice internals (length, capacity, backing array, aliasing) and map behavior (zero value, nil maps, non-deterministic iteration) by implementing a word-frequency counter.

### 2. Prerequisites

- Pack 01 complete (variables, functions, pointers)
- Comfortable running `go run .` and `go test ./...`

### 3. Step-by-step implementation

1. Create package `wordcount` with:

- `func Tokenize(text string) []string`
- `func Count(words []string) map[string]int`
- `func TopN(counts map[string]int, n int) []string`

Why: splitting tokenizing, counting, and ranking into three functions isolates the slice-heavy step (Tokenize), the map-heavy step (Count), and the step where map iteration order matters (TopN).

2. Implement `Tokenize`:

- Lowercase the input.
- Split on any run of non-letter characters (use `strings.FieldsFunc` with a predicate checking `unicode.IsLetter`).

Why: `strings.Fields`/`FieldsFunc` return a freshly allocated `[]string` — this is your first real slice whose `len` and `cap` you did not choose by hand.

3. Add a temporary debug line printing `len(words)` and `cap(words)` right after calling `Tokenize` on a sample sentence, then remove it once you have observed it.

Why: seeing `cap` be equal to or larger than `len` makes the len/cap/backing-array model concrete before you write any code that depends on it.

4. Implement `Count`:

- Declare `counts := make(map[string]int)`.
- For each word, `counts[word]++`.

Why: `make(map[string]int)` is required — a `var counts map[string]int` (nil map) panics on write. Incrementing a missing key works because Go maps return the zero value (`0` for `int`) for a key that isn't present yet, so `counts[word]++` is safe even the first time.

5. Implement `TopN`:

- Copy the map's keys into a `[]string`.
- Sort that slice by `counts[key]` descending using `sort.Slice`.
- Return the first `n` entries (or fewer, if there aren't `n` distinct words).

Why: map iteration order in Go is intentionally randomized — if you need a stable order, you always sort a slice derived from the map, you never rely on `for range someMap` order directly.

6. In `general/main.go`, call `Tokenize` → `Count` → `TopN` on a sample paragraph and print the top 3 words.

Why: wiring the three functions together in sequence is the same happy-path composition you already practiced with `score` in Pack 01, now with slices and maps instead of scalars.

7. Add tests:

- Table-driven test for `Tokenize` (empty string, punctuation, mixed case, multiple spaces).
- Table-driven test for `Count` (empty slice, all-unique words, repeated words).
- A test for `TopN` where `n` is larger than the number of distinct words.

### 4. Why this works

This exercise forces you to build and consume both core Go collection types end to end: a slice you didn't size by hand (from `Tokenize`), and a map used both for accumulation (`Count`) and as a source you must convert back to a slice for anything order-sensitive (`TopN`). The debug-print step in 3 is deliberate — most slice bugs in real code come from an incorrect mental model of when `append` reuses the backing array versus allocates a new one, and you cannot build that model without looking at `cap`, not just `len`.

### 5. Failure and edge-case drills

- Drill A: Declare `var m map[string]int` (no `make`) and attempt `m["x"] = 1`. Observe the exact panic message and explain why reading from a nil map (`m["x"]`) does *not* panic but writing does.
- Drill B: Take a slice `a := []int{1, 2, 3}`, create `b := a[0:2]`, then run `b = append(b, 99)`. Print `a` afterward. Explain why `a[2]` changed even though you only appended to `b` — this is the aliasing trap: `append` reused the backing array because `cap(b) > len(b)`.
- Drill C: Force a reallocation by appending enough elements to `b` to exceed its capacity, then repeat Drill B's mutation check on `a`. Confirm `a` is now unaffected, and explain why (the backing array was replaced).
- Drill D: Run your `Count` function's output through two `for key := range counts` loops printed back to back, without sorting. Run the program 2-3 times. Observe that the printed order is not guaranteed to be identical every run.

### 6. Questions to verify understanding

1. What are the three components of a slice's internal representation (pointer, len, cap), and what does each one mean?
2. Under what condition does `append` allocate a new backing array instead of reusing the existing one?
3. Why does `counts[word]++` work correctly even when `word` has never been seen before, for a map declared with `make`?
4. Why is it wrong to rely on `for range` order over a map for anything user-visible (e.g. printing a report)?

### 7. Practical challenge (with acceptance criteria)

Challenge: Add a function `Dedupe(words []string) []string` that returns a new slice with duplicate words removed, preserving the first-seen order, without mutating the input slice. Then add `MostCommon(text string) (string, int)` that returns the single most frequent word and its count, using `Tokenize`, `Count`, and `TopN` internally, returning `("", 0)` for empty input.

Acceptance criteria:

- `Dedupe` does not share a backing array with its input where mutation of the result could affect the caller's original slice.
- `Dedupe` preserves first-seen order (not sorted order).
- `MostCommon("")` returns `("", 0)` without panicking.
- A table-driven test covers: normal text, empty text, text where two words tie for most common (documented tie-breaking behavior is acceptable, but must be tested and intentional, not incidental).
- `go test ./...` passes and `gofmt -w .` produces no further changes.

### 8. Best-practice answer key

```go
package wordcount

import (
	"sort"
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	return strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func Count(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}

func TopN(counts map[string]int, n int) []string {
	keys := make([]string, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if counts[keys[i]] != counts[keys[j]] {
			return counts[keys[i]] > counts[keys[j]]
		}
		return keys[i] < keys[j] // stable tie-break: alphabetical
	})

	if n > len(keys) {
		n = len(keys)
	}
	return keys[:n]
}

func Dedupe(words []string) []string {
	seen := make(map[string]bool, len(words))
	result := make([]string, 0, len(words))
	for _, w := range words {
		if seen[w] {
			continue
		}
		seen[w] = true
		result = append(result, w)
	}
	return result
}

func MostCommon(text string) (string, int) {
	words := Tokenize(text)
	if len(words) == 0 {
		return "", 0
	}
	counts := Count(words)
	top := TopN(counts, 1)
	return top[0], counts[top[0]]
}
```

Best practices used:

- `make([]string, 0, len(words))` pre-sizes the result slice's capacity to avoid repeated reallocation during `append`.
- Sorting is always done on a slice derived from map keys — the map itself is never iterated for order-sensitive output.
- Tie-breaking in `TopN` is explicit and deterministic (alphabetical), not left to accidental map order.
- `Dedupe` builds a brand-new backing array via `make` + `append`, so the caller's original slice can never be mutated through the result.

### 9. Key takeaway

A slice is a view (pointer + len + cap) over a backing array that `append` may or may not reuse, and a map's iteration order is never a contract — always sort a derived slice when order matters.

---

## Exercise 2: Structs, Methods, Embedding, and Interfaces Through a Billing System

### 1. Goal

Understand value vs pointer receivers, struct embedding, and implicit interface satisfaction by modeling a small set of payment methods that all need to be "charged" polymorphically.

### 2. Prerequisites

- Exercise 1 in this pack complete
- Comfortable with Go's zero values and pointers from Pack 01

### 3. Step-by-step implementation

1. Create package `billing` with a struct:

```go
type Money struct {
	Cents int64
}
```

and a method `func (m Money) Add(other Money) Money`.

Why: `Money` is a value type on purpose — money should behave like a number (compared and copied by value), so its method uses a **value receiver**: it reads `m` but never needs to mutate the caller's original `Money`.

2. Define an interface:

```go
type Charger interface {
	Charge(amount Money) error
}
```

Why: any type with a `Charge(Money) error` method automatically satisfies `Charger` — there is no `implements` keyword in Go. This is "implicit implementation": satisfaction is structural, checked by the compiler from method sets alone.

3. Create two concrete types that satisfy `Charger`:

```go
type CreditCard struct {
	Balance Money
	Limit   Money
}

type WalletAccount struct {
	Balance Money
}
```

Give `CreditCard` a method `func (c *CreditCard) Charge(amount Money) error` and `WalletAccount` a method `func (w *WalletAccount) Charge(amount Money) error`.

Why: both use a **pointer receiver** here, deliberately different from `Money.Add`. `Charge` must mutate the receiver's balance, so it needs a pointer; if you used a value receiver, you'd be mutating a copy and the caller would never see the balance change.

4. Implement `CreditCard.Charge`: reject (return an error) if `Balance.Cents + amount.Cents > Limit.Cents`; otherwise add to `Balance` and return `nil`.

Implement `WalletAccount.Charge`: reject if `amount.Cents > Balance.Cents` (no overdraft); otherwise subtract from `Balance` and return `nil`.

Why: encoding two different business rules behind the same interface method is exactly what interfaces are for — callers depend on the `Charger` behavior, not on which concrete type they're holding.

5. Create an embedding example:

```go
type AuditedWalletAccount struct {
	WalletAccount
	History []Money
}
```

Override `Charge` on `AuditedWalletAccount` so it calls the embedded `WalletAccount.Charge`, and if it succeeds, appends `amount` to `History`.

Why: embedding promotes `WalletAccount`'s fields and methods onto `AuditedWalletAccount` (you can call `aw.Balance` directly), but defining `Charge` directly on `AuditedWalletAccount` shadows the promoted one — this is Go's version of "overriding", achieved by method resolution order, not inheritance.

6. Write a function `func ChargeAll(chargers []Charger, amount Money) []error` that attempts `Charge(amount)` on every element and collects any errors (nil entries excluded from the result).

Why: this is the polymorphism payoff — `ChargeAll` never mentions `CreditCard`, `WalletAccount`, or `AuditedWalletAccount` by name; it only depends on the `Charger` interface.

7. Add tests:

- Table-driven test for `CreditCard.Charge` (under limit, exactly at limit, over limit).
- Table-driven test for `WalletAccount.Charge` (sufficient balance, insufficient balance, exact balance).
- A test proving `AuditedWalletAccount.History` grows only on successful charges, not failed ones.
- A test calling `ChargeAll` with a mixed slice of `*CreditCard`, `*WalletAccount`, and `*AuditedWalletAccount`.

### 4. Why this works

This exercise deliberately mixes a value-receiver type (`Money`, copied like a primitive) with pointer-receiver types (`CreditCard`, `WalletAccount`, mutated in place), then ties them together with an interface that doesn't care which receiver style the concrete type uses — only that the method set matches. Embedding is introduced as composition-with-promotion, not inheritance: `AuditedWalletAccount` *has a* `WalletAccount`, and Go promotes its members, but there is no dynamic dispatch back down from parent to child the way base classes work in OOP languages.

### 5. Failure and edge-case drills

- Drill A: Change `CreditCard.Charge` to a value receiver (`func (c CreditCard) Charge(...)`) and call it through a `*CreditCard` stored in a `[]Charger`. Observe whether it still compiles and whether the balance mutation is visible to the caller afterward — explain why.
- Drill B: Try to put a `WalletAccount` (not `*WalletAccount`) value directly into a `[]Charger` when `Charge` is defined with a pointer receiver. Read the exact compiler error and explain it in terms of method sets: a pointer receiver method is not in the method set of the value type.
- Drill C: Call `ChargeAll` with an amount that makes every single charger fail. Confirm the returned `[]error` has the expected length and that no panic occurs from a nil error being dereferenced.
- Drill D: Remove the `Charge` override on `AuditedWalletAccount` and call `.Charge` on an `AuditedWalletAccount` value. Confirm it still compiles (the promoted method resolves), then explain why `History` no longer gets updated.

### 6. Questions to verify understanding

1. What determines whether a type satisfies an interface in Go, and where is that checked — compile time or runtime?
2. Why does `Money.Add` use a value receiver while `CreditCard.Charge` uses a pointer receiver? What rule of thumb decides this in general?
3. What does "method set" mean, and why can a `T` value not be used where a `Charger` requires a method defined with a `*T` receiver?
4. What actually happens when `AuditedWalletAccount` defines its own `Charge` method — is the embedded `WalletAccount.Charge` overridden, hidden, or something else? Be precise.

### 7. Practical challenge (with acceptance criteria)

Challenge: Add a `Refund(amount Money) error` method to the `Charger` interface (rename it or add a second interface `Refunder` — document your choice), implement it for `CreditCard` and `WalletAccount` (undoing a charge, capped so balance never goes negative or over any relevant limit), and write a `Statement(c Charger) string` function that type-switches on the concrete type to print a human-readable one-line summary for `*CreditCard`, `*WalletAccount`, and a default case for anything else.

Acceptance criteria:

- Interface design choice (extend `Charger` vs new `Refunder`) is explicitly justified in a comment.
- `Refund` never drives a balance negative or a credit card balance below zero.
- `Statement` uses a `switch v := c.(type)` type switch, with a `default` branch that does not panic on an unrecognized type.
- Tests cover: successful refund, refund larger than current balance (clamped, not negative), and `Statement` output for all three concrete types plus the default case.
- `go test ./...` passes and `gofmt -w .` produces no further changes.

### 8. Best-practice answer key

```go
package billing

import "fmt"

type Money struct {
	Cents int64
}

func (m Money) Add(other Money) Money {
	return Money{Cents: m.Cents + other.Cents}
}

// Refund is kept on the same Charger interface rather than split into a
// separate Refunder: every charger in this system is always refundable,
// so splitting would only add an interface without any type needing a
// narrower view.
type Charger interface {
	Charge(amount Money) error
	Refund(amount Money) error
}

type CreditCard struct {
	Balance Money
	Limit   Money
}

func (c *CreditCard) Charge(amount Money) error {
	if c.Balance.Cents+amount.Cents > c.Limit.Cents {
		return fmt.Errorf("charge of %d exceeds limit", amount.Cents)
	}
	c.Balance = c.Balance.Add(amount)
	return nil
}

func (c *CreditCard) Refund(amount Money) error {
	c.Balance.Cents -= amount.Cents
	if c.Balance.Cents < 0 {
		c.Balance.Cents = 0
	}
	return nil
}

type WalletAccount struct {
	Balance Money
}

func (w *WalletAccount) Charge(amount Money) error {
	if amount.Cents > w.Balance.Cents {
		return fmt.Errorf("insufficient balance: have %d, need %d", w.Balance.Cents, amount.Cents)
	}
	w.Balance.Cents -= amount.Cents
	return nil
}

func (w *WalletAccount) Refund(amount Money) error {
	w.Balance = w.Balance.Add(amount)
	return nil
}

type AuditedWalletAccount struct {
	WalletAccount
	History []Money
}

func (aw *AuditedWalletAccount) Charge(amount Money) error {
	if err := aw.WalletAccount.Charge(amount); err != nil {
		return err
	}
	aw.History = append(aw.History, amount)
	return nil
}

func ChargeAll(chargers []Charger, amount Money) []error {
	var errs []error
	for _, c := range chargers {
		if err := c.Charge(amount); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func Statement(c Charger) string {
	switch v := c.(type) {
	case *CreditCard:
		return fmt.Sprintf("credit card: balance %d / limit %d", v.Balance.Cents, v.Limit.Cents)
	case *WalletAccount:
		return fmt.Sprintf("wallet: balance %d", v.Balance.Cents)
	default:
		return "unknown charger type"
	}
}
```

Best practices used:

- Interface choice is documented inline with a comment explaining why it wasn't split further.
- Pointer receivers used consistently for every type that mutates its own state.
- `Refund` clamps rather than allowing negative balances, matching the same defensive-bounds pattern from Pack 01's `Bump`.
- Type switch has an explicit `default` so `Statement` never panics on an interface value it doesn't recognize.

### 9. Key takeaway

Interfaces in Go are satisfied structurally by method sets, not declared with an `implements` keyword — which means receiver choice (value vs pointer) is not a style preference, it decides whether a type even satisfies the interface at all call sites.

---

## Also see

- Self-check: `exercises/02-final-questions.md` — work through it only after finishing both exercises above, without re-reading this file first.
