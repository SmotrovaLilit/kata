package des1

func canConstruct(ransomNote string, magazine string) bool {
	var isFound bool
	for i := 0; i < len(ransomNote); i++ {
		isFound = false
		for j := 0; j < len(magazine); j++ {
			if ransomNote[i] == magazine[j] {
				magazine = magazine[:j] + magazine[j+1:]
				isFound = true
				break
			}
		}
		if !isFound {
			return false
		}
	}
	return true
}
