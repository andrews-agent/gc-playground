## Finding

TestMoveInBackground accessed shared store state without waiting for the background mutation to complete.

## Evidence

Before the fix, the targeted race-detector test failed with races on Store.Queries and the item map. Repository-wide searches found no other sleep-based synchronization.

## Selected Candidate

The background test race had a reproducible failure and a small, test-only fix.

## Rejected Candidates

Missing lint script, low-stock test defects, and report query inefficiency are separate cleanup classes. Previously reported documentation and arithmetic-test issues were excluded.

## Change

Updated inventory/inventory_test.go to receive the background move result through a channel before reading the store. Changes remain uncommitted.

## Validation

go test -race ./... passed after the fix. make check passed its test stage but failed because scripts/lint.sh is missing. gofmt and git diff --check passed. Go checks used writable temporary cache paths.

## Risk

Low: only test synchronization and error assertions changed; product behavior is unchanged.

## Rollback

Revert the changes to inventory/inventory_test.go.

## Follow-up

Restore or replace the missing lint script in a separate change.
