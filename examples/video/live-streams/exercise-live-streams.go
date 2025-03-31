package main

import (
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go/v7"
	"github.com/muxinc/mux-go/v7/examples/common"
)

// Exercises all live stream operations:
//   create-live-stream
//   list-live-streams
//   get-live-stream
//   delete-live-stream
//   create-live-stream-playback-id
//   delete-live-stream-playback-id
//   reset-stream-key
//   signal-live-stream-complete
//   create-live-stream-simulcast-target
//   get-live-stream-simulcast-target
//   delete-live-stream-simulcast-target

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

	// ========== get-asset-or-livestream-id =========
	playbackId := s.Data.PlaybackIds[0].Id
	pbResp, err := client.PlaybackIDApi.GetAssetOrLivestreamId(playbackId)
	common.AssertNoError(err)
	common.AssertNotNil(pbResp.Data)
	common.AssertStringEqualsValue(pbResp.Data.Object.Id, s.Data.Id)
	common.AssertStringEqualsValue(pbResp.Data.Object.Type, "live_stream")
	fmt.Println("get-asset-or-livestream-id OK ✅")

	// ========== create-live-stream-simulcast-target ==========
	cst := muxgo.CreateSimulcastTargetRequest{Passthrough: "foo", StreamKey: "bar", Url: "rtmp://this-is-a.test"}
	nst, err := client.LiveStreamsApi.CreateLiveStreamSimulcastTarget(s.Data.Id, cst)
	common.AssertNoError(err)
	common.AssertNotNil(nst.Data)
	fmt.Println("create-live-stream-simulcast-target OK ✅")

	// ========== get-live-stream-simulcast-target ==========
	st, err := client.LiveStreamsApi.GetLiveStreamSimulcastTarget(s.Data.Id, nst.Data.Id)
	common.AssertNoError(err)
	common.AssertNotNil(st.Data)
	common.AssertStringEqualsValue(nst.Data.Id, st.Data.Id)
	fmt.Println("get-live-stream-simulcast-target OK ✅")

	// ========== delete-live-stream-simulcast-target ==========
	err = client.LiveStreamsApi.DeleteLiveStreamSimulcastTarget(s.Data.Id, nst.Data.Id)
	common.AssertNoError(err)
	// Check it actually got deleted
	snost, err := client.LiveStreamsApi.GetLiveStream(s.Data.Id)
	common.AssertNoError(err)
	common.AssertNotNil(snost.Data)
	common.AssertNil(snost.Data.SimulcastTargets)
	fmt.Println("delete-live-stream-simulcast-target OK ✅")

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

	// ========== disable-live-stream ==========
	_, err = client.LiveStreamsApi.DisableLiveStream(s.Data.Id)
	common.AssertNoError(err)
	dls, err := client.LiveStreamsApi.GetLiveStream(s.Data.Id)
	common.AssertStringEqualsValue(string(dls.Data.Status), "disabled")
	fmt.Println("disable-live-stream OK ✅")

	// ========== enable-live-stream ==========
	_, err = client.LiveStreamsApi.EnableLiveStream(s.Data.Id)
	common.AssertNoError(err)
	els, err := client.LiveStreamsApi.GetLiveStream(s.Data.Id)
	common.AssertStringEqualsValue(string(els.Data.Status), "idle")
	fmt.Println("enable-live-stream OK ✅")

	// ========== delete-live-stream ==========
	err = client.LiveStreamsApi.DeleteLiveStream(s.Data.Id)
	common.AssertNoError(err)
	_, err = client.LiveStreamsApi.GetLiveStream(s.Data.Id)
	common.AssertNotNil(err)
	fmt.Println("delete-live-stream OK ✅")

}
