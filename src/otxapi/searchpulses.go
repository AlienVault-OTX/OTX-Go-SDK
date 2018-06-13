package otxapi

type OTXSearchPulsesSearvice struct {
	client *Client
}

// Pulse from search/pulse
type PulseSearchResponse struct {
	Next       *string `json:"next,omitempty"`
	ExactMatch *string `json:"exact_match,omitempty"`
	Previous   *string `json:"previous,omitempty"`
	Count      *int    `json:"count,omitempty"`
	Results    []Pulse `json:"results,omitempty"`
}

func (r PulseSearchResponse) String() string {
	return Stringify(r)
}
