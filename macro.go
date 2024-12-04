package kairos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetMacros(ctx context.Context) ([]*objects.MacroR, error) {
	ep := k.ep.Macros()
	var payload []*objects.MacroR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

func (k *kairosRestClient) GetMacro(ctx context.Context, id string) (*objects.MacroR, error) {
	ep := k.ep.Macro(id)
	var payload objects.MacroR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) PatchMacro(ctx context.Context, macroUuid, state string) error {
	payload := objects.MacroW{
		State: state,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return xerrors.Errorf("Failed to encode payload: %w", err)
	}

	ep := k.ep.Macro(macroUuid)
	var response objects.PatchResponsePayload
	if err := doPATCH(ctx, k, ep, &buf, &response); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	if response.Code != 200 {
		return xerrors.New(response.Text)
	}
	return nil
}
