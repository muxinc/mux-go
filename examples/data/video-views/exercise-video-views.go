package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all filter operations:
//   list-video-views
//   get-video-view

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== list-video-views ==========
	p := muxgo.ListVideoViewsParams{Filters: []string{"country:GB", "browser:Chrome"}, Timeframe: []string{"7:days"}}
	vs, err := client.VideoViewsApi.ListVideoViews(muxgo.WithParams(&p))
	common.AssertNoError(err)
	common.AssertNotNil(vs.Data)
	if len(vs.Data) < 1 {
		fmt.Println("No Video Views found.")
		os.Exit(255)
	}
	fmt.Println("list-video-views ✅")

	// ========== get-video-view ==========
	v, err := client.VideoViewsApi.GetVideoView(vs.Data[0].Id)
	common.AssertNoError(err)
	common.AssertNotNil(v.Data)
	fmt.Println("get-video-view ✅")
}