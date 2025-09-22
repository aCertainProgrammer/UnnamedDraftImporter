package main

import (
	"strings"
)

type Draft struct {
	Picks [10]string `json:"picks"`
	Bans  [10]string `json:"bans"`
}

type Snapshot struct {
	Picks [10]string `json:"picks"`
	Bans  [10]string `json:"bans"`

	Name     string `json:"name"`
	BlueName string `json:"blue_name"`
	RedName  string `json:"red_name"`
}

type (
	PicksAndBans []Draft
	Snapshots    []Snapshot
	Series       []string
)

func GetDraftFromDrafterBody(body string) (Draft, error) {
	draft := Draft{}
	words := (strings.Split(string(body), "\\\""))

	pickIndex := 0
	banIndex := 0

	for i := range words {
		word := words[i]

		if strings.Contains(word, "blueBan") || strings.Contains(word, "redBan") {
			draft.Bans[banIndex] = strings.ToLower(words[i+2])
			banIndex++
		}
		if strings.Contains(word, "bluePick") || strings.Contains(word, "redPick") {
			draft.Picks[pickIndex] = strings.ToLower(words[i+2])
			pickIndex++
		}
	}

	return NormalizeDraft(draft), nil
}

func GetSeriesFromDrafterBody(body string) (Series, error) {
	series := Series{}
	words := (strings.Split(string(body), "\\\""))

	for i := range words {
		word := words[i]
		if strings.Contains(word, "draftIDs") {
			for j := 1; words[i+j*2] != "finishedDraftIDs"; j++ {
				series = append(series, words[i+j*2])
			}

			break
		}
	}

	return series, nil
}
