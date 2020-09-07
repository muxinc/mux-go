package main

import (
	"encoding/json"
	"fmt"
	"os"

	muxgo "github.com/muxinc/mux-go"
)

func main() {

	s := `{
	"video_startup_preroll_load_time": 1234,
	"player_source_url": "https://stream.mux.com",
	"mux_api_version": "test",
	"asn": 1,
	"view_total_downscaling": "1234",
	"view_max_downscale_percentage": "0.001",
	"viewer_application_version": "1.2.3",
	"player_source_type": "hls",
	"view_average_request_throughput": 1000,
	"viewer_device_category": "test",
	"error_type_id": null,
	"preroll_ad_tag_hostname": "true",
	"player_error_code": "0",
	"mux_embed_version": "test",
	"time_to_first_frame": 1000,
	"platform_description": "test",
	"view_total_upscaling": "1234",
	"mux_viewer_id": "test",
	"view_total_content_playback_time": 12345,
	"player_poster": "test",
	"watched": true,
	"latitude": "10.1",
	"player_language": "en-us",
	"player_error_message": "test",
	"page_load_time": 1000,
	"viewer_os_architecture": "test",
	"cdn": "test",
	"city": "test",
	"player_version": "1.2.3",
	"view_max_request_latency": 1000,
	"viewer_application_name": "Chrome",
	"rebuffer_percentage": "0.001",
	"longitude": "10.1",
	"player_source_width": 640,
	"viewer_user_id": "1234",
	"player_source_stream_type": "hls",
	"session_id": "test",
	"video_content_type": "type",
	"viewer_os_version": "test",
	"view_average_request_latency": 1000,
	"short_time": " 5:22am Etc/UTC 2020-08-25",
	"player_software": "video.js",
	"region": "test",
	"requests_for_first_preroll": 1,
	"video_producer": "producer",
	"platform_summary": "Chrome on test",
	"page_url": "https://localhost",
	"player_mux_plugin_version": "1.2.3",
	"player_software_version": "1.2.3",
	"used_fullscreen": true,
	"player_startup_time": 1234,
	"video_language": "en-us",
	"view_seek_count": 1,
	"viewer_user_agent": "test",
	"buffering_rate": "0.001",
	"player_source_host_name": "test",
	"player_source_domain": "test",
	"startup_score": "0.1",
	"video_variant_name": "variant",
	"property_id": 41454,
	"video_duration": 1000,
	"buffering_duration": 1000,
	"metro": "test",
	"continent_code": "test",
	"watch_time": 1000,
	"player_name": "test",
	"viewer_experience_score": "0.1",
	"quality_score": "0.1",
	"experiment_name": "test",
	"video_id": "1234",
	"asn_name": "LVLT-1 - Level 3 Parent, LLC (AS1)",
	"page_type": "test",
	"video_encoding_variant": "variant",
	"preroll_requested": true,
	"rebuffering_score": "0.1",
	"player_height": 360,
	"video_variant_id": "variant_id",
	"viewer_application_engine": "chromium",
	"viewer_device_name": "test",
	"country_code": "test",
	"video_title": "Video Title",
	"player_source_duration": 1000,
	"playback_score": "0.1",
	"view_max_upscale_percentage": "0.001",
	"viewer_os_family": "test",
	"preroll_ad_asset_hostname": "true",
	"events": [],
	"country_name": "test",
	"player_preload": true,
	"video_stream_type": "live",
	"isp": "test",
	"player_source_height": 360,
	"player_width": 640,
	"exit_before_video_start": true,
	"player_view_count": 0,
	"player_instance_id": "test",
	"buffering_count": 1,
	"player_autoplay": true,
	"view_seek_duration": 1000,
	"player_load_time": 1000,
	"view_error_id": null,
	"video_startup_preroll_request_time": 1234,
	"preroll_played": true,
	"viewer_device_manufacturer": "test",
	"video_series": "series",
	"player_mux_plugin_name": "test",
	"updated_at": "2020-08-25T05:22:03Z",
	"view_start": "2020-08-25T05:22:03Z",
	"view_id": "dd36b9c5-c456-483e-b6ef-6a8b2b05d862",
	"inserted_at": "2020-08-25T05:22:03Z",
	"view_end": "2020-08-25T05:22:03Z",
	"id": "892e4239-1540-4014-8f84-6a7e67fea57b"
  }`

	var vv muxgo.VideoView

	err := json.Unmarshal([]byte(s), &vv)

	if err != nil {
		fmt.Println("Encountered error when unmarshaling known good Video View:")
		fmt.Println(err)
		os.Exit(255)
	} else {
		fmt.Println("unmarshal-video-view-type-check ✅")
	}

}
