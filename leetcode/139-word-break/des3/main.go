package des1

import "strings"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i < len(wordDict[j])-1 { // wordDict[j] is too long
				continue
			}
			if !strings.HasSuffix(s[:i+1], wordDict[j]) { // if ith substring of s is ended with wordDict[j]
				continue
			}

			if i == len(wordDict[j])-1 { // if s starts with wordDict[j]
				dp[i] = true
				break
			}
			if dp[i-len(wordDict[j])] { // or dp before wordDict[j] is true
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)-1]
}
