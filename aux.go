package kairos

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetAuxs(ctx context.Context) ([]*Aux, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []*Aux
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

type AuxIdentifier interface {
	~int | ~string
}

func getAux[T AuxIdentifier](ctx context.Context, k *kairosRestClient, aux T) (*Aux, error) {
	// input=id or number

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux/%v", net.JoinHostPort(k.ip, k.port), aux)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload *Aux
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetAuxByID(ctx context.Context, id string) (*Aux, error) {
	return getAux(ctx, k, id)
}

func (k *kairosRestClient) GetAuxByNumber(ctx context.Context, number int) (*Aux, error) {
	return getAux(ctx, k, number)
}

func (k *kairosRestClient) PatchAux(ctx context.Context) error {
	panic("TODO")
}
