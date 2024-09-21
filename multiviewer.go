package kairos

import (
	"context"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMultiviewers(ctx context.Context) ([]*Multiviewer, error) {
	// エンドポイントの設定
	ep := k.ep.Multiviewers()
	var payload []*Multiviewer
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

type MultiviewerIdentifier interface {
	~int | ~string
}

func getMultiviewer[T MultiviewerIdentifier](ctx context.Context, k *kairosRestClient, mv T) (*Multiviewer, error) {
	// エンドポイントの設定
	ep := endPointMultiviewers(k.ep, mv)
	var payload Multiviewer
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
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
