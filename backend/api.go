package main

import (
	"io"
	"net/http"
)

func GetResponseBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func GetDrafterDraftByURL(url string) (Draft, error) {
	draft := Draft{}

	respBody, err := GetResponseBody(url)
	if err != nil {
		return draft, err
	}

	draft, err = GetDraftFromDrafterBody(string(respBody))
	if err != nil {
		return draft, err
	}

	return draft, nil
}

func GetDrafterSeriesByURL(url string) (PicksAndBans, error) {
	picksAndBans := PicksAndBans{}

	respBody, err := GetResponseBody(url)
	if err != nil {
		return picksAndBans, err
	}

	series, err := GetSeriesFromDrafterBody(string(respBody))
	if err != nil {
		return picksAndBans, err
	}

	for i := range series {
		url := "https://drafter.lol/draft/" + series[i]
		draft, err := GetDrafterDraftByURL(url)
		if err != nil {
			return picksAndBans, err
		}

		picksAndBans = append(picksAndBans, draft)
	}

	return picksAndBans, nil
}
