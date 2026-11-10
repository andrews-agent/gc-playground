## Finding

Move permits transferring exactly the available quantity, but no test exercised that boundary.

## Evidence

Changing the rejection guard from > to >= passed the original inventory suite. The extended test rejected that mutation with “only 85 left”.

## Selected Candidate

Close the demonstrated stock-exhaustion coverage gap with a one-line extension to the existing test.

## Rejected Candidates

Stock-report batching and arithmetic tests duplicate excluded work. CLI argument restrictions lack an established contract; removing the unused median helper offered weaker evidence of value.

## Change

Added transfers of the remaining 85 units and zero units afterward. Production code is unchanged.

## Validation

Ran go mod download: no dependencies to download. Mutation testing proved the new coverage distinguishes the boundary regression. Restored production code; make check passed all tests and go vet. git diff --check passed. Writable temporary and Go cache directories resolved initial sandbox failures; make emitted cache warnings but completed successfully.

## Risk

Minimal: test-only change using existing assertions.

## Rollback

Remove the two added transfer quantities.

## Follow-up

None
