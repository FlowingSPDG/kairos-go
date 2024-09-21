package kairos

import (
	"context"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetInputs(ctx context.Context) ([]*Input, error) {
	ep := k.ep.Inputs()
	var payload []*Input
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

type InputIdentifier interface {
	~int | ~string
}

func getInput[T InputIdentifier](ctx context.Context, k *kairosRestClient, input T) (*Input, error) {
	ep := endPointInput(k.ep, input)
	var payload Input
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) GetInputByID(ctx context.Context, id string) (*Input, error) {
	return getInput(ctx, k, id)
}

func (k *kairosRestClient) GetInputByNumber(ctx context.Context, number int) (*Input, error) {
	return getInput(ctx, k, number)
}
