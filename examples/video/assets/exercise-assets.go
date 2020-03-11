package main

import (
	"fmt"
	"os"
	"time"

	muxgo "github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// ========== create-asset ==========
	asset, err := client.AssetsApi.CreateAsset(muxgo.CreateAssetRequest{
		Input: []muxgo.InputSettings{
			muxgo.InputSettings{
				Url: "https://storage.googleapis.com/muxdemofiles/mux-video-intro.mp4",
			},
			muxgo.InputSettings{
				Url:            "https://tears-of-steel-subtitles.s3.amazonaws.com/tears-fr.vtt",
				TextType:       "subtitles",
				Type:           "text",
				LanguageCode:   "fr",
				ClosedCaptions: false,
				Name:           "French",
			},
		},
		NormalizeAudio: true,
	})
	common.AssertNoError(err)
	common.AssertNotNil(asset.Data)
	fmt.Println("create-asset OK ✅")

	// ========== list-assets ==========
	lr, err := client.AssetsApi.ListAssets()
	common.AssertNoError(err)
	common.AssertNotNil(lr.Data)
	common.AssertStringEqualsValue(asset.Data.Id, lr.Data[0].Id)
	fmt.Println("list-assets OK ✅")

	// Wait for the asset to become ready
	if asset.Data.Status != "ready" {
		fmt.Println("    Waiting for asset to become ready...")
		for {
			// ========== get-asset ==========
			gr, err := client.AssetsApi.GetAsset(asset.Data.Id)
			common.AssertNoError(err)
			common.AssertNotNil(gr)
			common.AssertStringEqualsValue(asset.Data.Id, gr.Data.Id)
			if gr.Data.Status != "ready" {
				fmt.Println("    Asset still not ready.")
				time.Sleep(1 * time.Second)
			} else {
				// ========== get-asset-input-info ==========
				ir, err := client.AssetsApi.GetAssetInputInfo(asset.Data.Id)
				common.AssertNoError(err)
				common.AssertNotNil(ir.Data)
				break
			}
		}
	}
	fmt.Println("get-asset OK ✅")
	fmt.Println("get-asset-input-info OK ✅")

	// ========== create-asset-playback-id ==========
	capr := muxgo.CreatePlaybackIdRequest{muxgo.PUBLIC}
	capre, err := client.AssetsApi.CreateAssetPlaybackId(asset.Data.Id, capr)
	common.AssertNoError(err)
	common.AssertNotNil(capre.Data)
	common.AssertStringEqualsValue(string(capre.Data.Policy), "public")
	fmt.Println("create-asset-playback-id OK ✅")

	// ========== get-asset-playback-id ==========
	pbre, err := client.AssetsApi.GetAssetPlaybackId(asset.Data.Id, capre.Data.Id)
	common.AssertNoError(err)
	common.AssertNotNil(pbre.Data)
	common.AssertStringEqualsValue(capre.Data.Id, pbre.Data.Id)
	fmt.Println("get-asset-playback-id OK ✅")

	// ========== update-asset-mp4-support ==========
	mp4r := muxgo.UpdateAssetMp4SupportRequest{"standard"}
	mp4, err := client.AssetsApi.UpdateAssetMp4Support(asset.Data.Id, mp4r)
	common.AssertNoError(err)
	common.AssertNotNil(mp4.Data)
	common.AssertStringEqualsValue(asset.Data.Id, mp4.Data.Id)
	common.AssertStringEqualsValue(mp4.Data.Mp4Support, "standard")
	fmt.Println("update-asset-mp4-support OK ✅")

	// ========== update-asset-master-access ==========
	mr := muxgo.UpdateAssetMasterAccessRequest{"temporary"}
	mas, err := client.AssetsApi.UpdateAssetMasterAccess(asset.Data.Id, mr)
	common.AssertNoError(err)
	common.AssertNotNil(mas.Data)
	common.AssertStringEqualsValue(asset.Data.Id, mas.Data.Id)
	common.AssertStringEqualsValue(mas.Data.MasterAccess, "temporary")
	fmt.Println("update-asset-master-access OK ✅")

	// ========== create-asset-track ==========
	cat := muxgo.CreateTrackRequest{
		Url:            "https://tears-of-steel-subtitles.s3.amazonaws.com/tears-en.vtt",
		TextType:       "subtitles",
		Type:           "text",
		LanguageCode:   "en",
		ClosedCaptions: false,
		Name:           "English",
	}
	s, err := client.AssetsApi.CreateAssetTrack(asset.Data.Id, cat)
	common.AssertNoError(err)
	common.AssertNotNil(s.Data)
	common.AssertNotNil(s.Data.Id)
	common.AssertStringEqualsValue(s.Data.Name, "English")
	a2s, err := client.AssetsApi.GetAsset(asset.Data.Id)
	common.AssertIntEqualsValue(len(a2s.Data.Tracks), 4) // Audio, Video, French that we ingested with the asset, and the English we added here!
	fmt.Println("create-asset-track OK ✅")

	// ========== delete-asset-track ==========
	err = client.AssetsApi.DeleteAssetTrack(asset.Data.Id, s.Data.Id)
	common.AssertNoError(err)
	a1s, err := client.AssetsApi.GetAsset(asset.Data.Id)
	common.AssertIntEqualsValue(len(a1s.Data.Tracks), 3) // Audio, Video, French that we ingested with the asset
	fmt.Println("delete-asset-track OK ✅")

	// ========== delete-asset-playback-id ==========
	err = client.AssetsApi.DeleteAssetPlaybackId(asset.Data.Id, capre.Data.Id)
	common.AssertNoError(err)
	dpba, err := client.AssetsApi.GetAsset(asset.Data.Id)
	common.AssertNoError(err)
	if len(dpba.Data.PlaybackIds) > 0 {
		fmt.Println("List of playback IDs wasn't empty")
		os.Exit(255)
	}
	fmt.Println("delete-asset-playback-id OK ✅")

	// ========== delete-asset ==========
	err = client.AssetsApi.DeleteAsset(asset.Data.Id)
	common.AssertNoError(err)
	_, err = client.AssetsApi.GetAsset(asset.Data.Id)
	common.AssertNotNil(err)
	fmt.Println("delete-asset OK ✅")
}
