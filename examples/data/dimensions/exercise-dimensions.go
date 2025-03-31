package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go/v7"
	"github.com/muxinc/mux-go/v7/examples/common"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-dimensions ==========
	d, err := client.DimensionsApi.ListDimensions()
	common.AssertNoError(err)
	common.AssertNotNil(d.Data)
	common.AssertNotNil(d.Data.Basic)
	common.AssertNotNil(d.Data.Advanced)
	fmt.Println("list-dimensions ✅")

	// ========== list-dimension-values ==========
	ldp := muxgo.ListDimensionValuesParams{Timeframe: []string{"7:days"}}
	dv, err := client.DimensionsApi.ListDimensionValues("browser", muxgo.WithParams(&ldp))
	common.AssertNoError(err)
	common.AssertNotNil(dv.Data)
	fmt.Println("list-dimension-values ✅")
}
