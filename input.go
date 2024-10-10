package kairos

import (
	"context"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetInputs(ctx context.Context) ([]*objects.InputR, error) {
	ep := k.ep.Inputs()
	var payload []*objects.InputR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

type InputIdentifier interface {
	~int | ~string
}

func getInput[T InputIdentifier](ctx context.Context, k *kairosRestClient, input T) (*objects.InputR, error) {
	ep := endPointInput(k.ep, input)
	var payload objects.InputR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) GetInputByID(ctx context.Context, id string) (*objects.InputR, error) {
	return getInput(ctx, k, id)
}

func (k *kairosRestClient) GetInputByNumber(ctx context.Context, number int) (*objects.InputR, error) {
	return getInput(ctx, k, number)
}
