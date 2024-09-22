package objects

import (
	"encoding/json"
	"testing"

	"github.com/pixelbender/go-sdp/sdp"
	"github.com/stretchr/testify/assert"
)

type testCase[TIn, TOut any] struct {
	Name     string
	Input    TIn
	Expected TOut
	Error    error
}

func testUnmarshalJSON[T any](t *testing.T, incoming []byte, expected T) {
	actual := *new(T)
	if err := json.Unmarshal(incoming, &actual); err != nil {
		t.Fatalf("Failed to Unmarshal JSON:%v", err)
	}
	assert.Equal(t, expected, actual)
}

func assertUnmarshalJSON[TIn ~string, TOut any](t *testing.T, tcs []testCase[TIn, TOut]) {
	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			testUnmarshalJSON(t, []byte(tc.Input), tc.Expected)
		})
	}
}

func TestUnmarshalAuxR(t *testing.T) {
	t.Parallel()
	tcs := []testCase[string, AuxR]{
		{
			Name:     "Empty Input",
			Input:    `{}`,
			Expected: AuxR{},
			Error:    nil,
		},
		{
			Name: "Example on API document",
			Input: `{
"index":0,
"name": "IP-AUX1",
"source": "Main",
"sources": [
"Black",
"White"
],
"uuid": "04c03dbe-aa4e-5bdd-a98d-942f5c19ecbd"
}`,
			Expected: AuxR{
				Index: 0,
				Name:  "IP-AUX1",
				auxCommon: auxCommon{
					Source: "Main",
				},
				Sources: []string{
					"Black",
					"White",
				},
				base: base{
					UUID: "04c03dbe-aa4e-5bdd-a98d-942f5c19ecbd",
				},
			},
			Error: nil,
		},
	}

	assertUnmarshalJSON(t, tcs)
}

func TestUnmarshalAuxRs(t *testing.T) {
	t.Parallel()
	tcs := []testCase[string, []*AuxR]{
		{
			Name:     "Empty Input",
			Input:    `[]`,
			Expected: []*AuxR{},
			Error:    nil,
		},
		{
			Name: "Example on API document",
			Input: `
[
{
"index":0,
"name": "IP-AUX1 ",
"source": "IP1",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", "IP1"],
"uuid": "5d152781-63c6-5dcc-ad35-e8b5be025867"
},
{
"index":1,
"name": "IP-AUX2",
"source": "Main",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "952e8432-60aa-57c2-af7a-c7666c8fbf55"
},
{
"index":2,
"name": "IP-AUX3",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "32917296-436a-5edd-bf7c-7d9197bba74d"
},
{
"index":3,
"name": "IP-AUX4",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "313e020a-5f25-5a9b-83b7-3a3b0405ee76"
},
{
"index":4,
"name": "IP-AUX5",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "5ccee3f0-09ef-5f19-add8-16c1bc9ed5a3"
},
{
"index":5,
"name": "IP-AUX6",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "f1dfae31-9423-5596-a5ad-2e2c537f91cd" 
},
{
"index":40,
"name": "SDI-AUX1",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "26948dfc-06b4-5eca-b41e-a5dadce34856"
},
{
"index":41,
"name": "SDI-AUX2",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "bc1ea12a-64f3-5ae6-9c74-38abc90de4f0"
},
{
"index":72,
"name": "NDI-AUX1",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "afd163d7-b99b-528a-910e-9413304b3c94"
},
{
"index":73,
"name": "NDI-AUX2",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "cf2a841a-8bc9-58b3-b43b-c017da6eb4e7"
},
{
"index":74,
"name": "Stream-AUX1",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "d812a49f-1038-5bd9-a50f-658ecded4784"
},
{
"index":75,
"name": "Stream-AUX2",
"source": "Black",
"sources":["Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"],
"uuid": "d5127876-8b42-57e3-be03-2ef3803ceb09"
}
]`,
			Expected: []*AuxR{
				{
					Index:     0,
					Name:      "IP-AUX1 ", // I have no idea why the F they add this damn space.
					auxCommon: auxCommon{Source: "IP1"},
					// This is neither.
					Sources: []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", "IP1"},
					base:    base{UUID: "5d152781-63c6-5dcc-ad35-e8b5be025867"},
				},
				{
					Index:     1,
					Name:      "IP-AUX2",
					auxCommon: auxCommon{Source: "Main"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "952e8432-60aa-57c2-af7a-c7666c8fbf55"},
				},
				{
					Index:     2,
					Name:      "IP-AUX3",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "32917296-436a-5edd-bf7c-7d9197bba74d"},
				},
				{
					Index:     3,
					Name:      "IP-AUX4",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "313e020a-5f25-5a9b-83b7-3a3b0405ee76"},
				},
				{
					Index:     4,
					Name:      "IP-AUX5",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "5ccee3f0-09ef-5f19-add8-16c1bc9ed5a3"},
				},
				{
					Index:     5,
					Name:      "IP-AUX6",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "f1dfae31-9423-5596-a5ad-2e2c537f91cd"},
				},
				{
					Index:     40,
					Name:      "SDI-AUX1",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "26948dfc-06b4-5eca-b41e-a5dadce34856"},
				},
				{
					Index:     41,
					Name:      "SDI-AUX2",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "bc1ea12a-64f3-5ae6-9c74-38abc90de4f0"},
				},
				{
					Index:     72,
					Name:      "NDI-AUX1",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "afd163d7-b99b-528a-910e-9413304b3c94"},
				},
				{
					Index:     73,
					Name:      "NDI-AUX2",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "cf2a841a-8bc9-58b3-b43b-c017da6eb4e7"},
				},
				{
					Index:     74,
					Name:      "Stream-AUX1",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "d812a49f-1038-5bd9-a50f-658ecded4784"},
				},
				{
					Index:     75,
					Name:      "Stream-AUX2",
					auxCommon: auxCommon{Source: "Black"},
					Sources:   []string{"Black", "White", "ColA", "ColB", "ColC", "FxIN1", "FxIN1", " IP1"},
					base:      base{UUID: "d5127876-8b42-57e3-be03-2ef3803ceb09"},
				},
			},
			Error: nil,
		},
	}

	assertUnmarshalJSON(t, tcs)
}

