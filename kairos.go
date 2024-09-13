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

type KairosRestClient interface {
	// AUX
	GetAuxByID(ctx context.Context, id string) (map[string]any, error)
	GetAuxByNumber(ctx context.Context, number int) (map[string]any, error)
	GetAuxs(ctx context.Context) (map[string]any, error)
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
	GetMultiviewers(ctx context.Context) (map[string]any, error)
	// PatchMultiviewer(ctx context.Context) error

	// Scenes
	GetScene(ctx context.Context, scene string) (*Scene, error)
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

func NewKairosRestClient(ip, user, password string) KairosRestClient {
	return &kairosRestClient{
		ip:       ip,
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

func (k *kairosRestClient) GetInputs(ctx context.Context) ([]Input, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/inputs", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []Input
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}

	return payload, nil
}

type InputIdentifier interface {
	~int | ~string
}

func getInput[T InputIdentifier](ctx context.Context, k *kairosRestClient, input T) (*Input, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/inputs/%v", net.JoinHostPort(k.ip, k.port), input)
	req, err := http.NewRequest(http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload Input
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}

	return &payload, nil
}

func (k *kairosRestClient) GetInputByID(ctx context.Context, id string) (*Input, error) {
	return getInput(ctx, k, id)
}

func (k *kairosRestClient) GetInputByNumber(ctx context.Context, number int) (*Input, error) {
	return getInput(ctx, k, number)
}

func (k *kairosRestClient) GetMacros(ctx context.Context) ([]Macro, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var response []Macro
	if err := k.doRequest(req, &response); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	return response, nil
}

func (k *kairosRestClient) GetMacro(ctx context.Context, id string) (*Macro, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros/%s", net.JoinHostPort(k.ip, k.port), id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	response := Macro{}
	if err := k.doRequest(req, &response); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	return &response, nil
}

func (k *kairosRestClient) GetAuxs(ctx context.Context) (map[string]any, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	// TODO: type
	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

type AuxIdentifier interface {
	~int | ~string
}

func getAux[T AuxIdentifier](ctx context.Context, k *kairosRestClient, aux T) (map[string]any, error) {
	// input=id or number

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux/%v", net.JoinHostPort(k.ip, k.port), aux)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	// TODO: type
	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetAuxByID(ctx context.Context, id string) (map[string]any, error) {
	return getAux(ctx, k, id)
}

func (k *kairosRestClient) GetAuxByNumber(ctx context.Context, number int) (map[string]any, error) {
	return getAux(ctx, k, number)
}

func (k *kairosRestClient) GetMultiviewers(ctx context.Context) (map[string]any, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetMultiviewer(ctx context.Context, mv string) (map[string]any, error) {
	// input=id or number?

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers/%s", net.JoinHostPort(k.ip, k.port), mv)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	fmt.Printf("Payload: %+v\n", payload)

	return payload, nil
}

func (k *kairosRestClient) GetScenes(ctx context.Context) ([]Scene, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload []Scene
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
	}
	return payload, nil
}

func (k *kairosRestClient) GetScene(ctx context.Context, scene string) (*Scene, error) {
	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, xerrors.Errorf("Failed to create request: %w", err)
	}

	var payload Scene
	if err := k.doRequest(req, &payload); err != nil {
		return nil, xerrors.Errorf("Failed to do request: %w", err)
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
	if response.Code != 200 {
		return xerrors.New(response.Text)
	}
	return nil
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

	ep := "http://" + k.ip + ":1234/macros/" + macroUuid
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, ep, &buf)
	if err != nil {
		return xerrors.Errorf("Failed to create request: %w", err)
	}

	var response patchMacroResponsePayload
	if err := k.doRequest(req, &response); err != nil {
		return xerrors.Errorf("Failed to do request: %w", err)
	}
	if response.Code != 200 {
		return xerrors.New(response.Text)
	}
	return nil
}

func (k *kairosRestClient) PatchSnapshot(ctx context.Context) error {
	panic("TODO")
}

func (k *kairosRestClient) PatchAux(ctx context.Context) error {
	panic("TODO")
}

func (k *kairosRestClient) PatchMultiviewer(ctx context.Context) error {
	panic("TODO")
}
