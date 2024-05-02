// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type PlaybackRestriction struct {
	// Unique identifier for the Playback Restriction. Max 255 characters.
	Id string `json:"id,omitempty"`
	// Time the Playback Restriction was created, defined as a Unix timestamp (seconds since epoch).
	CreatedAt string `json:"created_at,omitempty"`
	// Time the Playback Restriction was last updated, defined as a Unix timestamp (seconds since epoch).
	UpdatedAt string                       `json:"updated_at,omitempty"`
	Referrer  ReferrerDomainRestriction    `json:"referrer,omitempty"`
	UserAgent UserAgentRestrictionSettings `json:"user_agent,omitempty"`
}
