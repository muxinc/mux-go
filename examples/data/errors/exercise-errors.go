package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
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
	e, err := client.ErrorsApi.ListErrors()
	common.AssertNoError(err)
	common.AssertNotNil(e.Data)
	fmt.Println("list-errors ✅")
}