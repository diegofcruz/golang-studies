# New Topic — Generate Exercise Pack

Generate a complete exercise pack for the topic: **$ARGUMENTS**

The argument is a topic number or slug from `exercises/TOPICS.md` (e.g. `3`, `03`, or `error-handling`).

---

## Workflow

### 1. Resolve the topic

Match `$ARGUMENTS` against the numbered sections in `exercises/TOPICS.md` (e.g. `## 3) Error Handling and API Design`). If it's ambiguous or not found, list the closest matching section headers and ask before proceeding.

### 2. Read context

Before generating anything, read:

- `CLAUDE.md` — exercise format rules, code style, exercise design rules
- `exercises/TOPICS.md` — to see prerequisites and what's already covered
- Every existing `exercises/NN-*.md` pack — to match tone, depth, and format exactly

### 3. Determine prerequisites

From `TOPICS.md` and existing packs, identify:
- Which earlier topics already have a pack (so you know what the student can be assumed to know)
- Which Go concepts from earlier packs this topic builds on — name them explicitly in the new pack's "Prerequisites" sections

### 4. Design the exercise sequence

Exactly 2 exercises (per `CLAUDE.md` Exercise Design Rules), unless the topic is unusually broad (concurrency, standard library deep dive, HTTP services) — then up to 3. Each exercise:

- Trains one coherent slice of the topic, not a single isolated fact
- Is strictly harder than the previous exercise in the same pack
- Follows the exact 9-part format from `CLAUDE.md` → "Default Exercise Format"
- Includes at least one failure/edge-case drill
- Ends with a practical challenge with concrete acceptance criteria, then a full runnable answer key, then a one-sentence key takeaway

### 5. Write the exercise pack

Output the complete markdown file to `exercises/NN-topic-slug.md`, where `NN` is the zero-padded topic number from `TOPICS.md` and `topic-slug` is a short kebab-case name.

File header:

```md
# Exercise Pack NN - <Topic Title>

This pack covers Topic <N> from TOPICS.md.

## Exercise 1: <Title>
...
## Exercise 2: <Title>
...
```

### 6. Write the final-questions file

Create `exercises/NN-final-questions.md` following `CLAUDE.md` → "Final Questions / Self-Check": Conceptual Questions, Code Prediction Questions, Implementation Challenges. Leave every question unanswered — this file is for the student to fill in themselves after finishing the pack.

### 7. Do not write solution code files

Per `CLAUDE.md` Exercise Design Rule 7, do not create the actual `general/<package>/` files. The pack's step-by-step instructions and answer key are enough for the student to build it themselves — that act of building it is the exercise.

### 8. Update TOPICS.md

Update the "Exercise Packs" list at the top of `exercises/TOPICS.md` to add a line for this pack, following the existing format (see Pack 01/02 entries). Do not otherwise rewrite the file's topic sections.

---

## Quality Checklist Before Writing

- [ ] Every exercise maps to the exact 9-part structure — no missing sections
- [ ] At least one failure-path drill exists per exercise
- [ ] The answer key is complete, idiomatic, compilable Go — no stubs, no TODOs
- [ ] Every "Key takeaway" is one sentence capturing the whole mental-model shift
- [ ] No exercise silently assumes a concept from a topic that has no pack yet
- [ ] No external dependencies unless the topic is specifically about a library

---

## Output

After creating the files, summarize:

1. Path to the exercise pack created
2. Path to the final-questions file created
3. Number of exercises generated and their titles
4. Which prior packs/topics this one assumes are already done
