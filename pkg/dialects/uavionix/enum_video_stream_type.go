//autogenerated:yes
//nolint:golint,misspell,govet,lll,dupl
package uavionix

import (
	"errors"
)

// Video stream types
type VIDEO_STREAM_TYPE int

const (
	// Stream is RTSP
	VIDEO_STREAM_TYPE_RTSP VIDEO_STREAM_TYPE = 0
	// Stream is RTP UDP (URI gives the port number)
	VIDEO_STREAM_TYPE_RTPUDP VIDEO_STREAM_TYPE = 1
	// Stream is MPEG on TCP
	VIDEO_STREAM_TYPE_TCP_MPEG VIDEO_STREAM_TYPE = 2
	// Stream is h.264 on MPEG TS (URI gives the port number)
	VIDEO_STREAM_TYPE_MPEG_TS_H264 VIDEO_STREAM_TYPE = 3
)

var labels_VIDEO_STREAM_TYPE = map[VIDEO_STREAM_TYPE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e VIDEO_STREAM_TYPE) MarshalText() ([]byte, error) {
	if l, ok := labels_VIDEO_STREAM_TYPE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_VIDEO_STREAM_TYPE = map[string]VIDEO_STREAM_TYPE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *VIDEO_STREAM_TYPE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_VIDEO_STREAM_TYPE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e VIDEO_STREAM_TYPE) String() string {
	if l, ok := labels_VIDEO_STREAM_TYPE[e]; ok {
		return l
	}
	return "invalid value"
}
