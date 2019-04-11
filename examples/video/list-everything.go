package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
}

func main() {
	// API Client Init
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("ACCESS_TOKEN_ID"), os.Getenv("ACCESS_TOKEN_SECRET")),
		))

	// List Assets
	fmt.Println("Listing Assets: \n")
	assets, err := client.AssetsApi.ListAssets()
	checkError(err)
	for _, asset := range assets.Data {
		fmt.Printf("Asset ID: %s\n", asset.Id)
		fmt.Printf("Status: %s\n", asset.Status)
		fmt.Printf("Duration: %f\n\n", asset.Duration)
	}

	// List Live Streams
	fmt.Println("Listing Live Streams: \n")
	streams, err := client.LiveStreamsApi.ListLiveStreams()
	checkError(err)
	for _, stream := range streams.Data {
		fmt.Printf("Live Stream ID: %s\n", stream.Id)
		fmt.Printf("Status: %s\n", stream.Status)
		fmt.Printf("Stream Key: %s\n\n", stream.StreamKey)
	}

	// List Direct Uploads
	fmt.Println("Listing Direct Uploads: \n")
	uploads, err := client.DirectUploadsApi.ListDirectUploads()
	checkError(err)
	for _, upload := range uploads.Data {
		fmt.Printf("Status: %s\n", upload.Status)
		fmt.Printf("New Asset ID: %s\n\n", upload.AssetId)
	}

	// List URL Signing Keys
	fmt.Println("Listing URL Signing Keys: \n")
	keys, err := client.URLSigningKeysApi.ListUrlSigningKeys()
	for _, key := range keys.Data {
		fmt.Printf("Signing Key ID: %s\n\n", key.Id)
	}
}
