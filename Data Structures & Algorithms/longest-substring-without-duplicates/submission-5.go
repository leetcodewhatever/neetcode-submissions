func lengthOfLongestSubstring(s string) int {
// use l,r
// normal case: if the next charachtere is different keep moving r
// edge case: if the next charachter is equal move l by one
// abcb
// l r
// if l = r || occurences[r] > 1 
    if s == "" {
		return 0
	}

    charOccurences := make(map[byte]int)
	l := 0
	// length := 1
	maxLength := 0
	// var substring []rune
	for r:=0 ; r<len(s) ; r++ {
     
	 char := s[r]
	 charOccurences[char] += 1

	 for charOccurences[char] > 1 {
        charL := s[l]
		charOccurences[charL]--
		l++

		// length = 1
		// charOccurences[char] += 1
		// clear(substring) 
	 } 
    //    append(substring, s[r])
		
	if r - l + 1 > maxLength {
		maxLength = r - l + 1
	}
	 

	}
	
	return maxLength
}
