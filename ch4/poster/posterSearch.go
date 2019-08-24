package poster

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// queries the Open Movie Database
func SearchMovies(terms []string) (*MoviesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(PosterURL + "?t" + q + "&apikey=" + MyAPIKey) // t for title.

	fmt.Println("debug response call:%v\n", resp)

	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result MoviesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
