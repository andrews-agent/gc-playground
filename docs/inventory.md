# Inventory

The `inventory` command prints reports over a small sample inventory. See the
[command source](../cmd/inventory/main.go) for usage and argument handling.

| Command | What it prints |
| --- | --- |
| `go run ./cmd/inventory stock` | every item with its warehouse |
| `go run ./cmd/inventory low` | items at or below the threshold, 5 by default |
| `go run ./cmd/inventory show <sku>` | one item; prints "no such item or warehouse" for an unknown SKU |

The code lives in [`inventory/`](../inventory) and the reports in
[`inventory/report.go`](../inventory/report.go).
