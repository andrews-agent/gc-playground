## Finding

AddWarehouse replacement behavior lacked test coverage, unlike the equivalent AddItem replacement path.

## Evidence

AddWarehouse assigns the supplied record to s.warehouses[w.ID]. A mutation ignoring existing IDs passed the previous inventory suite but failed all three new replacement cases.

## Selected Candidate

Warehouse replacement coverage: a demonstrated test gap with a small fix following existing complete-record comparison conventions.

## Rejected Candidates

Stock-report batching and arithmetic coverage overlap excluded PRs. Documentation and navigation were current. Loop rewrites are cosmetic; unused-helper removal lacked comparable evidence.

## Change

Added 20 lines in inventory/store_test.go covering replacement of each sample warehouse and preservation of unrelated records.

## Validation

go mod download completed: no external dependencies. make check passed all tests and go vet using writable temporary caches; a macOS cache warning was nonfatal. gofmt and git diff --check passed. Mutation verification confirmed the new test detects ignored replacements. No checks remain blocked.

## Risk

Low; tests only.

## Rollback

Remove TestAddWarehouseReplacesExistingID.

## Follow-up

None
