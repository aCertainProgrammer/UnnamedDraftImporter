package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetResponseBody(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:150.0) Gecko/20100101 Firefox/150.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accept-Language", "en-Us,en;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Priority", "u=0, i")
	req.Header.Add("TE", "trailers")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(respBody))

	return respBody, nil
}

func GetDrafterDraftByURL(url_string string) (Draft, error) {
	draft := Draft{}

	respBody, err := GetResponseBody(url_string)
	if err != nil {
		return draft, err
	}

	urlStruct, err := url.Parse(url_string)
	if err != nil {
		return draft, err
	}

	game, err := strconv.Atoi(urlStruct.Query().Get("game"))
	if err != nil {
		return draft, err
	}

	draft, err = GetDraftFromDrafterBody(string(respBody), game)
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
		url := "https://drafter.lol/draft/" + series[i] + "?game=" + strconv.Itoa(i+1)
		fmt.Println(url)
		draft, err := GetDrafterDraftByURL(url)
		if err != nil {
			return picksAndBans, err
		}

		picksAndBans = append(picksAndBans, draft)
	}

	return picksAndBans, nil
}
