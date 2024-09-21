package kairos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMacros(ctx context.Context) ([]*Macro, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []*Macro
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetMacro(ctx context.Context, id string) (*Macro, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros/%s", net.JoinHostPort(k.ip, k.port), id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	payload := Macro{}
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return &payload, nil
}

type patchMacroRequestPayload struct {
	State string `json:"state"` // play only??
}

type patchMacroResponsePayload struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (k *kairosRestClient) PatchMacro(ctx context.Context, macroUuid, state string) error {
	payload := patchMacroRequestPayload{
		State: state,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return xerrors.Errorf("Failed to encode payload: %w", err)
	}

	ep := fmt.Sprintf("http://%s/macros/%s", net.JoinHostPort(k.ip, k.port), macroUuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, ep, &buf)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	var response patchMacroResponsePayload
	if err := k.doRequest(req, &response); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	if response.Code != 200 {
		return xerrors.New(response.Text)
	}
	return nil
}
