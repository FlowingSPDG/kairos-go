package kairos

// HOLY FUDGE WHAT THE HELL IS WRONG WITH KAIROS "OFFICIAL" API DOCUMENTS

type Aux struct {
	UUID    string   `json:"uuid"`
	Index   int      `json:"index"`
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	Sources []string `json:"sources"`
}

type Macro struct {
	UUID  string `json:"uuid"`
	Color string `json:"color"`
	Name  string `json:"name"`
	State any    `json:"state"`
}

type Snapshot struct {
	UUID  string `json:"uuid"`
	State any
}

type Input struct {
	UUID  string `json:"uuid"`
	Index int    `json:"index"`
	Name  string `json:"name"`
	Tally int    `json:"tally"`
}

type Scene struct {
	UUID   string  `json:"uuid"`
	Layers []Layer `json:"layers"`
	Name   string  `json:"name"`
	Tally  int     `json:"tally"`
}

type Layer struct {
	UUID    string   `json:"uuid"`
	Name    string   `json:"name"`
	SourceA string   `json:"sourceA"`
	SourceB string   `json:"sourceB"`
	Sources []string `json:"sources"`
}

type Multiviewer struct {
	Index   int
	Name    string
	Preset  any
	Presets []MultiviewerPreset
	SDP     string
}

type MultiviewerPreset struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Usr  bool   `json:"usr"`
}

// TODO: SDP. Content-Type: "application/sdp"
