package kairos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

func (k *kairosRestClient) GetScenes(ctx context.Context) ([]*objects.SceneR, error) {
	ep := k.ep.Scenes()
	var payload []*objects.SceneR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return payload, nil
}

func (k *kairosRestClient) GetScene(ctx context.Context, scene string) (*objects.SceneR, error) {
	ep := k.ep.Scene(scene)
	var payload objects.SceneR
	if err := doGET(ctx, k, ep, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to get inputs: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) PatchScene(ctx context.Context, sceneUuid, layerUuid string, a, b *string) error {
	payload := objects.NewLayerWritePayload(a, b)
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return xerrors.Errorf("Failed to encode payload: %w", err)
	}

	// TODO: ep
	ep := path.Join(k.ep.Scene(sceneUuid), layerUuid)
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
