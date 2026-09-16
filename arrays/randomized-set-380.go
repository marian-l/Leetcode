package arrays

import "math/rand"

type RandomizedSet struct {
	values  []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indices: make(map[int]int),
		values:  make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, exists := rs.indices[val]
	if exists {
		return false
	}

	rs.values = append(rs.values, val)
	rs.indices[val] = len(rs.values) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	_, exists := rs.indices[val]
	if !exists {
		return false
	}

	index := rs.indices[val]
	last := rs.values[len(rs.values)-1]

	rs.values[index] = last
	rs.indices[last] = index

	delete(rs.indices, val)
	rs.values = rs.values[:len(rs.values)-1]

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	random := rand.Int() % len(rs.values)
	return rs.values[random]
}