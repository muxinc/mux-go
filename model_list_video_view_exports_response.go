// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type ListVideoViewExportsResponse struct {
	Data          []ExportDate `json:"data,omitempty"`
	TotalRowCount int32        `json:"total_row_count,omitempty"`
	Timeframe     []int32      `json:"timeframe,omitempty"`
}