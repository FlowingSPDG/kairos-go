package kairos

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMultiviewers(ctx context.Context) ([]*Multiviewer, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []*Multiviewer
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

type MultiviewerIdentifier interface {
	~int | ~string
}

func getMultiviewer[T MultiviewerIdentifier](ctx context.Context, k *kairosRestClient, mv T) (*Multiviewer, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers/%v", net.JoinHostPort(k.ip, k.port), mv)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload *Multiviewer
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetMultiviewerByID(ctx context.Context, mv string) (*Multiviewer, error) {
	return getMultiviewer(ctx, k, mv)
}

func (k *kairosRestClient) GetMultiviewerByNumber(ctx context.Context, mv int) (*Multiviewer, error) {
	return getMultiviewer(ctx, k, mv)
}

func (k *kairosRestClient) PatchMultiviewer(ctx context.Context) error {
	panic("TODO")
}
