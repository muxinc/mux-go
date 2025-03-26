package main

import (
	"log"
	"os"

	muxgo "github.com/muxinc/mux-go/v7"
)

func main() {
	log.Println("Starting...")

	muxClient := muxgo.NewAPIClient(
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

	asset, err := muxClient.AssetsApi.CreateAsset(createAssetRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created asset with ID: %s", asset.Data.Id)

	assets, err := muxClient.AssetsApi.ListAssets()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(assets.Data[0].Id)
}
