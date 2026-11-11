## Finding

The root CLI's tests covered malformed numbers but left numeric range errors untested.

## Evidence

main.go rejects all ParseFloat errors. A temporary mutation accepting range errors passed existing cases but failed both new cases, printing infinity results.

## Selected Candidate

Close this test gap using the existing command-test table; inventory CLI tests already cover overflow.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap open PRs. Removing unused median lacked sufficient evidence under the requested criteria.

## Change

Added two table cases in main_test.go; 12 lines, no product changes.

## Validation

Ran go mod download: no external dependencies. make check passed all tests and go vet with a writable GOCACHE. git diff --check passed. Mutation testing confirmed both cases detect accepted overflow; source restored. Initial checks encountered sandbox cache restrictions; make also emitted nonfatal xcrun cache warnings.

## Risk

Low: test-only change.

## Rollback

Remove the two added test cases.

## Follow-up

None
