package otxapi

type OTXIndicatorService struct {
	client *Client
}

type Indicator struct {
	ID          *string `json:"_id"`
	Indicator   *string `json:"indicator"`
	Type        *string `json:"type"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
	Title       *string `json:"title,omitempty"`
}

func (r Indicator) String() string {
	return Stringify(r)
}
