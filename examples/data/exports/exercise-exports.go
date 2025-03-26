package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go/v7"
	"github.com/muxinc/mux-go/v7/examples/common"
)

// Exercises all export operations:
//   list-exports

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-exports ==========
	e, err := client.ExportsApi.ListExportsViews()
	common.AssertNoError(err)
	common.AssertNotNil(e.Data)
	if len(e.Data) < 1 {
		fmt.Println("Didn't find any exports :( ")
		os.Exit(255)
	}
	fmt.Println("list-exports ✅")
}
