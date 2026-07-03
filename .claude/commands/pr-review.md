# PR Review — golang-studies

Review the current branch against the base branch: **$ARGUMENTS** (default: `main` when no argument is provided).

Generate a concise review report in **Portuguese**, prioritizing findings by severity (`alta`, `media`, `baixa`) across three lenses:

1. **Learning absorption** — primary lens
2. **Project patterns**
3. **Go best practices**

---

## Step-by-step Workflow

### 1. Resolve base branch

If no argument was provided, use `main` as the base branch.

### 2. Validate git context

Run these in order:

```bash
git rev-parse --is-inside-work-tree
git rev-parse --abbrev-ref HEAD
```

Verify the base branch exists:

```bash
git show-ref --verify --quiet refs/heads/<BASE_BRANCH> || git show-ref --verify --quiet refs/remotes/origin/<BASE_BRANCH>
```

### 3. Collect the diff

```bash
git diff --stat <BASE_BRANCH>...HEAD
git diff -U3 <BASE_BRANCH>...HEAD
```

If diff is empty, stop and return:

> `Sem achados: não há diferenças entre <BASE_BRANCH> e a branch atual.`

### 4. Run the Go toolchain over changed packages

The workspace root has no `go.mod` of its own — only `general/` and `leetcodes/` do — so `./...` does not resolve from the root by itself. From the repo root, run:

```bash
gofmt -l .
go vet ./general/... ./leetcodes/...
go test ./general/... ./leetcodes/...
```

Treat any `gofmt -l` output, `go vet` warning, or failing test as a finding in Lens 3, not just a passive note.

### 5. Run all three review lenses on the same diff

Evaluate all lenses in a single pass. Do not call sub-agents or external tools.

---

## Lens 1 — Learning Absorption

Evaluate whether changed exercise content (`exercises/*.md`) helps knowledge absorption:

- Clear progression, one concept slice at a time within each exercise
- Explicit `why` behind non-obvious steps, traced through compiler/runtime mechanics
- Failure paths, edge cases, and incorrect-usage drills present
- Full runnable answer key — no stubs, no TODOs
- A one-sentence "Key takeaway" present and specific, not generic

Flag these as `alta`:
- Missing failure-path drills for an exercise
- Answer key with stubs or TODO instead of working code
- Key takeaway absent or vague
- An exercise mixes more than one topic's worth of new concepts

Flag these as `media`:
- Weak or missing "Why" explanations on non-trivial steps
- Acceptance criteria for the practical challenge are vague or unchecked
- Prerequisites section missing or inconsistent with `TOPICS.md`

Flag these as `baixa`:
- Minor wording/clarity improvements to goals or steps

---

## Lens 2 — Project Patterns

Evaluate structural consistency with this repository:

- Pack naming: `exercises/NN-topic-slug.md` + `exercises/NN-final-questions.md`
- Code placement: `general/<package>/<package>.go` + `general/<package>/<package>_test.go`, matching the package name
- `go.work` lists every module that has a `go.mod` (no orphaned modules)
- `exercises/TOPICS.md` "Exercise Packs" list stays in sync with files actually present

Flag these as `alta`:
- Go code placed outside its module's expected package folder
- A new `go.mod` created but not added to `go.work`
- Exercise pack file placed outside `exercises/` or misnamed

Flag these as `media`:
- Test file naming inconsistent with `<package>_test.go`
- `TOPICS.md` not updated when a new pack was added

Flag these as `baixa`:
- Minor structural deviations that don't affect navigation

---

## Lens 3 — Go Best Practices

Evaluate maintainability, correctness, and idiomatic quality of any `.go` files in the diff:

- `gofmt -l .` output, `go vet ./...` warnings, failing `go test ./...`
- Errors as values: no panic for expected/input-validation failures
- Error wrapping with `fmt.Errorf("...: %w", err)` where context is added
- Table-driven tests with `t.Run` subtests and explicit failure messages
- `context.Context` used for cancellation/deadlines where relevant
- For concurrent code: `go test -race ./...` clean, explicit channel ownership

Flag these as `alta`:
- `gofmt -l` lists a changed file
- `go vet` warning or failing test in a changed package
- Panic used for an expected/validatable failure instead of returning an error
- Concurrent code without a passing `-race` run

Flag these as `media`:
- Missing table-driven structure for non-trivial test logic
- Error returned without context via `%w` where the caller would need it
- Package doing more than one clearly separable job

Flag these as `baixa`:
- Minor naming inconsistencies
- Redundant `fmt.Println` left over from debugging

---

## Consolidate and Output

1. Merge all findings from the three lenses.
2. Remove duplicates (keep the highest severity when a finding appears in multiple lenses).
3. Sort by severity: `alta` → `media` → `baixa`.
4. Output in Portuguese only.

Use this exact structure:

```md
## Findings

### Alta

- [lens] título curto do problema
  Evidência: caminho/arquivo:linha — trecho relevante
  Impacto: por que isso prejudica o aprendizado ou a qualidade
  Recomendação: ação concreta para corrigir

### Media

- ...

### Baixa

- ...
```

If a severity has no findings, include: `- Sem achados.`

---

## Rules

1. Review only what exists in the `<BASE_BRANCH>...HEAD` diff.
2. Do not invent files, behavior, or issues not present in the diff.
3. Prioritize learning-absorption findings when severity is tied.
4. Do not propose auto-fix or commit operations.
5. Keep recommendations actionable and tied to specific evidence.
6. Do not generate markdown outside the output contract above.
