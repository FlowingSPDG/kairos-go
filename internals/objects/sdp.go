package objects

import (
	"encoding/json"

	"github.com/pixelbender/go-sdp/sdp"
	"golang.org/x/xerrors"
)

var _ json.Unmarshaler = (*SDP)(nil)

type SDP struct {
	raw *sdp.Session
}

func (s *SDP) UnmarshalJSON(b []byte) (err error) {
	s, err = parseSDP(b)
	if err != nil {
		return xerrors.Errorf("Failed to Unmarshal sdp field: %w", err)
	}
	return nil
}

func (s *SDP) Raw() *sdp.Session {
	return s.raw
}

func parseSDP(b []byte) (*SDP, error) {
	s, err := sdp.Parse(b)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse SDP: %w", err)
	}
	return &SDP{
		raw: s,
	}, nil
}
func parseSDPStr(str string) (*SDP, error) {
	s, err := sdp.ParseString(str)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse SDP: %w", err)
	}
	return &SDP{
		raw: s,
	}, nil
}
