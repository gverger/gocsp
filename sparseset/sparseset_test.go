package sparseset

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewSet(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 0)
	is.Equal(s.AllValues(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	is.Equal(s.Max, 9)
	is.Equal(s.Min, 0)
	is.Equal(s.Size, 10)

	sofs := NewSet(10, 4)
	is.Equal(sofs.AllValues(), []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	is.Equal(sofs.Max, 13)
	is.Equal(sofs.Min, 4)
	is.Equal(sofs.Size, 10)
}

func TestRemove(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 0)
	s.Remove(7)
	is.Equal(s.AllValues(), []int{0, 1, 2, 3, 4, 5, 6, 8, 9})
	is.Equal(s.Max, 9)
	is.Equal(s.Min, 0)
	is.Equal(s.Size, 9)
	is.True(s.Contains(9))
	is.True(!s.Contains(7))
}

func TestRemoveOffset(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.Remove(7)
	is.Equal(s.AllValues(), []int{4, 5, 6, 8, 9, 10, 11, 12, 13})
	is.Equal(s.Max, 13)
	is.Equal(s.Min, 4)
	is.Equal(s.Size, 9)
	is.True(s.Contains(12))
	is.True(!s.Contains(7))
}

func TestRemoveAbsent(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.Remove(27)
	is.Equal(s.AllValues(), []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	is.Equal(s.Max, 13)
	is.Equal(s.Min, 4)
	is.Equal(s.Size, 10)
	is.True(s.Contains(12))
	is.True(!s.Contains(3))
}
