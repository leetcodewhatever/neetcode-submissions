func checkInclusion(s1 string, s2 string) bool {
	// s1 = "abc" s2 = "abdbdc"
	need := createFrequencyMapFor(s1)
	window := make(map[byte]int)
    l := 0

	for r, _ := range s2 {
       
	   char := s2[r]

	   if !existIn(need, char){
		  window = make(map[byte]int)
		  l=r+1
		  continue
	   }

	   window[char]++

	   for window[char] > need[char] {
	       charL := s2[l]
		   window[charL]--
		   l++
	   }

	   if (r-l+1) == len(s1){
			return true
	   }

	  
	} 

	return false
}

func createFrequencyMapFor(s string) map[byte]int {
	count := make(map[byte]int)
	for i:=0; i<len(s); i++{
		count[s[i]]++
	}
	return count
}


func existIn(s map[byte]int, char byte) bool{
	_, ok := s[char]
	return ok
}
