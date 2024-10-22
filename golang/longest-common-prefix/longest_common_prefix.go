package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var res = strs[0]
	var w string
	for _, word := range strs {
		if word == "" {
			return word
		}

		if res == word {
			continue
		}

		w = ""
		for i := range res {
			if len(word) <= i {
				break
			}

			if res[i] != word[i] {
				break
			}

			w += fmt.Sprintf("%c", res[i])
		}
		res = w
	}

	return res
}
