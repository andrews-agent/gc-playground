## Finding

Subcommands silently ignored trailing arguments, allowing malformed commands such as `inventory stock typo` to execute successfully.

## Evidence

The documented usage defines exact command forms, while the old implementation only checked minimum arguments and ignored extras. New subprocess tests distinguish the old behavior and fail before the fix.

## Selected Candidate

Exact argument validation is a repeated parser invariant with one small, observable fix.

## Rejected Candidates

Duplicate warehouse IDs in LowStock lacked measurable externally visible impact. Other report/query candidates overlapped prior open cleanup PRs.

## Change

Added argument-count checks in `cmd/inventory/main.go` and regression cases in `cmd/inventory/main_test.go`.

## Validation

`go mod download` completed with no dependencies. `go test ./...` passed. `make check` passed (`go test ./...` and `go vet ./...`); the environment emitted a non-fatal xcrun cache warning.

## Risk

Low; documented invocations are unchanged, while malformed extra-argument invocations now exit with status 2.

## Rollback

Revert the changes to `cmd/inventory/main.go` and `cmd/inventory/main_test.go`.

## Follow-up

None
