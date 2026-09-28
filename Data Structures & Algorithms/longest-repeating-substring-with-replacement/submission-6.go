func characterReplacement(s string, k int) int {

    l := 0
	// reps := 0
	longestRep := 0
	mostFrequent := 0
	count := make(map[byte]int)

	for r := 0; r < len(s); r++{

	   char := s[r]
	   count[char]++

	   if count[char] > mostFrequent {
			mostFrequent = count[char]
	   }

	   toReplace := (r - l + 1) - mostFrequent

	   if toReplace > k {
		  char = s[l]
		  count[char]--
		  l++
	   }
      
	  if (r - l + 1) > longestRep {
		longestRep = r - l + 1
	  }


	}
	return longestRep

}