func TestUnmarshalMultiviewerRs(t *testing.T) {
	t.Parallel()
	tcs := []testCase[string, []*MultiviewerR]{
		{
			Name:     "Empty Input",
			Input:    `[]`,
			Expected: []*MultiviewerR{},
			Error:    nil,
		},
		{
			Name: "Example on API document",
			Input: `[
 {
 "index": 0,
 "name": "Multiviewer1",
 "preset": null,
 "presets": [
 {
 "id": 0,
 "name": "Full",
 "usr": false
 },
 {
 "id": 1,
 "name": "Quad Split",
 "usr": false
 },
 {
 "id": 2,
 "name": "10 Split A",
 "usr": false
 },
 {
 "id": 3,
 "name": "10 Split B",
 "usr": false
 },
 {
 "id": 4,
 "name": "10 Split C",
 "usr": false
 },
 {
 "id": 5,
 "name": "10 Split D",
 "usr": false
 },
 {
 "id": 6,
 "name": "9 Split",
 "usr": false
 },
 {
 "id": 7,
 "name": "16 Split",
 "usr": false
 },
 {
 "id": 8,
 "name": "25 Split",
 "usr": false
 },
 {
 "id": 9,
 "name": "36 Split",
 "usr": false
 }
 ],
 "sdp": "v=0\ns=Multiviewer 0\no=- 1 1 IN IP4 192.168.10.42\nc=IN IP4 239.168.10.42\nt=0 0\nm=video 50000 RTP/AVP 96\na=rtpmap:96 H264/90000\n", 
 "uuid": "fd839e94-a9e9-570b-a5aa-bf99575f364f"
 },
 {
 "index":1,
 "name": "Multiviewer2",
 "preset": null,
 "presets": [ { "id": 0, "name": "Full", "usr": false }, { "id": 1, "name": "Quad Split", "usr": false }, { "id": 2, "name": "10 Split A", "usr":false }, { "id": 3, "name": "10 Split B", "usr": false }, { "id": 4, "name": "10 Split C", "usr": false }, { "id": 5, "name": "10 Split D", "usr": false },{ "id": 6, "name": "9 Split", "usr":false }, { "id": 7, "name": "16 Split", "usr": false }, { "id": 8, "name": "25 Split", "usr": false }, { "id": 9, "name": "36Split", "usr": false } ],
 "sdp": "v=0\ns=Multiviewer 1\no=- 1 1 IN IP4 192.168.10.42\nc=IN IP4 239.168.10.42\nt=0 0\nm=video 51000 RTP/AVP 96\na=rtpmap:96 H264/90000\n",
 "uuid": "ef884504-d688-59c7-a2a0-2b7d69e70e5b"
 }
 ]
`,
			Expected: []*MultiviewerR{
				{
					base: base{
						UUID: "fd839e94-a9e9-570b-a5aa-bf99575f364f",
					},
					Index: 0,
					Name:  "Multiviewer1",
					Presets: []MultiviewerPresetR{
						{
							ID:   0,
							Name: "Full",
							Usr:  false,
						},
						{
							ID:   1,
							Name: "Quad Split",
							Usr:  false,
						},
						{
							ID:   2,
							Name: "10 Split A",
							Usr:  false,
						},
						{
							ID:   3,
							Name: "10 Split B",
							Usr:  false,
						},
						{
							ID:   4,
							Name: "10 Split C",
							Usr:  false,
						},
						{
							ID:   5,
							Name: "10 Split D",
							Usr:  false,
						},
						{
							ID:   6,
							Name: "9 Split",
							Usr:  false,
						},
						{
							ID:   7,
							Name: "16 Split",
							Usr:  false,
						},
						{
							ID:   8,
							Name: "25 Split",
							Usr:  false,
						},
						{
							ID:   9,
							Name: "36 Split",
							Usr:  false,
						},
					},
					SDP: &SDP{
						raw: nil,
					},
				},
				{
					base: base{
						UUID: "ef884504-d688-59c7-a2a0-2b7d69e70e5b",
					},
					Index: 1,
					Name:  "Multiviewer2",
					Presets: []MultiviewerPresetR{
						{
							ID:   0,
							Name: "Full",
							Usr:  false,
						},
						{
							ID:   1,
							Name: "Quad Split",
							Usr:  false,
						},
						{
							ID:   2,
							Name: "10 Split A",
							Usr:  false,
						},
						{
							ID:   3,
							Name: "10 Split B",
							Usr:  false,
						},
						{
							ID:   4,
							Name: "10 Split C",
							Usr:  false,
						},
						{
							ID:   5,
							Name: "10 Split D",
							Usr:  false,
						},
						{
							ID:   6,
							Name: "9 Split",
							Usr:  false,
						},
						{
							ID:   7,
							Name: "16 Split",
							Usr:  false,
						},
						{
							ID:   8,
							Name: "25 Split",
							Usr:  false,
						},
						{
							ID:   9,
							Name: "36Split", // OH PLEASE
							Usr:  false,
						},
					},
					SDP: &SDP{
						raw: nil,
					},
				},
			},
			Error: nil,
		},
	}

	assertUnmarshalJSON(t, tcs)
}

