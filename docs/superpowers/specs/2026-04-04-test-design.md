# Test Design — mgjules-go

**Date:** 2026-04-04  
**Status:** Approved

## Context

The codebase currently has zero test files. This spec covers a two-phase test implementation that covers all meaningful units of logic using `testify` for assertions and a hand-written mock for the `Repository` interface.

## Decisions

| Decision | Choice | Rationale |
|---|---|---|
| Scope | All 8 areas, two phases | Design once for consistency; implement incrementally for early value |
| Library | `testify` (`assert`/`require`) | Better readability than stdlib; no codegen tooling needed |
| Mock strategy | Hand-written `MockRepository` in `internal/testutil` | Single definition reused by both `Fetcher` and `Projecter` tests |
| Test style | Mixed: table-driven for pure functions, individual named tests for stateful/concurrent behaviour | Matches Go idiom; named tests are more readable for complex scenarios |
| Coverage target | None enforced | Avoids gaming metrics with trivial tests; quality over quantity |
| Race detection | `-race` flag for concurrent packages | Catches data races in `Fetcher` and `Projecter` that `-count=1` alone won't |

## Shared Infrastructure

### `internal/testutil/mock_repository.go`

A hand-written struct with one `func` field per method on `repository.Repository`. Each field defaults to `nil`; tests set only the fields they need.

```go
type MockRepository struct {
    GetMetaFn          func(ctx context.Context, id string) (*entity.Meta, error)
    GetLinksFn         func(ctx context.Context) ([]entity.Link, error)
    GetIntroductionFn  func(ctx context.Context, id string) (*entity.Introduction, error)
    GetSectionsFn      func(ctx context.Context) ([]entity.Section, error)
    GetExperiencesFn   func(ctx context.Context) ([]entity.Experience, error)
    GetProjectsFn      func(ctx context.Context) ([]entity.Project, error)
    GetContributionsFn func(ctx context.Context) ([]entity.Contribution, error)
    GetAwardsFn        func(ctx context.Context) ([]entity.Award, error)
    GetInterestsFn     func(ctx context.Context) ([]entity.Interest, error)
    GetLanguagesFn     func(ctx context.Context) ([]entity.Language, error)
    GetPostsFn         func(ctx context.Context) ([]entity.Post, error)
    CleanupFn          func() error
}
```

Each method delegates to its `Fn` field. If the field is nil the method returns zero values, making tests that don't care about a particular method require no setup for it.

### `internal/testutil/entity_builders.go`

Helper functions returning minimal valid entity structs. Tests set only fields relevant to their assertion, keeping setup concise.

```go
func BuildMeta() entity.Meta          { ... }
func BuildExperience() entity.Experience { ... }
// one per entity type
```

### Conventions

- All test files use **external test packages** (`package foo_test`) — tests the public API only.
- Use `require` (fatal) for setup/preconditions; `assert` (non-fatal) for outcome assertions so all failures are reported.
- Table-driven tests for pure functions with multiple input/output variants.
- Individual named tests (`TestFetcher_Fetch_NotifiesSubscribers`) for stateful or concurrent behaviour.

---

## Phase 1 — Pure unit tests (no mocks or external dependencies)

### `internal/auth` → `auth_test.go`

Style: table-driven + one panic test.

| Test | Input | Expected |
|---|---|---|
| correct token | matching string | `true` |
| wrong token | non-matching string | `false` |
| empty input | `""` | `false` |
| `New("")` | empty token | panics |

### `internal/fetcher` (DeepCopy) → `copier_test.go`

Style: table-driven over several entity types.

Each case:
1. Creates an original entity.
2. Deep-copies it.
3. Mutates a field on the copy.
4. Asserts the original is unchanged.

Covers pointer fields, slice fields, and nested structs to confirm independence.

### `internal/config` → `config_test.go`

Style: table-driven using `t.Setenv`.

| Test | Env state | Expected |
|---|---|---|
| missing `AUTH_TOKEN` | unset | returns error |
| all required vars set | valid values | correct struct fields |
| defaults applied | `SERVER_HOST`/`SERVER_PORT` unset | `0.0.0.0` / `13337` |

### `internal/http` (render helpers) → `render_test.go`

Style: table-driven using `httptest.NewRecorder()`.

`renderJSON`:
- `nil` input → 204, no body
- valid struct → 200, `Content-Type: application/json`, correct JSON body

`renderHTML`:
- empty byte slice → 204, no body
- non-empty bytes → 200, `Content-Type: text/html`, `Content-Encoding: gzip`

### `internal/repository/static` → `static_test.go`

Style: one test per `GetFoo` method, each asserting structural integrity of returned data.

Assertions per entity collection:
- Length > 0
- No zero-value IDs
- No empty required string fields (Company, Position, Title, etc.)
- Valid `From` dates (non-zero `time.Time`)
- Optional pointer fields (`To *time.Time`) are nil or point to a date after `From`

Purpose: guard against silent regressions in the hardcoded content.

---

## Phase 2 — Mock-dependent tests

### `internal/fetcher` → `fetcher_test.go`

Uses `testutil.MockRepository`. Style: individual named tests.

| Test | Description |
|---|---|
| `TestFetcher_Fetch_PopulatesData` | Mock returns valid entities; all accessors return expected data after `Fetch()` |
| `TestFetcher_Fetch_PartialError` | One `GetFn` returns error; successful fields still populated |
| `TestFetcher_Fetch_NotifiesSubscribers` | Subscriber registered via `AddSubscriber` is called after successful `Fetch()` |
| `TestFetcher_Fetch_RaceCondition` | `Fetch()` and all accessors run concurrently; run with `-race` |
| `TestFetcher_FetchedAt` | Zero before first fetch, non-zero after |

### `internal/projecter` → `projecter_test.go`

Uses `testutil.MockRepository` (via a real `Fetcher` seeded with the mock). Style: individual named tests.

| Test | Description |
|---|---|
| `TestProjecter_Get_ExistingKey` | After `Build()`, `Get("index")` returns non-empty bytes and `found=true` |
| `TestProjecter_Get_EmptyKey` | `Get("")` returns `nil, false` — no projection is built for the empty key |
| `TestProjecter_Get_MissingKey` | `Get("nonexistent")` returns `nil, false` |
| `TestProjecter_Get_RebuildInDevMode` | With `prod=false`, `Get()` rebuilds on every call (verified via mock call counter) |

### `internal/http` (handlers) → `handler_test.go`

Requires a `Projecter` interface extracted in `internal/http` so a stub can be injected:

```go
type projecterInterface interface {
    Get(keys ...string) ([]byte, bool)
}
```

Style: individual named tests using `httptest.NewRecorder()`.

| Test | Description |
|---|---|
| `TestRootHandler_FoundInProjecter` | Path resolved from projections → 200, gzip HTML |
| `TestRootHandler_FallbackToIndexKey` | Bare path falls back to `path/index` key |
| `TestRootHandler_StaticFile` | Path not in projections but exists in static FS → file served |
| `TestRootHandler_NotFound` | Path in neither projections nor static FS → 404 HTML projection served |
| `TestNotFoundHandler_MissingProjection` | 404 projection not built → 500 JSON error |

---

## Out of Scope

- **Directus / EdgeDB repositories** — require live external services; candidates for future integration tests.
- **Template output correctness** — `*_templ.go` files are generated; their rendering is covered indirectly via `Projecter` tests.
- **Load / benchmark tests** — not planned.

---

## Running Tests

```bash
# All tests
go test ./...

# With race detector (required for Fetcher and Projecter)
go test -race ./...

# Single package
go test ./internal/auth/...
```
