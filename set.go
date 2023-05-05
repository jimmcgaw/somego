package main

// Set created for Leetcode exercise
// TOOD: this is a good use case for generics
type IntSet map[int]struct{}

func (intSet IntSet) Add(x int) {
	intSet[x] = struct{}{}
}

func (intSet IntSet) Update(nums []int) IntSet {
	combined := IntSet{}
	for x, _ := range intSet {
		combined.Add(x)
	}
	for _, x := range nums {
		combined.Add(x)
	}
	return combined
}

// Is X in Set?
func (intSet IntSet) Has(x int) bool {
	_, exists := intSet[x]
	return exists
}

func (intSet IntSet) Values() []int {
	vals := []int{}
	for x, _ := range intSet {
		vals = append(vals, x)
	}
	return vals
}

// SetA & SetB
func (intSet IntSet) Intersection(otherSet IntSet) IntSet {
	intersection := IntSet{}
	for x, _ := range intSet {
		isShared := otherSet.Has(x)
		if isShared {
			intersection.Add(x)
		}
	}
	return intersection
}

// SetA + SetB
func (intSet IntSet) Union(otherSet IntSet) IntSet {
	union := IntSet{}
	for x, _ := range intSet {
		union.Add(x)
	}
	for x, _ := range otherSet {
		union.Add(x)
	}
	return union
}

// SetA - SetB
func (intSet IntSet) Diff(otherSet IntSet) IntSet {
	diff := IntSet{}
	for x, _ := range intSet {
		isShared := otherSet.Has(x)
		if !isShared {
			diff.Add(x)
		}
	}
	return diff
}
