## Finding

The test suite did not enforce Max's documented largest-number contract. Max was the only exported stats helper without tests.

## Evidence

In isolated copies, the old suite passed three incorrect Max implementations: constant zero, zero initialization, and reversed comparison. The new suite rejected all three.

## Selected Candidate

Missing Max coverage: a demonstrated test weakness with a small, complete fix.

## Rejected Candidates

Command-name and empty-input documentation overlap earlier PRs. Sum loop changes are cosmetic. Unused median removal offers less concrete validation.

## Change

Added 24 lines to stats/stats_test.go. Changes remain uncommitted.

## Validation

go test -cover ./... and go vet ./... passed using a writable temporary GOCACHE. Stats coverage is 92.3%. gofmt and git diff --check passed. All three mutation checks distinguished the old and new suites.

## Risk

Low; tests only. No empty-input or special floating-point contract was introduced.

## Rollback

Remove the added TestMax function.

## Follow-up

None
