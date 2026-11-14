## Finding

The guide mentioned only unknown SKUs, while the command emits the same error when the item's warehouse is missing.

## Evidence

ItemLine performs both item and warehouse lookups, and the CLI maps either ErrNotFound result to the same message.

## Selected Candidate

A one-line documentation correction with direct code evidence and no behavior change.

## Rejected Candidates

Excluded query batching and arithmetic test changes because they overlap earlier open pull requests; no other distinct, well-grounded issue was found.

## Change

Updated docs/inventory.md to mention unknown SKUs or warehouses.

## Validation

go mod download completed with no dependencies. make check ran tests and vet but emitted a macOS cache-permission diagnostic; direct go test ./... and go vet ./... passed with writable workspace caches. git diff --check passed.

## Risk

None beyond documentation accuracy.

## Rollback

Revert the one-line change in docs/inventory.md.

## Follow-up

None
