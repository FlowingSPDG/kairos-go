package kairos

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type KairosRestClient struct {
	ip       string
	port     string
	c        *http.Client
	user     string
	password string
}

func NewKairosRestClient(ip, user, password string) *KairosRestClient {
	return &KairosRestClient{
		ip:       ip,
		c:        &http.Client{},
		user:     user,
		password: password,
	}
}

func (k *KairosRestClient) setHeaders(req *http.Request) {
	req.SetBasicAuth(k.user, k.password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}

func (k *KairosRestClient) doRequest(req *http.Request, response any) error {
	k.setHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

type getScenePayload []Scene

func (k *KairosRestClient) GetInputs() ([]Input, error) {
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/inputs", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	var payload []Input
	if err := k.doRequest(req, &payload); err != nil {
		return nil, err
	}

	return payload, nil
}

func (k *KairosRestClient) GetInput(input string) (*Input, error) {
	// input=id or number
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/inputs/%s", net.JoinHostPort(k.ip, k.port), input)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	var payload Input
	if err := k.doRequest(req, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (k *KairosRestClient) GetMacros() ([]Macro, error) {
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	var response []Macro
	if err := k.doRequest(req, &response); err != nil {
		return nil, err
	}
	return response, nil
}

func (k *KairosRestClient) GetMacro(id string) (*Macro, error) {
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/macros/%s", net.JoinHostPort(k.ip, k.port), id)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	response := Macro{}
	if err := k.doRequest(req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (k *KairosRestClient) GetAuxs() error {
	// input=id or number
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return err
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return err
	}
	fmt.Printf("Payload: %+v\n", payload)

	return nil
}

func (k *KairosRestClient) GetAux(aux string) error {
	// input=id or number
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/aux/%s", net.JoinHostPort(k.ip, k.port), aux)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return err
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return err
	}
	fmt.Printf("Payload: %+v\n", payload)

	return nil
}

func (k *KairosRestClient) GetMultiviewers() error {
	// input=id or number
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return err
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return err
	}
	fmt.Printf("Payload: %+v\n", payload)

	return nil
}

func (k *KairosRestClient) GetMultiviewer(mv string) error {
	// input=id or number
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/multiviewers/%s", net.JoinHostPort(k.ip, k.port), mv)
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return err
	}

	var payload map[string]any
	if err := k.doRequest(req, &payload); err != nil {
		return err
	}
	fmt.Printf("Payload: %+v\n", payload)

	return nil
}

func (k *KairosRestClient) GetScenes() ([]Scene, error) {
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	var payload []Scene
	if err := k.doRequest(req, &payload); err != nil {
		return nil, err
	}
	return payload, nil
}

func (k *KairosRestClient) GetScene(scene string) (*Scene, error) {
	// TODO: context

	// エンドポイントの設定
	ep := fmt.Sprintf("http://%s/scenes", net.JoinHostPort(k.ip, k.port))
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	var payload Scene
	if err := k.doRequest(req, &payload); err != nil {
		return nil, err
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

func (k *KairosRestClient) PatchScene(sceneUuid, layerUuid, a, b string, sources []string) error {
	payload := patchSceneRequestPayload{
		Name:    "Background",
		SourceA: a,
		SourceB: b,
		Sources: sources,
		UUID:    layerUuid,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return err
	}

	ep := "http://" + k.ip + ":1234/scenes/" + sceneUuid + "/" + layerUuid
	req, err := http.NewRequest("PATCH", ep, &buf)
	if err != nil {
		return err
	}

	var response patchSceneResponsePayload
	if err := k.doRequest(req, &response); err != nil {
		return err
	}
	if response.Code != 200 {
		return errors.New(response.Text)
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

func (k *KairosRestClient) PatchMacro(macroUuid, state string) error {
	payload := patchMacroRequestPayload{
		State: state,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(payload); err != nil {
		return err
	}

	ep := "http://" + k.ip + ":1234/macros/" + macroUuid
	req, err := http.NewRequest("PATCH", ep, &buf)
	if err != nil {
		return err
	}

	var response patchMacroResponsePayload
	if err := k.doRequest(req, &response); err != nil {
		return err
	}
	if response.Code != 200 {
		return errors.New(response.Text)
	}
	return nil
}

func (k *KairosRestClient) PatchSnapshot() error {
	panic("TODO")
}

func (k *KairosRestClient) PatchAux() error {
	panic("TODO")
}

func (k *KairosRestClient) PatchMultiviewer() error {
	panic("TODO")
}
