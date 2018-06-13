package otxapi

type OTXPulseDetailService struct {
	client *Client
}

// Pulse results from search/pulse
type Pulse struct {
	ID                *string     `json:"id,omitempty"`
	Name              *string     `json:"name,omitempty"`
	ModifiedTimestamp *string     `json:"modified,omitempty"`
	Adversary         *string     `json:"adversary,omitempty"`
	TLP               *string     `json:"TLP,omitempty"`
	TargetedCountries []string    `json:"targeted_countries"`
	AuthorName        *string     `json:"author_name,omitempty"`
	Revision          *int        `json:"revision,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Public            *int        `json:"public,omitempty"`
	References        []string    `json:"references,omitempty"`
	Tags              []string    `json:"tags,omitempty"`
	Industries        []string    `json:"industries,omitempty"`
	Indicators        []Indicator `json:"indicators,omitempty"`
}

func (r Pulse) String() string {
	return Stringify(r)
}
