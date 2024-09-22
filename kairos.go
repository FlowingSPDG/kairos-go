package kairos

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FlowingSPDG/kairos-go/internals/objects"
	"golang.org/x/xerrors"
)

// KairosRestClient is an interface to communicate with Panasonic Kairos.
// Currently version 1.4.0 is supported.
type KairosRestClient interface {
	// AUX
	GetAuxByID(ctx context.Context, id string) (*objects.AuxR, error)
	GetAuxByNumber(ctx context.Context, number int) (*objects.AuxR, error)
	GetAuxs(ctx context.Context) ([]*objects.AuxR, error)
	// PatchAux(ctx context.Context) error

	// Inputs
	GetInputByID(ctx context.Context, id string) (*objects.InputR, error)
	GetInputByNumber(ctx context.Context, number int) (*objects.InputR, error)
	GetInputs(ctx context.Context) ([]*objects.InputR, error)

	// Macros
	GetMacro(ctx context.Context, id string) (*objects.MacroR, error)
	GetMacros(ctx context.Context) ([]*objects.MacroR, error)
	PatchMacro(ctx context.Context, macroUuid string, state string) error

	// Multiviewers
	GetMultiviewerByID(ctx context.Context, mv string) (*objects.MultiviewerR, error)
	GetMultiviewerByNumber(ctx context.Context, mv int) (*objects.MultiviewerR, error)
	GetMultiviewers(ctx context.Context) ([]*objects.MultiviewerR, error)
	// PatchMultiviewer(ctx context.Context) error

	// Scenes
	GetScene(ctx context.Context, scene string) (*objects.SceneR, error)
	GetScenes(ctx context.Context) ([]*objects.SceneR, error)
	PatchScene(ctx context.Context, sceneUuid, layerUuid string, a, b *string, sources []string) error

	// Snapshot
	// PatchSnapshot(ctx context.Context) error
}

type kairosRestClient struct {
	c *http.Client

	user     string
	password string

	ep *Endpoints
}

func NewKairosRestClient(ip, port, user, password string) KairosRestClient {
	return &kairosRestClient{
		c: http.DefaultClient,

		user:     user,
		password: password,

		ep: NewEndpoints(ip, port),
	}
}

func (k *kairosRestClient) setGetHeaders(req *http.Request) {
	req.SetBasicAuth(k.user, k.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func (k *kairosRestClient) setPatchHeaders(req *http.Request) {
	req.SetBasicAuth(k.user, k.password)
	req.Header.Set("Content-Type", "application/json-patch+json")
	req.Header.Set("Accept", "application/json")
}

func (k *kairosRestClient) doGetRequest(req *http.Request, response any) error {
	k.setGetHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return xerrors.Errorf("Failed to decode response: %w", err)
	}

	return nil
}

func (k *kairosRestClient) doPatchRequest(req *http.Request, response any) error {
	k.setPatchHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return xerrors.Errorf("Failed to decode response: %w", err)
	}

	return nil
}

func doGET[T any](ctx context.Context, k *kairosRestClient, endpoint string, v *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	if err := k.doGetRequest(req, &v); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", v)

	return nil
}

func doPATCH[T any](ctx context.Context, k *kairosRestClient, endpoint string, body io.Reader, v *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, endpoint, body)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	if err := k.doPatchRequest(req, &v); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", v)

	return nil
}
