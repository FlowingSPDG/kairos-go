package objects

// HOLY FUDGE WHAT THE HELL IS WRONG WITH KAIROS "OFFICIAL" API DOCUMENTS

type base struct {
	UUID string `json:"uuid"` // R // can be replaced with google.uuid?
}

type auxCommon struct {
	Source string `json:"source"`
}

type AuxR struct {
	base
	auxCommon
	Index   int      `json:"index"`
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
}

type AuxW struct {
	auxCommon
}

func NewAuxWriteRequest(src string) *AuxW {
	return &AuxW{
		auxCommon{
			Source: src,
		},
	}
}

type MacroR struct {
	base
	Color string `json:"color"` // TODO: parse color tp struct...
	Name  string `json:"name"`
}

type MacroW struct {
	State any `json:"state"` // "play" or "recall"?
}

type SnapshotR struct {
	base
	Name string `json:"name"`
}

type SnapshotW struct {
	State any `json:"state"` // "play" or "recall"?
}

type InputR struct {
	base
	Index int    `json:"index"`
	Name  string `json:"name"`
	Tally int    `json:"tally"`
}

type Action struct {
	base
	Name  string `json:"name"`
	State any    `json:"state"`
}

type SceneR struct {
	base
	Actions   []Action    `json:"actions"`
	Name      string      `json:"name"`
	Path      string      `json:"path"`
	Tally     int         `json:"tally"`
	Layers    []LayerR    `json:"layers"`
	Macros    []MacroR    `json:"macros"`    // R?
	Snapshots []SnapshotR `json:"snapshots"` // R?
}

type layerCommon struct {
	SourceA *string `json:"sourceA,omitempty"`
	SourceB *string `json:"sourceB,omitempty"`
}

type LayerR struct {
	base
	layerCommon
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
}

type LayerW struct {
	layerCommon
}

func NewLayerWritePayload(a, b *string) *LayerW {
	return &LayerW{
		layerCommon{
			SourceA: a,
			SourceB: b,
		},
	}
}

type MultiviewerR struct {
	base
	Index   int                  `json:"index"`
	Name    string               `json:"name"`
	Presets []MultiviewerPresetR `json:"presets"`
	SDP     *SDP                 `json:"sdp"`
}

type MultiviewerW struct {
	Preset any `json:"preset"`
}

type MultiviewerPresetR struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Usr  bool   `json:"usr"`
}

type PatchResponsePayload struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
