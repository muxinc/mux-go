package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"golang.org/x/net/context"
)

// Exercises all asset operations:
//   get-asset
//   delete-asset
//   create-asset
//   list-assets
//   get-asset-input-info
//   create-asset-playback-id
//   get-asset-playback-id
//   delete-asset-playback-id
//   update-asset-mp4-support

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

	// API Client Initialization
	client := muxgo.NewAPIClient(muxgo.NewConfiguration())

	// ========== create-asset ==========
	car :=  muxgo.CreateAssetRequest{
				Input: []muxgo.InputSettings{muxgo.InputSettings{Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4"}},
			}

	_, _, err := client.AssetsApi.CreateAsset(auth, car)
	CheckError(err)
	fmt.Println("create-asset OK ✅")

	// ========== list-assets ==========
	_, _, err = client.AssetsApi.ListAssets(auth, nil)
	CheckError(err)
	fmt.Println("list-assets OK ✅")

		// ========== get-asset ==========
		// ========== get-asset-input-info ==========

	// ========== create-asset-playback-id ==========

	// ========== get-asset-playback-id ==========

	// ========== update-asset-mp4-support ==========

	// ========== delete-asset-playback-id ==========

	// ========== delete-asset ==========
}