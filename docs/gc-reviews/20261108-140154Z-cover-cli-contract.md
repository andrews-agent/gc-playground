## Finding

Existing tests exercised stats helpers but left main.go's parsing, output, and error-exit contracts unprotected.

## Evidence

Changing the sum label caused three new cases to fail. Changing the invalid-input exit status to zero caused the error case to fail. Both mutations were restored.

## Selected Candidate

Missing CLI contract coverage: directly grounded in main.go and distinct from the open arithmetic-test PRs.

## Rejected Candidates

Arithmetic tests and package documentation overlap open PRs. Sum loop style is cosmetic. Max empty-input behavior has no established contract. Unused median removal offers less value.

## Change

Added 66 lines in main_test.go covering all CLI branches. Production code is unchanged; changes remain uncommitted.

## Validation

go test ./... and go vet ./... passed with a writable GOCACHE. Both mutation checks failed as expected. git diff --check passed.

## Risk

Low. Tests add one command build per suite run and require the Go executable.

## Rollback

Remove main_test.go.

## Follow-up

None
