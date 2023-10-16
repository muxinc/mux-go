package main

import (
	"fmt"
	"os"
	"time"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all direct upload operations:
//   create-url-signing-key
//   list-url-signing-keys
//   get-url-signing-key
//   delete-url-signing-key


func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

		// ========== create-url-signing-key ==========
		k, err := client.URLSigningKeysApi.CreateUrlSigningKey()
		common.AssertNoError(err)
		common.AssertNotNil(k.Data)
		common.AssertNotNil(k.Data.PrivateKey)
		fmt.Println("create-url-signing-key OK ✅")

		// ========== list-url-signing-keys ==========
		ks, err := client.URLSigningKeysApi.ListUrlSigningKeys()
		common.AssertNoError(err)
		common.AssertNotNil(ks.Data)
		common.AssertStringEqualsValue(k.Data.Id, ks.Data[len(ks.Data)-1].Id)
		fmt.Println("list-url-signing-keys OK ✅")

		// ========== get-url-signing-key ==========
		kg, err := client.URLSigningKeysApi.GetUrlSigningKey(k.Data.Id)
		common.AssertNoError(err)
		common.AssertNotNil(kg.Data)
		common.AssertStringEqualsValue(k.Data.Id, kg.Data.Id)
		fmt.Println("get-url-signing-key OK ✅")

		// ========== delete-url-signing-key ==========
		err = client.URLSigningKeysApi.DeleteUrlSigningKey(k.Data.Id)
		common.AssertNoError(err)
		time.Sleep(60 * time.Second)
		_, err = client.URLSigningKeysApi.GetUrlSigningKey(k.Data.Id)
		common.AssertNotNil(err)
		fmt.Println("delete-url-signing-key OK ✅")
}