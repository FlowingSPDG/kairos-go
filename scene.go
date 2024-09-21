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

func (k *kairosRestClient) GetScenes(ctx context.Context) ([]*Scene, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	var payload []*Scene
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

func (k *kairosRestClient) GetScene(ctx context.Context, scene string) (*Scene, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	var payload Scene
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

type patchSceneRequestPayload struct {
	Name    string   `json:"name"`
	SourceA string   `json:"sourceA"`
	SourceB string   `json:"sourceB"`
	Sources []string `json:"sources"`
	UUID    string   `json:"uuid"` // Layer uuid
}

type patchSceneResponsePayload struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (k *kairosRestClient) PatchScene(ctx context.Context, sceneUuid, layerUuid, a, b string, sources []string) error {
	payload := patchSceneRequestPayload{
		Name:    "Background",
		SourceA: a,
		SourceB: b,
		Sources: sources,
		UUID:    layerUuid,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return xerrors.Errorf("Failed to encode payload: %w", err)
	}

	ep := fmt.Sprintf("http://%s/scenes/%s/%s", net.JoinHostPort(k.ip, k.port), sceneUuid, layerUuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, ep, &buf)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	var response patchSceneResponsePayload
	if err := k.doRequest(req, &response); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	if response.Code != 200 {
		return xerrors.New(response.Text)
	}
	return nil
}
