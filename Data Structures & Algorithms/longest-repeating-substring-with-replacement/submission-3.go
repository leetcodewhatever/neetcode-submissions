func characterReplacement(s string, k int) int {
// move r as long as  l != r
// move l
/*
reps = l
for s[l] != s[r] && reps < k
	replace s[r] with s[l]
	k--
*/ 
    l := 0
	reps := 0
	longestRep := 0
	count := make(map[byte]int)

	for r := 0; r < len(s); r++{

	   char := s[l]
       
	   if s[l] != s[r] && reps < k {
		  count[char]++
		  reps++
		  continue
	   }
	   
	   if s[l] != s[r] && reps == k {
		  if count[char] > longestRep { longestRep = count[char]}
		  count[char] = r - l
		  reps = 0
		  reps++
		  l++
		  continue	
	   }
       
	   

		count[char]++

		if longestRep == 0 && r == len(s) - 1 {
			return count[char]
	   	}
	}

	return longestRep

}
