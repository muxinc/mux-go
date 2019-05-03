package main

import (
	"fmt"
	"os"
	"github.com/muxinc/mux-go"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// Create Asset
	asset, err := client.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			muxgo.InputSettings{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
		},
		PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC},
	})
	
	if err == nil {
		fmt.Printf("Playback URL: https://stream.mux.com/%s.m3u8 \n", asset.Data.PlaybackIds[0].Id)
	} else {
		fmt.Printf("Oh no, there was an error: %s \n", err)
	}

}
