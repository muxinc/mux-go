// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type MonitoringHistogramTimeseriesDatapoint struct {
	Timestamp     string                                      `json:"timestamp,omitempty"`
	Sum           int64                                       `json:"sum,omitempty"`
	P95           float64                                     `json:"p95,omitempty"`
	Median        float64                                     `json:"median,omitempty"`
	MaxPercentage float64                                     `json:"max_percentage,omitempty"`
	BucketValues  []MonitoringHistogramTimeseriesBucketValues `json:"bucket_values,omitempty"`
	Average       float64                                     `json:"average,omitempty"`
}
