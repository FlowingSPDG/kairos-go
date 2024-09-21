package kairos

// HOLY FUDGE WHAT THE HELL IS WRONG WITH KAIROS "OFFICIAL" API DOCUMENTS

type Aux struct {
	UUID    string   `json:"uuid"`    // R // can be replaced with google.uuid?
	Index   int      `json:"index"`   // R
	Name    string   `json:"name"`    // R
	Source  string   `json:"source"`  // R/W
	Sources []string `json:"sources"` // W
}

type Macro struct {
	UUID  string `json:"uuid"`  // R
	Color string `json:"color"` // R // TODO: parse color tp struct...
	Name  string `json:"name"`  // R
	State any    `json:"state"` // W // null if read // "play" or "recall"?
}

type Snapshot struct {
	UUID  string `json:"uuid"` // R
	Name  string `json:"name"` // R
	State any    // W // null if read
}

type Input struct {
	UUID  string `json:"uuid"`  // R
	Index int    `json:"index"` // R
	Name  string `json:"name"`  // R
	Tally int    `json:"tally"` // R
}

type Scene struct {
	UUID   string  `json:"uuid"`   // R
	Name   string  `json:"name"`   // R
	Tally  int     `json:"tally"`  // R
	Layers []Layer `json:"layers"` // R

	// following fields are not documented
	Macros    []Macro    `json:"macros"`    // R?
	Snapshots []Snapshot `json:"snapshots"` // R?
}

type Layer struct {
	UUID    string   `json:"uuid"`              // R
	Name    string   `json:"name"`              // R
	SourceA *string  `json:"sourceA,omitempty"` // R/W // optional
	SourceB *string  `json:"sourceB,omitempty"` // R/W // optional
	Sources []string `json:"sources"`           // R
}

type Multiviewer struct {
	Index   int                 // R
	Name    string              // R
	Preset  any                 // W // null if read
	Presets []MultiviewerPreset // R
	SDP     string              // R
}

type MultiviewerPreset struct {
	ID   int    `json:"id"`   // R
	Name string `json:"name"` // R
	Usr  bool   `json:"usr"`  // R
}

// TODO: SDP. Content-Type: "application/sdp"
