package openrec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetMovies Get movie list with specified oa/sort/page status.
func GetMovies(OnAirStatus bool, sort string, page uint) ([]Movie, error) {
	oa := 0
	if OnAirStatus {
		oa = 1
	}
	url := fmt.Sprintf("https://public.openrec.tv/external/api/v5/movies?onair_status=%d&sort=%s&page=%d", oa, sort, page)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	list := make([]Movie, 40)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Printf("BODY : %v\n", string(body))

	if err := json.Unmarshal(body, &list); err != nil {
		return nil, err
	}
	return list, err
}

// GetMoviesByChannelID Get movie list by ChannelID.
func GetMoviesByChannelID(chanID string) ([]Movie, error) {
	url := fmt.Sprintf("https://public.openrec.tv/external/api/v5/movies?channel_id=%s", chanID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	list := make([]Movie, 40)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Printf("BODY : %v\n", string(body))

	if err := json.Unmarshal(body, &list); err != nil {
		return nil, err
	}
	return list, err
}
