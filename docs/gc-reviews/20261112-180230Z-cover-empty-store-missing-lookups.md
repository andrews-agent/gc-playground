## Finding

Existing missing-record tests used populated stores, leaving empty-store error handling unprotected.

## Evidence

A temporary mutation suppressed missing-record errors only in empty stores. The existing suite passed; all three new cases failed. The mutation was removed.

## Selected Candidate

Empty-store lookup contract coverage: explicit error contract, complete coverage of all three lookup methods, and demonstrated test value.

## Rejected Candidates

Stock-report batching and arithmetic tests overlap excluded PRs. Removing unused median lacked sufficient supporting evidence. Documentation and navigation screening found no concrete defect.

## Change

Added 26 lines in inventory/store_test.go. Production code is unchanged.

## Validation

Ran go mod download: no external dependencies. make check passed all tests and go vet with writable Go cache paths after the default cache was denied. macOS emitted an xcrun cache warning, but checks completed. git diff --check and gofmt verification passed.

## Risk

Low; tests only.

## Rollback

Remove TestMissingLookupsInEmptyStore.

## Follow-up

None
