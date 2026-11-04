func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]bool, len(s))
	maxSub := 0
	start:= 0

for i, r := range s {
	if set[r] {
		for set[r] {
			delete(set, rune(s[start]))
			start++
		}
	}

	set[r] = true
	maxSub = max(maxSub, i - start + 1)

}

return maxSub
}
