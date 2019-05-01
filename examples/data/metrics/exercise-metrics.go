package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all metrics operations:
//   list-breakdown-values
//   get-overall-values
//   list-insights
//   get-metric-timeseries-data
//   list-all-metric-values

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-breakdown-values ==========
	bdvopts := muxgo.ListBreakdownValuesParams{GroupBy: "browser", Timeframe: []string{"7:days"}}
	bdv, err := client.MetricsApi.ListBreakdownValues("video_startup_time", muxgo.WithParams(&bdvopts))
	common.AssertNoError(err)
	common.AssertNotNil(bdv.Data)
	fmt.Println("list-breakdown-values ✅")
	
	// ========== get-overall-values ==========
	ovopts := muxgo.GetOverallValuesParams{Timeframe: []string{"7:days"}}
	ov, err := client.MetricsApi.GetOverallValues("video_startup_time", muxgo.WithParams(&ovopts))
	common.AssertNoError(err)
	common.AssertNotNil(ov.Data)
	fmt.Println("get-overall-values ✅")

	// ========== list-insights ==========
	iopts := muxgo.ListInsightsParams{Timeframe: []string{"7:days"}}
	is, err := client.MetricsApi.ListInsights("video_startup_time", muxgo.WithParams(&iopts))
	common.AssertNoError(err)
	common.AssertNotNil(is.Data)
	fmt.Println("list-insights ✅")

	// ========== get-metric-timeseries-data ==========
	tsopts := muxgo.GetMetricTimeseriesDataParams{Timeframe: []string{"7:days"}}
	ts, err := client.MetricsApi.GetMetricTimeseriesData("video_startup_time", muxgo.WithParams(&tsopts))
	common.AssertNoError(err)
	common.AssertNotNil(ts.Data)
	fmt.Println("get-metic-timeseries-data ✅")

	// ========== list-all-metric-values ==========
	lamv, err := client.MetricsApi.ListAllMetricValues()
	common.AssertNoError(err)
	common.AssertNotNil(lamv.Data)
	fmt.Println("list-all-metric-values ✅")
}