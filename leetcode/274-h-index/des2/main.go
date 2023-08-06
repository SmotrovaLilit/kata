package des1

func hIndex(citations []int) int {
	papersCounts := make([]int, len(citations)+1)
	nPapers := len(citations)
	for _, count := range citations {
		if count >= nPapers {
			papersCounts[nPapers]++ // count of papers with more than papers len citations
		} else {
			papersCounts[count]++ // count of papers with <count> citations
		}
	}
	k := len(citations)
	for s := papersCounts[len(citations)]; s < k; s += papersCounts[k] {
		// s - it is count of papers with >=k citations
		k--
	}
	return k
}
