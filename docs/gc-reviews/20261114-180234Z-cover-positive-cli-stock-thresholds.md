## Finding

The CLI must honor an explicit threshold, but its tests covered no positive argument.

## Evidence

Temporarily forcing positive thresholds to 5 caused both new cases to fail. The production code was restored afterward.

## Selected Candidate

Complete positive-threshold CLI coverage using the existing command test table.

## Rejected Candidates

Stock-report batching repeats an excluded open PR. Arithmetic coverage overlaps another excluded PR. Screening the remaining code, docs, and tests found no stronger safe change.

## Change

Added two cases in cmd/inventory/main_test.go, totaling 14 lines. No product behavior changed.

## Validation

Ran go mod download: no dependencies to download. Targeted CLI tests, make check (all tests and go vet), and git diff --check passed. Used a writable Go cache; make emitted an xcrun cache permission warning but completed successfully.

## Risk

Low; test-only change.

## Rollback

Remove the two added test cases.

## Follow-up

None
