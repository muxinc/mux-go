package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go/v6"
	"github.com/muxinc/mux-go/v6/examples/common"
)

// Exercises all direct upload operations:
// create-direct-upload
// list-direct-uploads
// get-direct-upload
// cancel-direct-upload

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== create-direct-upload ==========
	car := muxgo.CreateAssetRequest{PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC}}
	cur := muxgo.CreateUploadRequest{NewAssetSettings: car, Timeout: 3600, CorsOrigin: "philcluff.co.uk"}
	u, err := client.DirectUploadsApi.CreateDirectUpload(cur)
	common.AssertNoError(err)
	common.AssertNotNil(u.Data)
	fmt.Println("create-direct-upload OK ✅")

	// ========== create-direct-upload ==========
	us, err := client.DirectUploadsApi.ListDirectUploads()
	common.AssertNoError(err)
	common.AssertNotNil(u.Data)
	common.AssertStringEqualsValue(u.Data.Id, us.Data[0].Id)
	fmt.Println("list-direct-uploads ✅")

	// ========== get-direct-upload ==========
	ug, err := client.DirectUploadsApi.GetDirectUpload(u.Data.Id)
	common.AssertNoError(err)
	common.AssertNotNil(ug.Data)
	common.AssertStringEqualsValue(u.Data.Id, ug.Data.Id)
	fmt.Println("get-direct-upload ✅")

	// ========== cancel-direct-upload ==========
	uc, err := client.DirectUploadsApi.CancelDirectUpload(u.Data.Id)
	common.AssertNoError(err)
	common.AssertNotNil(uc.Data)
	common.AssertStringEqualsValue(u.Data.Id, uc.Data.Id)
	common.AssertStringEqualsValue(uc.Data.Status, "cancelled")
	fmt.Println("cancel-direct-upload ✅")
}
