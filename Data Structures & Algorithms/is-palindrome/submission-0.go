func isPalindrome(s string) bool {

	s = reconstructString(s)
	fmt.Println(s)
	mid := len(s) - 1/2
	j := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}

	return true
}

func reconstructString(s string) string {
	var newString []rune
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"

	alphabetMap := make(map[rune]bool)
	for _, char := range alphabet {
		alphabetMap[char] = true
	}

	s = strings.ToLower(s)

	for _, v := range s {

		ok, _ := alphabetMap[v]
		
        if !ok {
			continue
		}
        
		newString = append(newString, v)
	}

	return string(newString)

}