package otxapi

type OTXThreatIntelFeedService struct {
	client *Client
}

type ThreatIntelFeed struct {
	Pulses []Pulse `json:"results"`
	// These fields provide the page values for paginating through a set of
	// results.  Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.
	//NextPageNum  int   Coming soon
	//PrevPageNum  int   Coming soon
	NextPageString *string `json:"next"`
	PrevPageString *string `json:"prev"`
	Count          int     `json:"count"`
}

func (r ThreatIntelFeed) String() string {
	return Stringify(r)
}
