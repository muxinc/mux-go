package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"golang.org/x/net/context"
)

// To run: go run go-demo.go

// Dependencies:
// go get github.com/stretchr/testify/assert
// go get golang.org/x/oauth2
// go get golang.org/x/net/context
// go get github.com/antihax/optional

func main() {

	// Auth Setup
	auth := context.WithValue(context.Background(), muxgo.ContextBasicAuth, muxgo.BasicAuth{
		UserName: os.Getenv("MUX_TOKEN_ID"),
		Password: os.Getenv("MUX_TOKEN_SECRET"),
	})

	// API Client Init
	client := muxgo.NewAPIClient(muxgo.NewConfiguration())

	// List Assets
	fmt.Println("Listing Assets in account: \n")
	lasr, _, err := client.AssetsApi.ListAssets(auth, nil)
	if err != nil {
		fmt.Printf("err: %s \n\n", err)
		os.Exit(255)
	}
	for _, asset := range lasr.Data {
		fmt.Printf("Asset ID: %s\n", asset.Id)
		fmt.Printf("Status: %s\n", asset.Status)
		fmt.Printf("Duration: %f\n\n", asset.Duration)
	}
}
