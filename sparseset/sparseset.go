package sparseset

type Set struct {
	values []int
	index  []int

	n      int
	offset int

	Size int // nb of values
	min  int // min value (before offset)
	max  int // max value (before offset)
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
		min:    0,
		max:    n - 1,
	}
}

func (s Set) Min() int {
	if s.IsEmpty() {
		panic("empty set")
	}
	return s.min + s.offset
}

func (s Set) Max() int {
	if s.IsEmpty() {
		panic("empty set")
	}
	return s.max + s.offset
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

// Remove the value from the set
func (s *Set) Remove(value int) bool {
	if !s.Contains(value) {
		return false
	}

	value -= s.offset
	s.exchangeValues(s.index[value], s.Size-1)

	s.Size--
	s.updateBoundsValRemoved(value)
	return true
}

// Return true if the value is in the set
func (s *Set) Contains(value int) bool {
	value -= s.offset
	return s.contains(value)
}

func (s *Set) contains(value int) bool {
	if value < 0 || value >= s.n {
		return false
	}
	return s.index[value] < s.Size
}

func (s *Set) IsEmpty() bool {
	return s.Size == 0
}

// update min and max.
// Warning val is already shifted
func (s *Set) updateBoundsValRemoved(val int) {
	s.updateMaxValRemoved(val)
	s.updateMinValRemoved(val)
}

func (s *Set) updateMaxValRemoved(val int) {
	if s.IsEmpty() || s.max != val {
		return
	}

	if s.contains(val) {
		panic("updateMaxValRemoved")
	}

	for v := val - 1; v >= s.min; v-- {
		if s.contains(v) {
			s.max = v
			return
		}
	}
}

func (s *Set) updateMinValRemoved(val int) {
	if s.IsEmpty() || s.min != val {
		return
	}

	if s.contains(val) {
		panic("updateMinValRemoved")
	}

	for v := val + 1; v <= s.max; v++ {
		if s.contains(v) {
			s.min = v
			return
		}
	}
}

func (s *Set) exchangeValues(a, b int) {
	i1 := s.values[a]
	i2 := s.values[b]

	s.index[i1] = b
	s.index[i2] = a
	s.values[a] = i2
	s.values[b] = i1
}
