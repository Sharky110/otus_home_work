package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	str = strings.ReplaceAll(str, "\n\t", " ")
	str = strings.ReplaceAll(str, "\n", " ")
	str = strings.ReplaceAll(str, "\t", " ")
	words := strings.Split(str, " ")
	var dict []word
	for _, v := range words {
		if v == "" {
			continue
		}
		newWord := true
		for i := range dict {
			if dict[i].val == v {
				dict[i].amount += 1
				newWord = false
				break
			}
		}
		if newWord {
			dict = append(dict, word{
				val:    v,
				amount: 1,
			})
		}
	}

	sort.Slice(dict, func(i, j int) bool {
		if dict[i].amount == dict[j].amount {
			return strings.Compare(dict[i].val, dict[j].val) < 0
		} else {
			return dict[i].amount > dict[j].amount
		}
	})

	var length int
	if len(dict) > 10 {
		length = 10
	} else {
		length = len(dict)
	}

	var sls = make([]string, length)
	for i, v := range dict {
		if i > 9 {
			break
		}
		sls[i] = v.val
	}
	return sls
}

type word struct {
	val    string
	amount int
}
