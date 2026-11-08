## Finding

LowStock used < despite its inclusive threshold contract. Its test repeatedly exercised only the first table case.

## Evidence

The corrected test failed on the original implementation at thresholds 0 and 5, omitting GEAR-L and GEAR-S respectively. Repository-wide search found one affected comparison.

## Selected Candidate

The inclusive threshold bug: explicit contract, direct code evidence, and a regression test that distinguishes old from new behavior.

## Rejected Candidates

Build checks, inventory navigation, background synchronization, and arithmetic coverage overlap earlier PRs. Query batching and error wrapping were deferred to keep this change limited to one bug.

## Change

Changed < to <= and repaired LowStock coverage. Two files changed; changes remain uncommitted.

## Validation

Regression test failed before the fix. All tests and go vet ./... passed afterward; git diff --check passed. make check failed because scripts/lint.sh is missing, an existing issue covered by an earlier PR.

## Risk

Low: items exactly at the threshold now appear as documented.

## Rollback

Revert the changes in inventory/report.go and inventory/inventory_test.go.

## Follow-up

None
