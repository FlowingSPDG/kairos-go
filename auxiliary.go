package kairos

import (
	"context"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetAuxs(ctx context.Context) ([]*objects.AuxR, error) {
	// エンドポイントの設定
	ep := k.ep.Auxs()
	var payload []*objects.AuxR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

func getAux[T AuxIdentifier](ctx context.Context, k *kairosRestClient, aux T) (*objects.AuxR, error) {
	ep := endPointAux(k.ep, aux)
	var payload objects.AuxR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) GetAuxByID(ctx context.Context, id string) (*objects.AuxR, error) {
	return getAux(ctx, k, id)
}

func (k *kairosRestClient) GetAuxByNumber(ctx context.Context, number int) (*objects.AuxR, error) {
	return getAux(ctx, k, number)
}

func (k *kairosRestClient) PatchAux(ctx context.Context) error {
	panic("TODO")
}
