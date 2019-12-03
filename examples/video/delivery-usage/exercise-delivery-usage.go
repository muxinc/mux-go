package main

import (
	"fmt"
	"os"
	"time"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

func main() {

	t := int32(time.Now().Unix())
	fmt.Println(t)


	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

		// ========== list-delivery-usage ==========
		// OK, so here's the deal. Until we have actual meaningful data flowing into an account this is
		// really hard to to test, so instead I just create a timeframe that's falid and call the API.
		// We've manually verified this works during development.
		// To use specific times:
		// p := muxgo.ListDeliveryUsageParams{Timeframe: []string{"1574175600", "1574305200"}}
		// du, err := client.DeliveryUsageApi.ListDeliveryUsage(muxgo.WithParams(&p))
		// Default time window is 1 hour representing usage from 13th to 12th hour from when the request is made.
		du, err := client.DeliveryUsageApi.ListDeliveryUsage()
		common.AssertNoError(err)
		common.AssertNotNil(du.Data)
		fmt.Println("list-delivery-usage OK ✅")
}