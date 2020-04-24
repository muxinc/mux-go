package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all error operations:
//   list-errors

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-errors ==========
	lep := muxgo.ListErrorsParams{Filters: []string{"browser:Safari"}, Timeframe: []string{"7:days"}}
	e, err := client.ErrorsApi.ListErrors(muxgo.WithParams(&lep))
	common.AssertNoError(err)
	common.AssertNotNil(e.Data)
	if len(e.Data) < 1 {
		fmt.Println("Didn't find any errors :( ")
		os.Exit(255)
	}
	fmt.Println("list-errors ✅")
}
