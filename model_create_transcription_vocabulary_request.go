// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateTranscriptionVocabularyRequest struct {
	// The user-supplied name of the Transcription Vocabulary.
	Name string `json:"name,omitempty"`
	// Phrases, individual words, or proper names to include in the Transcription Vocabulary. When the Transcription Vocabulary is attached to a live stream's `generated_subtitles`, the probability of successful speech recognition for these words or phrases is boosted.
	Phrases []string `json:"phrases"`
	// Arbitrary user-supplied metadata set for the Transcription Vocabulary. Max 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
}
