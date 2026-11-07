## Finding

Documented commands must work with the current repository. Two README commands violated this invariant.

## Evidence

`--verbose` fails with `not a number: --verbose`; main.go parses every argument as a number. `make test` fails because no target exists. CI uses `go test ./...`.

## Selected Candidate

Stale README commands: directly reproducible, low risk, and completely fixable in one file.

## Rejected Candidates

Unused median helper: lower value than broken user instructions. Max empty-input handling: no established contract. Sum loop style: cosmetic.

## Change

Removed the verbose example and corrected the test command. Searched all tracked files for other instances; none remain. Changes are uncommitted.

## Validation

With a writable temporary GOCACHE, go vet ./... and go test ./... passed. The retained example printed sum 6 and mean 2. git diff --check passed.

## Risk

Minimal; documentation only.

## Rollback

Restore the two changed README lines.

## Follow-up

None
