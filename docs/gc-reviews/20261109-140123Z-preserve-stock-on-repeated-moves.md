## Finding

Store.Move overwrote existing destination stock, violating quantity conservation.

## Evidence

The regression test failed before the fix: transfers of 20 and 15 left 15 instead of 35 at the destination; a subsequent zero transfer erased it.

## Selected Candidate

Destination quantity overwrite: a concrete stock-loss bug with a deterministic regression test. Repository-wide write-path search found one affected implementation.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap listed open PRs. Unused median removal had less practical value. Documentation and navigation screening revealed no stronger candidate.

## Change

One production-line fix plus a 27-line regression test in inventory/store.go and inventory/inventory_test.go. Changes remain uncommitted.

## Validation

Regression test failed before the fix. make check passed with writable TMPDIR and GOCACHE overrides, running all tests and go vet. Formatting and git diff --check passed. make emitted a nonfatal sandbox cache warning.

## Risk

Low; repeated transfers now preserve existing destination quantities.

## Rollback

Revert the two-file diff.

## Follow-up

None
