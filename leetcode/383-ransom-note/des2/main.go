package des1

func canConstruct(ransomNote string, magazine string) bool {
	lettersCount := [26]int{}
	for i := 0; i < len(magazine); i++ {
		lettersCount[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		index := ransomNote[i] - 'a'
		if lettersCount[index] == 0 {
			return false
		}
		lettersCount[index]--
	}
	return true
}
