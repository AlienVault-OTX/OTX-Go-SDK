package otxapi

import (
	"strconv"
	"time"
)

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an RFC3339 or Unix timestamp. This is necessary for some
// fields since the GitHub API is inconsistent in how it represents times. All
// exported methods of time.Time can be called on Timestamp.
type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in RFC3339 or Unix format.
func (ts *Timestamp) UnmarshalJSON(data []byte) error {
	layouts := []string{
		time.RFC3339,
		"\"2006-01-02T15:04:05\"",
		"\\\"2006-01-02T15:04:05\\\"",
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, string(data))
		if err == nil {
			ts.Time = t
			return nil
		}
	}

	// If we fall through to here, let's try to parse the timestamp as
	// a UNIX time.
	n, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	ts.Time = time.Unix(n, 0)
	return nil
}

// Equal reports whether t and u are equal based on time.Equal
func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}
