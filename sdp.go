package kairos

import (
	"github.com/pixelbender/go-sdp/sdp"
	"golang.org/x/xerrors"
)

// ep := fmt.Sprintf("http://%s/multiviewers/%v/sdp", net.JoinHostPort(k.ip, k.port), mv)

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
