package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"golang.org/x/net/context"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
}

func main() {

	// Authentication Setup
	auth := context.WithValue(context.Background(), muxgo.ContextBasicAuth, muxgo.BasicAuth{
		UserName: os.Getenv("MUX_TOKEN_ID"),
		Password: os.Getenv("MUX_TOKEN_SECRET"),
	})

	// API Client Init
	client := muxgo.NewAPIClient(muxgo.NewConfiguration())

	// List Assets
	fmt.Println("Listing Assets: \n")
	assets, _, err := client.AssetsApi.ListAssets(auth, nil)
	CheckError(err)
	for _, asset := range assets.Data {
		fmt.Printf("Asset ID: %s\n", asset.Id)
		fmt.Printf("Status: %s\n", asset.Status)
		fmt.Printf("Duration: %f\n\n", asset.Duration)
	}

	// List Live Streams
	fmt.Println("Listing Live Streams: \n")
	streams, _, err := client.LiveStreamsApi.ListLiveStreams(auth, nil)
	CheckError(err)
	for _, stream := range streams.Data {
		fmt.Printf("Live Stream ID: %s\n", stream.Id)
		fmt.Printf("Status: %s\n", stream.Status)
		fmt.Printf("Stream Key: %s\n\n", stream.StreamKey)
	}

	// List Direct Uploads
	fmt.Println("Listing Direct Uploads: \n")
	uploads, _, err := client.DirectUploadsApi.ListDirectUploads(auth, nil)
	CheckError(err)
	for _, upload := range uploads.Data {
		fmt.Printf("Status: %s\n", upload.Status)
		fmt.Printf("New Asset ID: %s\n\n", upload.AssetId)
	}

	// List URL Signing Keys
	fmt.Println("Listing URL Signing Keys: \n")
	keys, _, err := client.URLSigningKeysApi.ListUrlSigningKeys(auth, nil)
	for _, key := range keys.Data {
		fmt.Printf("Signing Key ID: %s\n\n", key.Id)
	}
}
