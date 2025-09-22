package main

func NormalizeDraft(draft Draft) Draft {
	draft.Picks = [10]string(NormalizeChampionsArray(draft.Picks[:]))
	draft.Bans = [10]string(NormalizeChampionsArray(draft.Bans[:]))

	return draft
}

func NormalizeChampionsArray(champions []string) []string {
	for i := range champions {
		switch champions[i] {
		case "monkeyking":
			champions[i] = "wukong"
		case "jarvaniv":
			champions[i] = "jarvan"
		case "none":
			champions[i] = ""
		}
	}

	return champions
}
