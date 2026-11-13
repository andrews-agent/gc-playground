## Finding

Single-record error tests discarded returned records, allowing bogus records alongside errors to pass.

## Evidence

Both implementations explicitly return zero values on missing records. A temporary mutation returning bogus records passed existing lookup/error-wrapping tests but failed all four new cases.

## Selected Candidate

Close the single-record error-result coverage gap, following the existing batch lookup assertion pattern.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation navigation was consistent. Unused median removal offered weaker value.

## Change

Added 32 lines in inventory/store_test.go covering both single-record lookup methods in empty and populated stores.

## Validation

go mod download completed: no dependencies to download. Targeted tests passed; mutation verification failed as expected, then production code was restored. make check passed all tests and go vet using writable cache paths, despite a nonfatal macOS cache warning. git diff --check and gofmt verification passed.

## Risk

Low; tests only.

## Rollback

Remove TestMissingSingleLookupResults.

## Follow-up

None
