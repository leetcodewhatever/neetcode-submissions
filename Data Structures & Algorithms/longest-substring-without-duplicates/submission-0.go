func lengthOfLongestSubstring(s string) int {
// use l,r
// normal case: if the next charachtere is different keep moving r
// edge case: if the next charachter is equal move l by one
// abcb
// l r
// if l = r || occurences[r] > 1 
    charOccurences := make(map[byte]int)
	l := 0
	length := 0
	maxLength := 0
	// var substring []rune
	for r:=1 ; r<len(s) ; r++ {
     
	 char := s[r]
	 charOccurences[char] += 1

	 if s[l] == s[r] || charOccurences[char] > 1 {
        l=r
		charOccurences = make(map[byte]int)
		length = 1
		charOccurences[char] += 1
		// clear(substring) 
	 } else {
    //    append(substring, s[r])
		length++
		
		if length > maxLength {
			maxLength = length
		}
	 }

	}
	
	if maxLength == 0 {
		return 1
	}
	return maxLength
}
