package main

import (
	"reflect"
	"testing"
)

func TestNormalizeDraft(t *testing.T) {
	draft := Draft{
		Picks: [10]string{"", "", "", "monkeyking", "none", "jarvaniv", "ksante", "chogath", "reksai", "fiddlesticks"},
		Bans:  [10]string{"", "", "", "monkeyking", "none", "jarvaniv", "ksante", "chogath", "reksai", "fiddlesticks"},
	}

	expected_result := Draft{
		Picks: [10]string{"", "", "", "wukong", "", "jarvan", "ksante", "chogath", "reksai", "fiddlesticks"},
		Bans:  [10]string{"", "", "", "wukong", "", "jarvan", "ksante", "chogath", "reksai", "fiddlesticks"},
	}

	result := normalizeDraft(draft)

	if !reflect.DeepEqual(result, expected_result) {
		t.Errorf("NormalizeDraft() returned %s, expected %s\n", result, expected_result)
	}
}
