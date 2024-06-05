package sparseset

import (
	"errors"

	"gverger.com/csp/statemanager"
)

var errNoElement = errors.New("no such element")

type Set struct {
	values []int
	index  []int

	n      int
	offset int

	size *statemanager.StateInt // nb of values
	min  *statemanager.StateInt // min value (before offset)
	max  *statemanager.StateInt // max value (before offset)
}

// Create a set of n consecutive values starting at offset
func NewSet(stateManager *statemanager.StateManager, n int, offset int) *Set {
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
		size:   stateManager.NewInt(n),
		min:    stateManager.NewInt(0),
		max:    stateManager.NewInt(n - 1),
	}
}

func (s Set) Size() int {
	return s.size.Value
}

// Return the min value, with an error if the set is empty
func (s Set) Min() (int, error) {
	if s.IsEmpty() {
		return 0, errNoElement
	}
	return s.min.Value + s.offset, nil
}

// Return the max value, with an error if the set is empty
func (s Set) Max() (int, error) {
	if s.IsEmpty() {
		return 0, errNoElement
	}
	return s.max.Value + s.offset, nil
}

// Return all values of the set in order
func (s Set) AllValues() []int {
	values := make([]int, 0, s.size.Value)
	for i := 0; i < s.n; i++ {
		if s.index[i] < s.size.Value {
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
	s.exchangeValues(s.index[value], s.size.Value-1)

	s.size.Value--
	s.updateBoundsValRemoved(value)
	return true
}

// Remove all values
func (s *Set) RemoveAll() {
	s.size.Value = 0
}

// Remove all values strictly smaller than value
func (s *Set) RemoveBelow(value int) {
	if s.IsEmpty() {
		return
	}
	if s.max.Value+s.offset < value {
		s.RemoveAll()
	} else {
		for val := s.min.Value; val < value-s.offset; val++ {
			s.Remove(val + s.offset)
		}
	}
}

// Remove all values strictly greater than value
func (s *Set) RemoveAbove(value int) {
	if s.IsEmpty() {
		return
	}
	if s.min.Value+s.offset > value {
		s.RemoveAll()
	} else {
		for val := value - s.offset + 1; val <= s.max.Value; val++ {
			s.Remove(val + s.offset)
		}
	}
}

// Remove all values strictly greater than value
func (s *Set) RemoveAllBut(value int) {
	if !s.Contains(value) {
		panic("remove all but non contained value")
	}

	value -= s.offset

	if s.index[value] != 0 {
		s.exchangeValues(s.index[value], 0)
	}

	s.size.Value = 1
	s.min.Value = value
	s.max.Value = value
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
	return s.index[value] < s.size.Value
}

func (s *Set) IsEmpty() bool {
	return s.size.Value == 0
}

// update min and max.
// Warning val is already shifted
func (s *Set) updateBoundsValRemoved(val int) {
	s.updateMaxValRemoved(val)
	s.updateMinValRemoved(val)
}

func (s *Set) updateMaxValRemoved(val int) {
	if s.IsEmpty() || s.max.Value != val {
		return
	}

	if s.contains(val) {
		panic("updateMaxValRemoved")
	}

	for v := val - 1; v >= s.min.Value; v-- {
		if s.contains(v) {
			s.max.Value = v
			return
		}
	}
}

func (s *Set) updateMinValRemoved(val int) {
	if s.IsEmpty() || s.min.Value != val {
		return
	}

	if s.contains(val) {
		panic("updateMinValRemoved")
	}

	for v := val + 1; v <= s.max.Value; v++ {
		if s.contains(v) {
			s.min.Value = v
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
