package otxapi

import "net/http"

type UserDetail struct {
	AwardCount      *int    `json:"award_count"`
	FollowerCount   *int    `json:"follower_count"`
	SubscriberCount *int    `json:"subscriber_count"`
	IndicatorCount  *int    `json:"indicator_count"`
	PulseCount      *int    `json:"pulse_count"`
	MemberSince     *string `json:"member_since"`
	UserId          *int    `json:"user_id"`
	Username        *string `json:"username"`
}

func (r UserDetail) String() string {
	return Stringify(r)
}

type OTXUserDetailService struct {
	client *Client
}

func (c *OTXUserDetailService) Details() (*UserDetail, error) {
	req, err := c.client.newRequest(http.MethodGet, UserURLPath, nil)
	if err != nil {
		return nil, err
	}

	var user UserDetail
	if err := c.client.do(req, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
