package kairos

import (
	"context"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMultiviewers(ctx context.Context) ([]*objects.MultiviewerR, error) {
	ep := k.ep.Multiviewers()
	var payload []*objects.MultiviewerR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

type MultiviewerIdentifier interface {
	~int | ~string
}

func getMultiviewer[T MultiviewerIdentifier](ctx context.Context, k *kairosRestClient, mv T) (*objects.MultiviewerR, error) {
	ep := endPointMultiviewers(k.ep, mv)
	var payload objects.MultiviewerR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) GetMultiviewerByID(ctx context.Context, mv string) (*objects.MultiviewerR, error) {
	return getMultiviewer(ctx, k, mv)
}

func (k *kairosRestClient) GetMultiviewerByNumber(ctx context.Context, mv int) (*objects.MultiviewerR, error) {
	return getMultiviewer(ctx, k, mv)
}

func (k *kairosRestClient) PatchMultiviewer(ctx context.Context) error {
	panic("TODO")
}
