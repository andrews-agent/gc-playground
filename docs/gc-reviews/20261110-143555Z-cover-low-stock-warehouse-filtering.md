## Finding

Existing LowStock tests cannot detect warehouse lookups performed before quantity filtering.

## Evidence

Moving warehouse ID collection before filtering passes the old test but fails both new cases. The mutation was reverted.

## Selected Candidate

Strengthen LowStock filtering coverage: a concrete invariant already implemented in inventory/report.go.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap listed open PRs. Documentation links and commands are consistent. Removing unused median offered less substantiated value.

## Change

Added one 27-line table-driven test in inventory/inventory_test.go. Changes remain uncommitted.

## Validation

Targeted tests, make check (all tests and go vet), and git diff --check passed. Used a writable Go cache after sandbox restrictions blocked the default cache; make emitted nonfatal cache warnings.

## Risk

Low; test-only change using the existing missing-warehouse fixture convention.

## Rollback

Remove TestLowStockSkipsUnselectedWarehouses.

## Follow-up

None
