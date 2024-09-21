package kairos

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMultiviewers(ctx context.Context) ([]any, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetMultiviewer(ctx context.Context, mv string) (map[string]any, error) {
	// input=id or number?

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers/%s", net.JoinHostPort(k.ip, k.port), mv)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) PatchMultiviewer(ctx context.Context) error {
	panic("TODO")
}
