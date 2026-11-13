## Finding

The nil-ID and empty-ID test cases expected initialized empty maps but also accepted nil results.

## Evidence

A temporary mutation returning (nil, nil) for empty IDs passed the original assertions. The strengthened test failed for both affected cases.

## Selected Candidate

Close the demonstrated warehouse-result assertion gap with one shared assertion.

## Rejected Candidates

Stock-report batching, lookup failure attribution, and arithmetic coverage overlap open PRs. Other source and documentation candidates lacked sufficient evidence for a safe change.

## Change

Added a three-line non-nil assertion in inventory/store_test.go. Reviewed all maps.Equal uses; no other empty-map expectations needed this fix.

## Validation

Ran go mod download: no external dependencies. Targeted test passed; mutation proof distinguished old and new assertions. make check passed all tests and go vet using workspace-local caches after default cache access was denied. git diff --check passed. Temporary mutations and caches were removed.

## Risk

Low; test-only change with no product behavior changes.

## Rollback

Remove the added non-nil assertion.

## Follow-up

None
