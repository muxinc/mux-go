package main

import (
	"fmt"
	"os"

	"github.com/muxinc/mux-go"
	"github.com/muxinc/mux-go/examples/common"
)

// Exercises all direct upload operations:
//   create-live-stream
//   list-live-streams
//   get-live-stream
//   delete-live-stream
//   create-live-stream-playback-id
//   delete-live-stream-playback-id
//   reset-stream-key
//   signal-live-stream-complete

func main() {

	// API Client Initialization
	client := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

		// ========== create-live-stream ==========
		car := muxgo.CreateAssetRequest{PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC}}
		csr := muxgo.CreateLiveStreamRequest{NewAssetSettings: car, PlaybackPolicy: []muxgo.PlaybackPolicy{muxgo.PUBLIC}}
		s, err := client.LiveStreamsApi.CreateLiveStream(csr)
		common.AssertNoError(err)
		common.AssertNotNil(s.Data)
		fmt.Println("create-live-stream OK ✅")
		
		// ========== list-live-streams ==========
		ss, err := client.LiveStreamsApi.ListLiveStreams()
		common.AssertNoError(err)
		common.AssertNotNil(ss.Data)
		common.AssertStringEqualsValue(s.Data.Id, ss.Data[0].Id)
		fmt.Println("list-live-streams OK ✅")

		// ========== get-live-stream ==========
		gs, err := client.LiveStreamsApi.GetLiveStream(s.Data.Id)
		common.AssertNoError(err)
		common.AssertNotNil(gs.Data)
		common.AssertStringEqualsValue(s.Data.Id, gs.Data.Id)
		fmt.Println("get-live-stream OK ✅")
		
		// ========== create-live-stream-playback-id ==========
		cpbidr := muxgo.CreatePlaybackIdRequest{Policy: muxgo.SIGNED}
		cpbids, err := client.LiveStreamsApi.CreateLiveStreamPlaybackId(s.Data.Id, cpbidr)
		common.AssertNoError(err)
		common.AssertNotNil(cpbids.Data)
		common.AssertStringEqualsValue(string(cpbids.Data.Policy), "signed")
		fmt.Println("create-live-stream-playback-id OK ✅")

		// ========== delete-live-stream-playback-id ==========
		err = client.LiveStreamsApi.DeleteLiveStreamPlaybackId(s.Data.Id, cpbids.Data.Id)
		common.AssertNoError(err)
		s1pbid, err := client.LiveStreamsApi.GetLiveStream(s.Data.Id)
		common.AssertNoError(err)
		common.AssertNotNil(s1pbid.Data)
		common.AssertStringEqualsValue(string(s1pbid.Data.PlaybackIds[0].Policy), "public")
		fmt.Println("delete-live-stream-playback-id OK ✅")

		// ========== reset-stream-key ==========
		sk, err := client.LiveStreamsApi.ResetStreamKey(s.Data.Id)
		common.AssertNoError(err)
		common.AssertNotNil(sk.Data)
		if s.Data.StreamKey == sk.Data.StreamKey {
			fmt.Println("Stream Key didn't change on stream key reset.")
			os.Exit(255)
		}
		fmt.Println("reset-stream-key OK ✅")

		// ========== signal-live-stream-complete ==========
		_, err = client.LiveStreamsApi.SignalLiveStreamComplete(s.Data.Id)
		common.AssertNoError(err)
		fmt.Println("signal-live-stream-complete OK ✅")

		// ========== delete-live-stream ==========
		err = client.LiveStreamsApi.DeleteLiveStream(s.Data.Id)
		common.AssertNoError(err)
		_, err = client.LiveStreamsApi.GetLiveStream(s.Data.Id)
		common.AssertNotNil(err)
		fmt.Println("delete-live-stream OK ✅")
		
}