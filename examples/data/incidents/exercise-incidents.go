package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// Test coverage here is poor due to not knowning if the account we're testing against has any incidents.

	// ========== list-incidents ==========
	i, err := client.IncidentsApi.ListIncidents()
	common.AssertNoError(err)
	common.AssertNotNil(i.Data)
	fmt.Println("list-incidents ✅")

	// ========== get-incident ==========
	fmt.Println("get-incident SKIP ⚠️")

	// ========== list-related-incidents ==========
	fmt.Println("list-related-incidents SKIP ⚠️")
}
