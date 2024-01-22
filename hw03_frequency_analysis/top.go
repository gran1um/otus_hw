package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) ([]string, error) {
	words := strings.Fields(text)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	type kv struct {
		Key   string
		Value int
	}
	ss := make([]kv, 0, len(frequency))
	for k, v := range frequency {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	var top10 []string
	for i := 0; i < len(ss) && i < 10; i++ {
		top10 = append(top10, ss[i].Key)
	}

	return top10, nil
}
