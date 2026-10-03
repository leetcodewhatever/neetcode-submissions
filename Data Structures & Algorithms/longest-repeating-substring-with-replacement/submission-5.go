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
		  if count[char] > longestRep { longestRep = count[char]}
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

		if s[l] == s[r] && r == len(s) - 1 {
			return count[char]
		}
		
		// fmt.Printf("r %v \n", r)
		// fmt.Printf("reps %v \n", reps)
		// fmt.Printf("count %v \n", count)
		// fmt.Printf("longestRep %v \n", longestRep)
	}

	return longestRep

}
