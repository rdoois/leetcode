package main

import "testing"

func TestZigzagConversion(t *testing.T) {
	tableTests := []struct {
		description string
		s           string
		rows        int
		expected    string
	}{
		{"Case 1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"Case 2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"Case 3", "A", 3, "A"},
		{"Case 4", "wlrbbmqbhcdarzowkkyhiddqscdxrjmowfrxsjybldbefsarcbynecdyggxxpklorellnmpapqfwkhopkmcoqhnwnkuewhsqmgbbuqcljjivswmdkqtbxixmvtrrbljptnsnfwzqfjmafadrrwsofsbcnuvqhffbsaqxwpqcacehchzvfrkmlnozjkpqpxrjxkitzyxacbhhkicqcoendtomfgdwdwfcgpxiqvkuytdlcgdewhtaciohordtqkvwcsgspqoqmsboaguwnnyqxnzlgdgwpbtrwblnsadeuguumoqcdrubetokyxhoachwdvmxxrdryxlmndqtukwagmlejuukwcibxubumenmeyatdrmydiajxloghiqfmzhl", 80, "wfalfbocrhshhbqaxwbvqydmuxkvqnwombcptxhbqexcscbrdfaudaocrrrsedyzwhcxorcqlwrhomkdzmnkavudyffuqhargtimkuudjmekdfldwqqnaaszosgcwznmdfjllxnkberspwjjnqrumtptuopxbkwjrpwfljwcrbxgixrkdbsrigxjttluyvzzbbmynulxxxmdiaqebxcynebbnmfthnesqhwyakkuardigtcmcadbwqoryscbmnvosyeiemdcjnqidjdoayltqjgcopxgqmslxufgoxbgsgpbdchkgwwilmdvqoqwkfrsfqmehctzlwgdhleprlnuxomkihpnqoawvipnkcqhuafqytwothkcdwhmleokcdpg"},
	}

	for _, tt := range tableTests {
		t.Run(tt.description, func(t *testing.T) {
			res := convert(tt.s, tt.rows)
			if res != tt.expected {
				t.Errorf("given %s got %s, want %s", tt.s, res, tt.expected)
			}
		})
	}

}
