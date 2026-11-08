// Command inventory prints reports over the sample inventory.
//
// Usage:
//
//	inventory stock
//	inventory low [threshold]
//	inventory show <sku>
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/andrews-agent/gc-playground/inventory"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: inventory stock | low [threshold] | show <sku>")
		os.Exit(2)
	}
	s := inventory.Sample()
	var out string
	var err error
	switch os.Args[1] {
	case "stock":
		out, err = inventory.StockReport(s)
	case "low":
		threshold := 5
		if len(os.Args) > 2 {
			threshold, _ = strconv.Atoi(os.Args[2])
		}
		out, err = inventory.LowStock(s, threshold)
	case "show":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "usage: inventory show <sku>")
			os.Exit(2)
		}
		out, err = inventory.ItemLine(s, os.Args[2])
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", os.Args[1])
		os.Exit(2)
	}
	if errors.Is(err, inventory.ErrNotFound) {
		fmt.Fprintln(os.Stderr, "no such item or warehouse")
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println(out)
}
