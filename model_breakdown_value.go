// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type BreakdownValue struct {
	Views            int64   `json:"views,omitempty"`
	Value            float64 `json:"value,omitempty"`
	TotalWatchTime   int64   `json:"total_watch_time,omitempty"`
	TotalPlayingTime int64   `json:"total_playing_time,omitempty"`
	NegativeImpact   int32   `json:"negative_impact,omitempty"`
	Field            string  `json:"field,omitempty"`
}
