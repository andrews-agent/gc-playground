## Finding

The inventory guide mentions a default threshold but omits how to override it.

## Evidence

cmd/inventory/main.go accepts `low [threshold]`; existing CLI tests cover explicit thresholds. The guide previously showed only `low`.

## Selected Candidate

Incomplete command documentation: directly established by implementation and existing tests, with a small documentation-only fix.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Loop styling is cosmetic; removing unused median has weaker justification. Navigation links and CLI tests already cover their established contracts.

## Change

Updated docs/inventory.md with the optional argument and a zero-stock example. Searched other usage references; they already include it. Changes remain uncommitted.

## Validation

make check passed all tests and go vet using a writable GOCACHE and TMPDIR; make emitted a nonfatal cache warning. The documented example printed GEAR-L with quantity 0. git diff --check passed.

## Risk

Minimal; documentation only.

## Rollback

Revert the changes to docs/inventory.md.

## Follow-up

None
