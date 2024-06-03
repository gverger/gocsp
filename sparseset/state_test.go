package sparseset

import (
	"testing"

	"github.com/matryer/is"
)

func TestState(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)

	s.Remove(7)
	s.Remove(12)

	is.Equal(s.AllValues(), []int{4, 5, 6, 8, 9, 10, 11, 13})

	state1 := s.SaveState()

	s.Remove(5)
	is.Equal(s.AllValues(), []int{4, 6, 8, 9, 10, 11, 13})

	state2 := s.SaveState()

	s.RemoveAll()
	is.Equal(s.AllValues(), []int{})

	s.RestoreState(state2)

	is.Equal(s.AllValues(), []int{4, 6, 8, 9, 10, 11, 13})

	s.RestoreState(state1)
	is.Equal(s.AllValues(), []int{4, 5, 6, 8, 9, 10, 11, 13})
}
