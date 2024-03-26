package kairos

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type KairosRestClient struct {
	ip       string
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

type getScenePayload []Scene

type Scene struct {
	Layers []Layer `json:"layers"`
	Name   string  `json:"name"`
	Tally  int     `json:"tally"`
	UUID   string  `json:"uuid"`
}

type Layer struct {
	Name    string   `json:"name"`
	SourceA string   `json:"sourceA"`
	SourceB string   `json:"sourceB"`
	Sources []string `json:"sources"`
	UUID    string   `json:"uuid"`
}

func (k *KairosRestClient) GetScene() (getScenePayload, error) {
	// TODO: context

	// エンドポイントの設定
	ep := "http://" + k.ip + ":1234/scenes"
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	// 認証情報の設定
	k.setHeaders(req)

	// リクエストの送信
	resp, err := k.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのデコード
	var payload getScenePayload
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
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

	k.setHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response patchSceneResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if response.Code != 200 {
		return errors.New(response.Text)
	}

	return nil

}

type Macro struct {
	Color string `json:"color"`
	Name  string `json:"name"`
	State any    `json:"state"`
	UUID  string `json:"uuid"`
}

type getMacrosPayload []Macro

func (k *KairosRestClient) GetMacros() (getMacrosPayload, error) {
	// TODO: context

	// エンドポイントの設定
	ep := "http://" + k.ip + ":1234/macros"
	req, err := http.NewRequest("GET", ep, nil)
	if err != nil {
		return nil, err
	}

	// 認証情報の設定
	k.setHeaders(req)

	// リクエストの送信
	resp, err := k.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのデコード
	var payload getMacrosPayload
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
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

	k.setHeaders(req)

	resp, err := k.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var response patchMacroResponsePayload
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	if response.Code != 200 {
		return errors.New(response.Text)
	}

	return nil

}

// TODO: generics
