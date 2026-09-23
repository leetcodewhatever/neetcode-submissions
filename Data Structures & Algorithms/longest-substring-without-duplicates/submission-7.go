func lengthOfLongestSubstring(s string) int {
    if s == "" {
		return 0
	}

    charOccurences := make(map[byte]int)
	l, maxLength := 0,0

	for r := 0 ; r < len(s) ; r++ {
     
		char := s[r]
		charOccurences[char] += 1

		for charOccurences[char] > 1 {
			charL := s[l]
			charOccurences[charL]--
			l++
		}
 
		if r - l + 1 > maxLength {
			maxLength = r - l + 1
		}
	 
	}
	
	return maxLength
}
