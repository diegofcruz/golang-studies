# TOPICS - golang-studies

This roadmap is progressive: from first principles to production-level Go engineering.

## Exercise Packs

**Status legend:** `[x]` done, `[~]` pack created / in progress, `[ ]` not started (no pack yet). Generate new packs on demand with `/new-topic <topic-number-or-slug>`.

- [x] Pack 01 — Setup and Language Fundamentals — [`01-setup-and-fundamentals.md`](01-setup-and-fundamentals.md), self-check: [`01-final-questions.md`](01-final-questions.md)
- [~] Pack 02 — Core Data Structures and Types — [`02-core-data-structures.md`](02-core-data-structures.md), self-check: [`02-final-questions.md`](02-final-questions.md)

## 0) Setup and Tooling Foundations

- Install Go toolchain and verify environment (go version, go env)
- Configure editor, formatter, and linting workflow
- Understand modules, GOPATH history, and modern module-based workflow
- Use gofmt, go test, go vet, and staticcheck basics

## 1) Language Fundamentals

- Variables, constants, zero values, and type system basics
- Control flow: if, for, switch, defer
- Functions: multiple returns, named returns, variadic, closures
- Pointers and value semantics mental model

## 2) Core Data Structures and Types

- Arrays, slices, and slice internals (len, cap, backing array)
- Maps: behavior, nil maps, iteration caveats
- Structs, methods, method sets, and embedding
- Interfaces: implicit implementation and polymorphism

## 3) Error Handling and API Design

- Error-as-value philosophy and explicit error flow
- Sentinel errors, custom errors, wrapping and unwrapping
- errors.Is and errors.As patterns
- Designing small, stable package APIs

## 4) Packages, Modules, and Project Layout

- Module boundaries and semantic import versioning
- Internal packages, visibility, and encapsulation
- Package naming conventions and cohesion
- Organizing small apps vs growing services

## 5) Standard Library Deep Dive

- strings, bytes, strconv, time, math, sort
- io and os packages for file and stream processing
- encoding/json and common serialization pitfalls
- net/http client and server essentials

## 6) Testing Strategy

- Unit tests and table-driven tests
- Subtests and test helpers
- Mocks, fakes, and interface-driven testability
- Golden tests and deterministic test design
- Race detector and concurrency-focused testing

## 7) Concurrency Mental Models

- Goroutines lifecycle and scheduler basics
- Channels: unbuffered, buffered, direction, ownership
- Select patterns and timeout handling
- Fan-in, fan-out, worker pools, pipelines
- sync primitives: Mutex, RWMutex, WaitGroup, Once, Cond

## 8) Context and Cancellation

- context.Context propagation patterns
- Deadlines, cancellation, and cleanup discipline
- Avoiding context misuse and value abuse
- Integrating context into I/O and server workflows

## 9) Memory, Performance, and Runtime

- Escape analysis basics and stack vs heap implications
- Allocation patterns and reducing garbage pressure
- Profiling CPU, memory, goroutines, and blocking
- Benchmarking with go test -bench and interpretation

## 10) Networking and HTTP Services

- Building APIs with net/http
- Middleware patterns and request lifecycle
- Robust clients: retries, backoff, timeouts, idempotency
- JSON API contracts, validation, and compatibility

## 11) Data and Persistence

- database/sql fundamentals and connection pooling
- Transaction patterns and isolation awareness
- SQL safety and query parameterization
- Repository patterns and trade-offs

## 12) Observability and Reliability

- Structured logging patterns
- Metrics fundamentals and SLI/SLO mindset
- Tracing basics and request correlation
- Health checks, readiness, and graceful shutdown

## 13) Configuration and Environment Management

- 12-factor style configuration practices
- Env var parsing and validation
- Secrets handling and operational safety
- Feature flags and runtime behavior control

## 14) Security Essentials

- Input validation and output encoding
- Authentication and authorization fundamentals
- Secure defaults in HTTP servers and clients
- Dependency and supply-chain awareness

## 15) CLI and Automation

- Building CLIs with flag and command patterns
- File/process automation and robust scripting
- UX considerations for command-line tools
- Packaging and distribution basics

## 16) Advanced Language Features and Patterns

- Generics fundamentals and real-world use cases
- Reflection trade-offs and safe usage boundaries
- Code generation workflows and maintenance costs
- Functional options and composition patterns

## 17) Architecture and Production Patterns

- Layered design and dependency direction
- Domain boundaries and service decomposition
- Idempotent workflows and distributed failure handling
- Background jobs and asynchronous processing

## 18) Deployment and Operations

- Build reproducibility and release versioning
- Containerization and runtime tuning basics
- Deployment strategies and rollback practices
- Post-deploy verification and incident response

## 19) Capstone Projects

- Concurrent data pipeline with backpressure
- Production-style REST API with auth, persistence, and observability
- Resilient worker system with retries and dead-letter handling
- CLI plus service integration project

## Exercise Policy for Every Topic

Each topic should include:

- Happy path implementation exercise
- Failure-path exercise (invalid input, timeout, partial failure)
- Edge-case drill (boundary values, nil handling, race scenarios)
- Refactoring challenge focused on readability and maintainability
- Performance or reliability challenge where applicable

## Completion Milestones

- M1: Fundamentals and core types complete
- M2: Errors, modules, standard library, and testing complete
- M3: Concurrency, context, and runtime performance complete
- M4: HTTP, persistence, and observability complete
- M5: Security, architecture, and capstone complete
