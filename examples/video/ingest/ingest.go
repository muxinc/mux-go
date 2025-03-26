package main

import (
	"fmt"
	"log"
	"os"

	muxgo "github.com/muxinc/mux-go/v7"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	createAssetRequest := muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			{Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4"},
		},
		PlaybackPolicy:   []muxgo.PlaybackPolicy{muxgo.PUBLIC},
		StaticRenditions: []muxgo.CreateStaticRenditionRequest{{Resolution: "highest"}},
	}

	asset, err := client.AssetsApi.CreateAsset(createAssetRequest)
	if err != nil {
		log.Fatal(err)
	}

	if err == nil {
		fmt.Printf("HLS Playback URL: https://stream.mux.com/%s.m3u8 \n", asset.Data.PlaybackIds[0].Id)
		fmt.Printf("MP4 Playback URL: https://stream.mux.com/%s/highest.mp4 \n", asset.Data.PlaybackIds[0].Id)
	} else {
		fmt.Printf("Oh no, there was an error: %s \n", err)
	}

}
