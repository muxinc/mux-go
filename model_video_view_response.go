// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type VideoViewResponse struct {
	Data          VideoView `json:"data,omitempty"`
	Timeframe     []int64   `json:"timeframe,omitempty"`
	TotalRowCount int64     `json:"total_row_count,omitempty"`
}
