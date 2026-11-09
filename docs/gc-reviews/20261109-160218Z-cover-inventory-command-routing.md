## Finding

cmd/inventory/main_test.go exercised only low, leaving other command routes and their stdout, stderr, and exit-status contracts unprotected.

## Evidence

Deliberately breaking show routing passed the original tests but failed all three new show cases. The mutation was restored.

## Selected Candidate

Missing inventory CLI contract coverage: concrete, low-risk, and consistent with the root CLI test pattern.

## Rejected Candidates

Stock-report batching, repeated moves, and arithmetic coverage overlap excluded PRs. Sum loop style is cosmetic. Other inventory behavior changes lacked sufficiently established contracts.

## Change

Updated one test file: 44 insertions and 6 deletions. Added six cases and generalized the existing argument table. Changes remain uncommitted.

## Validation

Targeted tests, make check (go test ./... and go vet ./...), and git diff --check passed. make emitted a sandbox cache warning but exited successfully.

## Risk

Low; tests only, using the existing subprocess harness.

## Rollback

Revert cmd/inventory/main_test.go.

## Follow-up

None
