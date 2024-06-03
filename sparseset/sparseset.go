package sparseset

type Set struct {
	values []int
	index  []int

	n      int
	offset int

	Size int
	Min  int
	Max  int
}

// Create a set of n consecutive values starting at offset
func NewSet(n int, offset int) *Set {
	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = i
	}

	index := make([]int, n)
	copy(index, values)

	return &Set{
		values: values,
		index:  index,
		n:      n,
		offset: offset,
		Size:   n,
		Min:    offset,
		Max:    n + offset - 1,
	}
}

// Return all values of the set in order
func (s Set) AllValues() []int {
	values := make([]int, 0, s.Size)
	for i := 0; i < s.n; i++ {
		if s.index[i] < s.Size {
			values = append(values, i+s.offset)
		}
	}

	return values
}

func (s *Set) Remove(value int) {
	if !s.Contains(value) {
		return
	}

	value -= s.offset
	s.exchangeValues(s.index[value], s.Size-1)

	s.Size--
}

func (s *Set) Contains(value int) bool {
	value -= s.offset
	if value < 0 || value >= s.n {
		return false
	}
	return s.index[value] < s.Size
}

func (s *Set) exchangeValues(a, b int) {
	i1 := s.values[a]
	i2 := s.values[b]

	s.index[i1] = b
	s.index[i2] = a
	s.values[a] = i2
	s.values[b] = i1
}
