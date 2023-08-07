package des1

import "math/rand"

type RandomizedSet struct {
	values []int
	loc    map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values: make([]int, 0),
		loc:    make(map[int]struct{}),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.loc[val]; ok {
		return false
	}
	r.values = append(r.values, val)
	r.loc[val] = struct{}{}
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if _, ok := r.loc[val]; !ok {
		return false
	}
	delete(r.loc, val)
	for i, v := range r.values {
		if v == val {
			r.values = append(r.values[:i], r.values[i+1:]...)
			return true
		}
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.values[rand.Intn(len(r.values))]
}
