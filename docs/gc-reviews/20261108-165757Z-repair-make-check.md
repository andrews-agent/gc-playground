## Finding

The documented `make check` target must run available repository checks, but it invoked nonexistent scripts/lint.sh.

## Evidence

Screened all source, tests, documentation, and tooling. CI runs go test and go vet. Before the change, make check passed tests then failed with 'No such file or directory' for the lint script.

## Selected Candidate

Broken local check target: directly reproducible, isolated, and supported by CI's established convention.

## Rejected Candidates

Inventory threshold tests and error wrapping require separate fixes. Previously reported documentation and test issues were excluded. Arithmetic style and unused helper cleanup offered weaker value.

## Change

Changed one Makefile recipe line to go vet ./.... Searched for other missing lint-script references; found none. Changes remain uncommitted.

## Validation

make check passed with TMPDIR and GOCACHE redirected to writable workspace locations; tests and go vet passed. git diff --check passed. macOS emitted a nonfatal cache permission warning.

## Risk

Low; only the local check recipe changes.

## Rollback

Restore the previous Makefile recipe line.

## Follow-up

None
