package otxapi

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

type OTXPulseDetailService struct {
	client *Client
}

// Get returns a single *Pulse.
func (svc *OTXPulseDetailService) Get(pulseID string) (*Pulse, error) {
	req, err := svc.client.newRequest(http.MethodGet, PulseDetailURLPath+pulseID, nil)
	if err != nil {
		return nil, err
	}

	var p Pulse
	if err := svc.client.do(req, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Pulse represents an OTX Pulse.
type Pulse struct {
	ID                string           `json:"id"`
	Author            string           `json:"author_name"`
	Name              string           `json:"name"`
	Description       string           `json:"description,omitempty"`
	CreatedAt         Timestamp        `json:"created,omitempty"`
	ModifiedAt        Timestamp        `json:"modified"`
	References        []string         `json:"references,omitempty"`
	Tags              []string         `json:"tags,omitempty"`
	Indicators        []PulseIndicator `json:"indicators,omitempty"`
	Revision          float32          `json:"revision,omitempty"`
	TLP               string           `json:"tlp"`
	Public            bool             `json:"public"`
	Adversary         string           `json:"adversary"`
	TargetedCountries []string         `json:"targeted_countries"`
	Industries        []string         `json:"industries"`
}

func (p Pulse) String() string {
	//return Stringify(r)
	s := fmt.Sprintf("%s\t%s\t%s\t(rev %f)", p.ID, p.Name, p.CreatedAt, p.Revision)
	for _, ind := range p.Indicators {
		s += fmt.Sprintf("\n\t%s", ind.String())
	}
	return s
}

type PulseIndicator struct {
	Content      string     `json:"content"`
	Indicator    string     `json:"indicator"`
	Description  string     `json:"description,omitempty"`
	Created      Timestamp  `json:"created"`
	Expiration   *Timestamp `json:"expiration,omitempty"`
	Active       int        `json:"is_active"`
	Title        string     `json:"title"`
	AccessReason string     `json:"access_reason"`
	AccessType   string     `json:"access_type"`
	AccessGroups []string   `json:"access_groups"`
	Type         string     `json:"type"`
	ID           int        `json:"id"`
	Observations int        `json:"observations"`

	// These fields are currently always nil, in responses.
	//Role json.RawMessage `json:"role"`
}

func (p PulseIndicator) String() string {
	return fmt.Sprintf("%d\t%s\t%s\t%s\t%s",
		p.ID,
		p.Created,
		p.Title,
		p.Type,
		p.Indicator,
	)
}

// List returns a *PulseList, which maps to a single page of pulses the user
// is subscribed to.
//
// Passing nil *ListOptions will result in retrieving the first page of
// subscribed pulses, with the maximum number of results for the page (20).
func (svc *OTXPulseDetailService) List(opt *ListOptions) (*PulseList, error) {
	if opt == nil {
		opt = &ListOptions{
			Page:    1,
			PerPage: 20,
		}
	}

	req, err := svc.client.newRequest(http.MethodGet, SubscriptionsURLPath, nil)
	if err != nil {
		return nil, err
	}

	if err := addURLOptions(req.URL, opt); err != nil {
		return nil, fmt.Errorf("applying url options: %v", err)
	}

	var list PulseList
	if err := svc.client.do(req, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// PulseList represents a single page of results, where each page holds a list
// of pulses the user has subscribed to.
type PulseList struct {
	Pulses []Pulse `json:"results"`

	// These fields provide the page values for paginating through a set of
	// results.  Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.
	//NextPageNum  int   Coming soon
	//PrevPageNum  int   Coming soon
	NextPageString *string `json:"next"`
	PrevPageString *string `json:"prev"`

	// Count is the total number of results, across all pages.
	Count int `json:"count"`
}

func (r PulseList) String() string {
	return Stringify(r)
}

// ErrNoPage indicates that there is no "next" or "previous" page.
var ErrNoPage = errors.New("no more pages")

// NextPageOptions returns *ListOptions that can be used for retrieving the
// next page of results.
//
// If the PulseList contains the last page of results, this method will
// return ErrNoPage.
func (r PulseList) NextPageOptions() (*ListOptions, error) {
	if r.NextPageString == nil {
		return nil, ErrNoPage
	}
	opts, err := parseListOptions(*r.NextPageString)
	if err != nil {
		return nil, err
	}
	return opts, nil
}

// PrevPageOptions returns *ListOptions that can be used for retrieving the
// previous page of results.
//
// If the PulseList contains the first page of results, this method will return
// ErrNoPage.
func (r PulseList) PrevPageOptions() (*ListOptions, error) {
	if r.PrevPageString == nil {
		return nil, ErrNoPage
	}
	opts, err := parseListOptions(*r.PrevPageString)
	if err != nil {
		return nil, err
	}
	return opts, nil
}

// parseListOptions returns *ListOptions that have been parsed from the given
// URL string.
//
// If the "?limit" URL query parameter is < 0, or > 20, then limit of results
// per page will be set to 20.
//
// NOTE(nesv): 20 is currently the maximum number of results allowed, per page.
func parseListOptions(urlStr string) (*ListOptions, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	page, err := strconv.Atoi(u.Query().Get("page"))
	if err != nil {
		return nil, err
	}
	if page < 0 {
		page = 1
	}

	limit, err := strconv.Atoi(u.Query().Get("limit"))
	if err != nil {
		return nil, err
	}

	// Make sure we stay within page size limits.
	if limit > 20 || limit < 0 {
		limit = 20
	}

	return &ListOptions{
		Page:    page,
		PerPage: limit,
	}, nil
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.  our list options: ?limit=50&page_num=1
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	// The maximum value is 20.
	PerPage int `url:"limit,omitempty"`
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt *ListOptions) (string, error) {
	if opt == nil {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return "", err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// addURLOptions modifies the query parameters of u, in place.
//
// This function will return a non-nil error if u is nil.
func addURLOptions(u *url.URL, opt *ListOptions) error {
	if u == nil {
		return errors.New("nil url")
	}
	if opt == nil {
		return nil
	}

	qs, err := query.Values(opt)
	if err != nil {
		return err
	}

	u.RawQuery = qs.Encode()
	return nil
}
