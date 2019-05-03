package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all filter operations:
//   list-filters
//   list-filter-values

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-filters ==========
	f, err := client.FiltersApi.ListFilters()
	common.AssertNoError(err)
	common.AssertNotNil(f.Data)
	common.AssertNotNil(f.Data.Basic)
	common.AssertNotNil(f.Data.Advanced)
	fmt.Println("list-filters ✅")

	// ========== list-filter-values ==========
	fp := muxgo.ListFilterValuesParams{Timeframe: []string{"7:days"}}
	fv, err := client.FiltersApi.ListFilterValues("browser", muxgo.WithParams(&fp))
	common.AssertNoError(err)
	common.AssertNotNil(fv.Data)
	fmt.Println("list-filter-values ✅")
}