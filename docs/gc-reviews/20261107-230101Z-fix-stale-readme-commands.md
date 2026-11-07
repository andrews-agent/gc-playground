## Finding

README commands must work against the current repository. Two documented commands violated that invariant.

## Evidence

The CLI rejects --verbose with 'not a number: --verbose'. make test fails because no test target exists. CI runs go test ./....

## Selected Candidate

Stale README commands: directly reproducible, fully screened across the repository, and fixable without changing product behavior.

## Rejected Candidates

Unused median helper removal is a separate cleanup class. Max empty-input behavior has no established contract. Loop style changes would be cosmetic.

## Change

README.md only: removed the invalid example and corrected the test command. Left uncommitted.

## Validation

go test ./... and go vet ./... passed using a writable GOCACHE. The retained example printed sum 6 and mean 2. git diff --check passed.

## Risk

Minimal; documentation only.

## Rollback

Restore the two changed README lines.

## Follow-up

None
