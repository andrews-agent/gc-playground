## Finding

`median` was dead code with no repository references beyond its declaration.

## Evidence

Repository-wide search found no call sites; the source comment identified it as leftover code.

## Selected Candidate

Unused helper removal, because it is a concrete, behavior-preserving cleanup.

## Rejected Candidates

None; other apparent issues matched prior open cleanup reviews or would alter behavior.

## Change

Deleted `median` and its stale comment from `stats/stats.go`.

## Validation

`go mod download` completed with no dependencies. `make check` could not run because the sandbox blocked macOS cache creation. Running `go test ./...` and `go vet ./...` directly with writable `TMPDIR`/`GOCACHE` passed. `git diff --check` passed.

## Risk

None expected; the function was unexported and unused.

## Rollback

Restore the deleted helper in `stats/stats.go`.

## Follow-up

None
