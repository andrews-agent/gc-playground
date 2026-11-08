## Finding

Local documentation links must resolve. Two links in docs/inventory.md violated this invariant.

## Evidence

The original targets commands.md and inventory/reports.go do not exist. cmd/inventory/main.go contains usage and argument handling; inventory/report.go contains the reports.

## Selected Candidate

Broken documentation navigation: concrete, complete, and safe to fix without changing product behavior.

## Rejected Candidates

Considered ignored table-test cases, sleep-based synchronization, per-item warehouse queries, and the missing lint script. Deferred these separate cleanup classes to keep this change focused.

## Change

Updated both broken links in docs/inventory.md. One file changed, with two insertions and two deletions. Changes remain uncommitted.

## Validation

Confirmed both original targets fail existence checks. Checked all three Markdown files after editing: no broken local links remain. Go tests passed. git diff --check passed. make check failed at the pre-existing missing scripts/lint.sh.

## Risk

Low; documentation-only change.

## Rollback

Revert the two changed lines in docs/inventory.md.

## Follow-up

Restore or replace the missing lint script in a separate change.
