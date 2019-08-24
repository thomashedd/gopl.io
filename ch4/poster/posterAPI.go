// provides a Go API for downloading movie posters
// this is the setup go file
package poster

import "time"

const MyAPIKey = "dc6882ff" // not used because we use the below
const ApiAppend = "&apikey=dc6882ff"
const PosterURL = "http://img.omdbapi.com/?"

const (
	TimestampFormat = "06 Oct 1939" // convert to Go base time
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error

	t.Time, err = time.Parse(`"`+TimestampFormat+`"`, string(data))
	//*t = OMDbTime(parsed)
	return err
}

type MoviesSearchResult struct {
	Items []*Movie
}

type Movie struct {
	Title    string
	Year     string
	Released Time `json:"released"`
	Runtime  int
	Actors   []string
}
