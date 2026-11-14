## Finding

Two independent test loops used the parent testing.T, so t.Fatal skipped subsequent cases.

## Evidence

Temporary fault-injection overlays showed the original loops stopping after their first cases. After the fix, all six remaining cases ran and passed. Other independent test tables already use t.Run.

## Selected Candidate

Fix both remaining independent case loops with fatal assertions, following existing subtest conventions.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Removing unused median lacked sufficient contract evidence. The accumulated-move loop intentionally shares sequential state.

## Change

Wrapped cases in TestLowStock and TestAddWarehouseReplacesExistingID in named subtests. Assertions and product code are unchanged.

## Validation

go mod download completed: no dependencies to download. Eight targeted subtests passed. make check passed, including go test ./... and go vet ./..., using a writable Go cache. git diff --check passed. macOS make emitted a nonfatal cache warning; no checks remained blocked.

## Risk

Low: test organization only.

## Rollback

Revert the changes to the two inventory test files.

## Follow-up

None