func TestUnmarshalMultiviewerR(t *testing.T) {
	t.Parallel()
	tcs := []testCase[string, MultiviewerR]{
		{
			Name:     "Empty JSON",
			Input:    "{}",
			Expected: MultiviewerR{},
			Error:    nil,
		},
		{
			Name: "Example on API document",
			Input: `{
"index": 0,
"name": "Multiviewer1",
"preset": null,
"presets": [
{
"id": 0,
"name": "Full",
"usr": false
},
{
"id": 1,
"name": "Quad Split",
"usr": false
},
{
"id": 2,
"name": "10 Split A",
"usr": false
},
{
"id": 3,
"name": "10 Split B",
"usr": false
},
{
"id": 4,
"name": "10 Split C",
"usr": false
},
{
"id": 5,
"name": "10 Split D",
"usr": false
},
{
"id": 6,
"name": "9 Split",
"usr": false
},
{
"id": 7,
"name": "16 Split",
"usr": false
},
{ "id": 8,
"name": "25 Split",
"usr": false
},
{
"id": 9,
"name": "36 Split",
"usr": false
}
],
"sdp": "v=0\ns=Multiviewer 0\no=- 1 1 IN IP4 192.168.10.42\nc=IN IP4 239.168.10.42\nt=0 0\nm=video 50000 RTP/AVP 96\na=rtpmap:96 H264/90000\n",
"uuid": "fd839e94-a9e9-570b-a5aa-bf99575f364f"
}`,
			Expected: MultiviewerR{
				base: base{
					UUID: "fd839e94-a9e9-570b-a5aa-bf99575f364f",
				},
				Index: 0,
				Name:  "Multiviewer1",
				Presets: []MultiviewerPresetR{
					{
						ID:   0,
						Name: "Full",
						Usr:  false,
					},
					{
						ID:   1,
						Name: "Quad Split",
						Usr:  false,
					},
					{
						ID:   2,
						Name: "10 Split A",
						Usr:  false,
					},
					{
						ID:   3,
						Name: "10 Split B",
						Usr:  false,
					},
					{
						ID:   4,
						Name: "10 Split C",
						Usr:  false,
					},
					{
						ID:   5,
						Name: "10 Split D",
						Usr:  false,
					},
					{
						ID:   6,
						Name: "9 Split",
						Usr:  false,
					},
					{
						ID:   7,
						Name: "16 Split",
						Usr:  false,
					},
					{
						ID:   8,
						Name: "25 Split",
						Usr:  false,
					},
					{
						ID:   9,
						Name: "36 Split",
						Usr:  false,
					},
				},
				SDP: &SDP{
					raw: &sdp.Session{},
				},
			},
			Error: nil,
		},
	}

	assertUnmarshalJSON(t, tcs)
}

//func TestUnmarshalSceneRs
