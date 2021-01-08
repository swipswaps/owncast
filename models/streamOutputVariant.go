package models

import "encoding/json"

// StreamOutputVariant defines the output specifics of a single HLS stream variant.
type StreamOutputVariant struct {
	// Enable passthrough to copy the video and/or audio directly from the
	// incoming stream and disable any transcoding.  It will ignore any of
	// the below settings.
	IsVideoPassthrough bool `yaml:"videoPassthrough" json:"videoPassthrough"`
	IsAudioPassthrough bool `yaml:"audioPassthrough" json:"audioPassthrough"`

	VideoBitrate int `yaml:"videoBitrate" json:"videoBitrate"`
	AudioBitrate int `yaml:"audioBitrate" json:"audioBitrate"`

	// Set only one of these in order to keep your current aspect ratio.
	// Or set neither to not scale the video.
	ScaledWidth  int `yaml:"scaledWidth" json:"scaledWidth,omitempty"`
	ScaledHeight int `yaml:"scaledHeight" json:"scaledHeight,omitempty"`

	Framerate     int    `yaml:"framerate" json:"framerate"`
	EncoderPreset string `yaml:"encoderPreset" json:"encoderPreset"`
}

// GetFramerate returns the framerate or default.
func (q *StreamOutputVariant) GetFramerate() int {
	if q.IsVideoPassthrough {
		return 0
	}

	if q.Framerate > 0 {
		return q.Framerate
	}

	return 24
}

// GetEncoderPreset returns the preset or default.
func (q *StreamOutputVariant) GetEncoderPreset() string {
	if q.IsVideoPassthrough {
		return ""
	}

	if q.EncoderPreset != "" {
		return q.EncoderPreset
	}

	return "veryfast"
}

func (q *StreamOutputVariant) GetIsAudioPassthrough() bool {
	if q.IsAudioPassthrough {
		return true
	}

	if q.AudioBitrate == 0 {
		return true
	}

	return false
}

// MarshalJSON is a custom JSON marshal function for video stream qualities.
func (q *StreamOutputVariant) MarshalJSON() ([]byte, error) {
	type Alias StreamOutputVariant
	return json.Marshal(&struct {
		Framerate int `json:"framerate"`
		*Alias
	}{
		Framerate: q.GetFramerate(),
		Alias:     (*Alias)(q),
	})
}
