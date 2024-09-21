package kairos

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

// KairosRestClient is an interface to communicate with Panasonic Kairos.
// Currently version 1.4.0 is supported.
type KairosRestClient interface {
	// AUX
	GetAuxByID(ctx context.Context, id string) (*Aux, error)
	GetAuxByNumber(ctx context.Context, number int) (*Aux, error)
	GetAuxs(ctx context.Context) ([]*Aux, error)
	// PatchAux(ctx context.Context) error

	// Inputs
	GetInputByID(ctx context.Context, id string) (*Input, error)
	GetInputByNumber(ctx context.Context, number int) (*Input, error)
	GetInputs(ctx context.Context) ([]*Input, error)

	// Macros
	GetMacro(ctx context.Context, id string) (*Macro, error)
	GetMacros(ctx context.Context) ([]*Macro, error)
	PatchMacro(ctx context.Context, macroUuid string, state string) error

	// Multiviewers
	GetMultiviewerByID(ctx context.Context, mv string) (*Multiviewer, error)
	GetMultiviewerByNumber(ctx context.Context, mv int) (*Multiviewer, error)
	GetMultiviewers(ctx context.Context) ([]*Multiviewer, error)
	// PatchMultiviewer(ctx context.Context) error

	// Scenes
	GetScene(ctx context.Context, scene string) (*Scene, error)
	GetScenes(ctx context.Context) ([]*Scene, error)
	PatchScene(ctx context.Context, sceneUuid string, layerUuid string, a string, b string, sources []string) error

	// Snapshot
	// PatchSnapshot(ctx context.Context) error
}

type kairosRestClient struct {
	ip       string
	port     string
	c        *http.Client
	user     string
	password string

	ep *Endpoints
}

func NewKairosRestClient(ip string, port string, user, password string) KairosRestClient {
	return &kairosRestClient{
		ip:       ip,
		port:     port,
		c:        &http.Client{},
		user:     user,
		password: password,

		ep: NewEndpoints(ip, port),
	}
}

func (k *kairosRestClient) setHeaders(req *http.Request) {
	req.SetBasicAuth(k.user, k.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func (k *kairosRestClient) doRequest(req *http.Request, response any) error {
	k.setHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

// TODO: エンドポイントを解決する関数を作成する

func doGET[T any](ctx context.Context, k *kairosRestClient, endpoint string, v *T) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	if err := k.doRequest(req, &v); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", v)

	return nil
}

// TODO: doPatch
