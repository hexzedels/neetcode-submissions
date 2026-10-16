func groupAnagrams(strs []string) [][]string {
	alph := "abcdefghijklmnopqrstuvwxyz"
	hashmap := make(map[[26]uint8][]string)
	for _, s := range strs {
		encoded := [26]uint8{}
		for _, chr := range s {
			encoded[strings.IndexRune(alph, chr)]++
		}

		anagrams, ok := hashmap[encoded]
		if !ok {
			hashmap[encoded] = []string{s}
		}
		hashmap[encoded] = append(anagrams, s)
	}

	answ := make([][]string, 0, len(hashmap))
	for _, value := range hashmap {
		answ = append(answ, value)
	}

	return answ
}