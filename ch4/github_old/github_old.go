//provides a Go API for the Github issue tracker.
// see https://developer.github.com/v3/search/#search-issues
package github

import (
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

const (
	TimestampFormat = "2006-01-02T15:04:05Z"
)

//type Time time.Time

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	t.Time, err = time.Parse(`"`+TimestampFormat+`"`, string(data))
	//*t = Time(parsed)
	//pp.Println("unmarshaling github time:", string(data), parsed.String(), t.Day())

	return err
}

type IssuesSearchResult struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt Time   `json:"created_at"`
	UpdatedAt Time   `json:"updated_at"`
	Body      string // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
