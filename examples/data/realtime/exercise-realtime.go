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

	// Test coverage here isn't fantastic due to not knowning if the account we're testing against has
	// any real-time data. The behaviour has been manually verified against real-world data.

	// ========== list-realtime-dimensions ==========
	d, err := client.RealTimeApi.ListRealtimeDimensions()
	common.AssertNoError(err)
	common.AssertNotNil(d.Data)
	common.AssertIntGreaterThanZero(len(d.Data))
	common.AssertStringNotEqualsValue(d.Data[0].Name, "")
	common.AssertStringNotEqualsValue(d.Data[0].DisplayName, "")
	fmt.Println("list-realtime-dimensions ✅")

	// ========== list-realtime-metrics ==========
	m, err := client.RealTimeApi.ListRealtimeMetrics()
	common.AssertNoError(err)
	common.AssertNotNil(m.Data)
	common.AssertIntGreaterThanZero(len(m.Data))
	common.AssertStringNotEqualsValue(m.Data[0].Name, "")
	common.AssertStringNotEqualsValue(m.Data[0].DisplayName, "")
	fmt.Println("list-realtime-metrics ✅")

	// ========== get-realtime-breakdown ==========
	rbp := muxgo.GetRealtimeBreakdownParams{Dimension: "asn"}
	b, err := client.RealTimeApi.GetRealtimeBreakdown("current-rebuffering-percentage", muxgo.WithParams(&rbp))
	common.AssertNoError(err)
	common.AssertNotNil(b.Data)
	fmt.Println("get-realtime-breakdown ✅")

	// ========== get-realtime-histogram-timeseries ==========
	ht, err := client.RealTimeApi.GetRealtimeHistogramTimeseries("video-startup-time")
	common.AssertNoError(err)
	common.AssertNotNil(ht.Meta)
	common.AssertNotNil(ht.Meta.Buckets)
	common.AssertIntGreaterThanZero(len(ht.Meta.Buckets))
	common.AssertNotNil(ht.Data)
	common.AssertIntGreaterThanZero(len(ht.Data))
	fmt.Println("get-realtime-histogram-timeseries ✅")

	// ========== get-realtime-timeseries ==========
	t, err := client.RealTimeApi.GetRealtimeTimeseries("current-rebuffering-percentage")
	common.AssertNoError(err)
	common.AssertNotNil(t.Data)
	common.AssertIntGreaterThanZero(len(t.Data))
	common.AssertStringNotEqualsValue(t.Data[0].Date, "")
	fmt.Println("get-realtime-timeseries ✅")

}
