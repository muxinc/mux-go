package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go/v6"
	"github.com/muxinc/mux-go/v6/examples/common"
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
	p := muxgo.ListVideoViewsParams{Filters: []string{"country:US", "browser:Safari"}, Timeframe: []string{"7:days"}}
	vs, err := client.VideoViewsApi.ListVideoViews(muxgo.WithParams(&p))
	common.AssertNoError(err)
	common.AssertNotNil(vs.Data)
	if len(vs.Data) < 1 {
		fmt.Println("No Video Views found.")
		os.Exit(255)
	}
	fmt.Println("list-video-views ✅")

	fmt.Println(vs.Data[0].Id)

	// ========== get-video-view ==========
	v, err := client.VideoViewsApi.GetVideoView(vs.Data[0].Id)
	common.AssertNoError(err)
	common.AssertNotNil(v.Data)
	fmt.Println("get-video-view ✅")

	fmt.Println(v.Data.Events)
}
