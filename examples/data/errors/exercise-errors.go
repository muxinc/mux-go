package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-errors ==========
	e, err := client.ErrorsApi.ListErrors()
	common.AssertNoError(err)
	fmt.Println(e)
	fmt.Println("list-errors ✅")
}