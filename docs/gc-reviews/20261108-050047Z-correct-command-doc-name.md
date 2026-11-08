## Finding

Command documentation should identify the command it describes. main.go incorrectly named it stats.

## Evidence

go.mod and README.md name gc-playground; go list reports an executable target ending in gc-playground. The old comment named stats.

## Selected Candidate

Stale command documentation: directly verifiable and fixable without changing behavior. A repository-wide search found one instance.

## Rejected Candidates

Unused median removal lacked the required supporting evidence. Sum loop changes were cosmetic. Max empty-input documentation overlaps the excluded open PR. Existing README commands are correct.

## Change

Changed one comment in main.go. Left the change uncommitted.

## Validation

go vet ./..., go test ./..., and git diff --check passed using writable temporary Go caches. The README example printed sum 6 and mean 2. Reviewed the complete diff.

## Risk

Minimal; documentation only.

## Rollback

Restore the previous first line of main.go.

## Follow-up

None
