package des1

func rob(nums []int) int {
	r := newRobbing(nums)
	return r.maxRob(0)
}

type robbing struct {
	cache map[int]int
	nums  []int
}

func newRobbing(nums []int) *robbing {
	return &robbing{
		cache: make(map[int]int),
		nums:  nums,
	}
}

func (r *robbing) maxRob(index int) int {
	if v, ok := r.cache[index]; ok {
		return v
	}
	if index >= len(r.nums) {
		return 0
	}
	sum := max(r.maxRob(index+1), r.maxRob(index+2)+r.nums[index])
	r.cache[index] = sum
	return sum
}
