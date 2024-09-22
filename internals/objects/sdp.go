package objects

import (
	"encoding/json"
	"strings"

	"github.com/pixelbender/go-sdp/sdp"
	"golang.org/x/xerrors"
)

var _ json.Unmarshaler = (*SDP)(nil)

type SDP struct {
	raw *sdp.Session
}

//nolint:ineffassign,staticcheck
func (s *SDP) UnmarshalJSON(b []byte) error {
	session, err := parseSDP(b)
	if err != nil {
		return xerrors.Errorf("Failed to Unmarshal sdp field: %w", err)
	}
	s.raw = session.raw
	return nil
}

func (s *SDP) Raw() *sdp.Session {
	return s.raw
}

func parseSDP(b []byte) (*SDP, error) {
	// Replace new line
	bs := strings.ReplaceAll(string(b), "\\n", "\n")
	// remove last line
	bs = strings.TrimSuffix(bs, "\n")
	// remove double quotes
	bs = strings.TrimLeft(bs, `"`)
	bs = strings.TrimRight(bs, `"`)
	s, err := sdp.ParseString(bs)
	if err != nil {
		return nil, xerrors.Errorf("failed to parse SDP: %w", err)
	}
	return &SDP{
		raw: s,
	}, nil
}
