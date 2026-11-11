## Finding

Successful AddItem replacements lacked coverage despite existing tests for fresh inserts and rejected replacements.

## Evidence

A mutation that silently skipped existing SKUs passed the original full suite. All three new cases rejected it. The mutation was restored.

## Selected Candidate

Successful replacement coverage: a demonstrated test gap in the existing normalized-SKU storage behavior.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation and navigation were current. Unused-helper removal and loop rewrites lacked sufficient evidence.

## Change

Added one 23-line table-driven test in inventory/inventory_test.go. No production changes.

## Validation

go mod download reported no dependencies. Targeted tests and make check passed, including all tests and go vet. gofmt and git diff --check were clean. Used a writable Go cache; make emitted nonfatal macOS cache warnings. No checks remained blocked.

## Risk

Low; test-only change following existing conventions.

## Rollback

Remove TestAddItemReplacesExistingSKU.

## Follow-up

None
