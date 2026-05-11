# Exercise Pack 01 - Setup and Language Fundamentals

This pack covers Topic 0 and Topic 1 from TOPICS.md.

## Exercise 1: Environment and Tooling Baseline

### 1. Goal

Set up a reliable Go development environment and confirm you can compile, format, vet, and test a minimal module.

### 2. Prerequisites

- macOS terminal access
- VS Code installed
- Internet access to install Go and tooling

### 3. Step-by-step implementation

1. Install Go from the official distribution.
   Why: You need the compiler, standard library, and go command to build anything in Go.

2. Verify installation:

```bash
go version
go env GOROOT GOPATH GOMODCACHE
```

Why: This confirms your runtime and module cache paths so you can reason about where dependencies and toolchain files live.

3. In this repository, initialize your first module:

```bash
go mod init golang-studies
```

Why: Modules are the dependency and build boundary in modern Go. Everything else in this course depends on module-aware behavior.

4. Create main.go with a tiny program:

```go
package main

import "fmt"

func main() {
	fmt.Println("go studies ready")
}
```

Why: A minimal executable validates compiler, package layout, and entrypoint behavior.

5. Format and run quality checks:

```bash
gofmt -w .
go vet ./...
go test ./...
```

Why: gofmt enforces idiomatic style automatically, go vet catches suspicious constructs, and go test validates package-level correctness.

6. Run the program:

```bash
go run .
```

Expected output:

```text
go studies ready
```

Why: Final verification that module resolution and compilation pipeline work end-to-end.

### 4. Why this works

The go command unifies build, test, module, and package operations. By starting with module init plus format/vet/test/run, you exercise the complete baseline workflow every professional Go project depends on.

### 5. Failure and edge-case drills

- Drill A: Remove go.mod and run go test ./... again. Observe module resolution failure, then restore with go mod init.
- Drill B: Introduce an unused variable and run go test ./.... Observe compile-time strictness.
- Drill C: Write badly formatted code and run gofmt -w . to confirm deterministic formatting.

### 6. Questions to verify understanding

1. Why does go mod init change how imports and dependencies are resolved?
2. What problem does gofmt solve that manual formatting does not?
3. What is the practical difference between go vet and go test?

### 7. Practical challenge (with acceptance criteria)

Challenge: Add a package named bootcheck with a function EnvironmentOK() bool and test it.

Acceptance criteria:

- bootcheck package exists with at least one exported function.
- A table-driven test file validates true and false scenarios.
- go test ./... passes.
- gofmt -w . produces no further changes.

### 8. Best-practice answer key

- Keep logic in small packages instead of putting everything in main.
- Use table-driven tests even for simple logic to build repeatable testing habits.
- Handle setup assumptions explicitly in tests instead of relying on machine state.

Example:

```go
package bootcheck

func EnvironmentOK(version string) bool {
	return version != ""
}
```

```go
package bootcheck

import "testing"

func TestEnvironmentOK(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{name: "valid", version: "go1.22.0", want: true},
		{name: "empty", version: "", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := EnvironmentOK(tc.version)
			if got != tc.want {
				t.Fatalf("EnvironmentOK(%q) = %v, want %v", tc.version, got, tc.want)
			}
		})
	}
}
```

Best practices used:

- Small focused function
- Table-driven testing
- Explicit assertion messages

### 9. Key takeaway

The non-negotiable Go baseline is module-aware development plus automated formatting, vetting, and testing on every change.

---

## Exercise 2: Language Fundamentals Through a Mini CLI

### 1. Goal

Practice variables, constants, control flow, functions, and pointers by building a small score classifier CLI.

### 2. Prerequisites

- Exercise 1 complete
- Comfortable running go run . and go test ./...

### 3. Step-by-step implementation

1. Create package score with:

- const PassMark = 70
- func Classify(score int) (string, error)
- func Bump(score \*int, delta int)

Why: This single package touches constants, function returns, validation, and pointer-based mutation in a controlled scope.

2. Implement Classify rules:

- score < 0 or > 100 returns error
- score >= 90 => "excellent"
- score >= PassMark => "pass"
- otherwise => "fail"

Why: Ordered branching is a practical way to encode domain rules while enforcing input constraints early.

3. Implement Bump:

- Add delta to the pointed score.
- Clamp final value between 0 and 100.

Why: Pointer parameters let a function modify caller state; clamping introduces defensive bounds logic.

4. In main.go, parse one integer from os.Args and print classification.

Why: Wiring package logic into an executable reinforces package boundaries and function composition.

5. Add tests:

- Table-driven tests for Classify
- Edge tests for -1, 0, 70, 90, 100, 101
- Direct test for Bump mutation and clamping

Why: Fundamentals become durable only when you test edge conditions, not just expected input.

### 4. Why this works

This exercise combines core language constructs into one realistic flow: validate input, branch on ranges, return explicit errors, and use pointers intentionally. It builds mental models for value vs reference behavior without introducing concurrency yet.

### 5. Failure and edge-case drills

- Drill A: Pass invalid command-line input and confirm graceful error message.
- Drill B: Remove bounds checks in Classify and observe incorrect classifications.
- Drill C: Pass a nil pointer to Bump and decide whether to guard or panic; document your API choice.

### 6. Questions to verify understanding

1. Why is returning (string, error) preferred over panic for invalid score input?
2. What changes when Bump receives int vs \*int?
3. Why should validation happen before classification branches?
4. What are the benefits of constants like PassMark vs magic numbers?

### 7. Practical challenge (with acceptance criteria)

Challenge: Extend classification to include "distinction" for scores >= 95 and add a Normalize function that rounds any score to nearest multiple of 5 before classification.

Acceptance criteria:

- New category is correctly prioritized in logic.
- Normalize behavior is deterministic and tested with table-driven tests.
- Existing pass/fail behavior remains unchanged for unaffected values.
- go test ./... passes with new and existing cases.

### 8. Best-practice answer key

Guidance:

- Validate first, then normalize, then classify.
- Keep Classify single-purpose and move rounding to Normalize.
- Wrap input errors with context using fmt.Errorf and %w where helpful.

Reference implementation:

```go
package score

import "fmt"

const PassMark = 70

func Normalize(v int) int {
	if v%5 == 0 {
		return v
	}
	rem := v % 5
	if rem >= 3 {
		return v + (5 - rem)
	}
	return v - rem
}

func Classify(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", fmt.Errorf("invalid score: %d", score)
	}

	score = Normalize(score)

	switch {
	case score >= 95:
		return "distinction", nil
	case score >= 90:
		return "excellent", nil
	case score >= PassMark:
		return "pass", nil
	default:
		return "fail", nil
	}
}

func Bump(score *int, delta int) {
	if score == nil {
		return
	}
	*score += delta
	if *score < 0 {
		*score = 0
	}
	if *score > 100 {
		*score = 100
	}
}
```

Best practices used:

- Early validation and explicit error returns
- Clear package-level constant for domain rule
- Small focused functions with single responsibilities
- Defensive nil-pointer handling
- Switch with ordered guard logic for readability

### 9. Key takeaway

Go fundamentals become reliable when you combine simple language features with strict validation, explicit errors, and edge-focused tests.
