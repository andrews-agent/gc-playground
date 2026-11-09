## Finding

Documented CLI defaults must match argument handling. The inventory guide incorrectly states 10.

## Evidence

docs/inventory.md stated “10 by default”; cmd/inventory/main.go:31 initializes threshold to 5. Repository-wide searches found no other incorrect default claims.

## Selected Candidate

Stale threshold documentation: directly supported by code and completely resolved with one value correction.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Sum loop styling is cosmetic; unused median removal offers weaker evidence. Stock-report assertions would touch coverage included in the batching PR.

## Change

Changed 10 to 5 in docs/inventory.md. Left changes uncommitted; created no report file.

## Validation

make check passed go test ./... and go vet ./... with a writable GOCACHE, despite a nonfatal macOS cache warning. git diff --check passed. Reviewed the complete one-line diff.

## Risk

Minimal; documentation only.

## Rollback

Restore the changed value in docs/inventory.md.

## Follow-up

None
