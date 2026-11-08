## Finding

Public API documentation omitted Max's nonempty-input requirement.

## Evidence

stats/stats.go initializes max from numbers[0] unconditionally. The updated go doc output explicitly describes the resulting panic.

## Selected Candidate

Incomplete API documentation: directly supported by code and safely fixed with one sentence. All exported helpers were inspected for the same omission.

## Rejected Candidates

Changing Max's empty-input result lacks an established contract. Sum loop changes are cosmetic. Removing the unused median helper offered weaker justification. README commands are current.

## Change

Added one documentation line in stats/stats.go. Changes remain uncommitted.

## Validation

go vet ./..., go test ./..., and git diff --check passed. go doc with a writable GOCACHE confirmed the updated documentation; the initial documentation invocation failed to resolve the package.

## Risk

Minimal; documentation only.

## Rollback

Remove the added comment sentence.

## Follow-up

None
