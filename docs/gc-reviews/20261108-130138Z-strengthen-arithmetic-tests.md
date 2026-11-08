## Finding

Both arithmetic tests relied on the same positive input, leaving their input-dependent calculation contracts weakly protected.

## Evidence

In an isolated copy, replacing each Sum contribution with 2 passed the original tests but failed eight new subtests across Sum and Mean.

## Selected Candidate

Strengthen both arithmetic tests using the same fix recipe, supported by a demonstrated coverage gap.

## Rejected Candidates

Max coverage and command documentation overlap excluded PRs. Sum loop style is cosmetic. Removing unused median offered less demonstrable value.

## Change

Updated stats/stats_test.go: 27 insertions and 7 deletions. Production code is unchanged; changes remain uncommitted.

## Validation

go test ./... and go vet ./... passed using a writable GOCACHE. gofmt reported no differences; git diff --check passed. Mutation comparison confirmed stronger detection.

## Risk

Low: tests only, with exactly representable fractional expectations.

## Rollback

Restore stats/stats_test.go.

## Follow-up

None
