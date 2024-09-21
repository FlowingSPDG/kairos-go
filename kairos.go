package kairos

import (
	"context"
	"encoding/json"
	"net/http"

	"golang.org/x/xerrors"
)

type KairosRestClient interface {
	// AUX
	GetAuxByID(ctx context.Context, id string) (map[string]any, error)
	GetAuxByNumber(ctx context.Context, number int) (map[string]any, error)
	GetAuxs(ctx context.Context) ([]any, error)
	// PatchAux(ctx context.Context) error

	// Inputs
	GetInputByID(ctx context.Context, id string) (*Input, error)
	GetInputByNumber(ctx context.Context, number int) (*Input, error)
	GetInputs(ctx context.Context) ([]Input, error)

	// Macros
	GetMacro(ctx context.Context, id string) (*Macro, error)
	GetMacros(ctx context.Context) ([]Macro, error)
	PatchMacro(ctx context.Context, macroUuid string, state string) error

	// Multiviewers
	GetMultiviewer(ctx context.Context, mv string) (map[string]any, error)
	GetMultiviewers(ctx context.Context) ([]any, error)
	// PatchMultiviewer(ctx context.Context) error

	// Scenes
	GetScene(ctx context.Context, scene string) ([]any, error)
	GetScenes(ctx context.Context) ([]Scene, error)
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
}

func NewKairosRestClient(ip string, port string, user, password string) KairosRestClient {
	return &kairosRestClient{
		ip:       ip,
		port:     port,
		c:        &http.Client{},
		user:     user,
		password: password,
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
