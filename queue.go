package main

// inefficient slice implementation
// also good case for generics :)
type Queue []int

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}
