# Inventory

The `inventory` command prints reports over a small sample inventory. See the
[command source](../cmd/inventory/main.go) for usage and argument handling.

| Command | What it prints |
| --- | --- |
| `go run ./cmd/inventory stock` | every item with its warehouse |
| `go run ./cmd/inventory low [threshold]` | items at or below the threshold, 5 by default |
| `go run ./cmd/inventory show <sku>` | one item; prints "no such item or warehouse" for an unknown SKU |

`[threshold]` is an optional integer argument. For example,
`go run ./cmd/inventory low 0` lists only items with zero stock in the sample.

The code lives in [`inventory/`](../inventory) and the reports in
[`inventory/report.go`](../inventory/report.go).
