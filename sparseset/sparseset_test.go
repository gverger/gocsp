package sparseset

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewSet(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 0)
	is.Equal(s.AllValues(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 9)
	is.Equal(valueOf(s.Min()), 0)
	is.Equal(s.Size, 10)

	sofs := NewSet(10, 4)
	is.Equal(sofs.AllValues(), []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(sofs.Max()), 13)
	is.Equal(valueOf(sofs.Min()), 4)
	is.Equal(sofs.Size, 10)
}

func TestRemove(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 0)
	is.True(s.Remove(7))
	is.Equal(s.AllValues(), []int{0, 1, 2, 3, 4, 5, 6, 8, 9})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 9)
	is.Equal(valueOf(s.Min()), 0)
	is.Equal(s.Size, 9)
	is.True(s.Contains(9))
	is.True(!s.Contains(7))
}

func TestRemoveOffset(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	is.True(s.Remove(7))
	is.Equal(s.AllValues(), []int{4, 5, 6, 8, 9, 10, 11, 12, 13})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 13)
	is.Equal(valueOf(s.Min()), 4)
	is.Equal(s.Size, 9)
	is.True(s.Contains(12))
	is.True(!s.Contains(7))
}

func TestRemoveMax(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	is.True(s.Remove(13))
	is.Equal(s.AllValues(), []int{4, 5, 6, 7, 8, 9, 10, 11, 12})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 12)
	is.Equal(valueOf(s.Min()), 4)
	is.Equal(s.Size, 9)
	is.True(s.Contains(12))
	is.True(!s.Contains(13))
}

func TestRemoveMin(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	is.True(s.Remove(4))
	is.Equal(s.AllValues(), []int{5, 6, 7, 8, 9, 10, 11, 12, 13})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 13)
	is.Equal(valueOf(s.Min()), 5)
	is.Equal(s.Size, 9)
	is.True(s.Contains(12))
	is.True(!s.Contains(4))
}

func TestRemoveAbsent(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	is.True(!s.Remove(27))
	is.Equal(s.AllValues(), []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 13)
	is.Equal(valueOf(s.Min()), 4)
	is.Equal(s.Size, 10)
	is.True(s.Contains(12))
	is.True(!s.Contains(3))
}

func TestRemoveAll(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.RemoveAll()
	is.Equal(s.AllValues(), []int{})
	is.True(errorOf(s.Max()) != nil)
	is.True(errorOf(s.Min()) != nil)
	is.Equal(s.Size, 0)
	is.True(!s.Contains(12))
	is.True(!s.Contains(3))
}

func TestRemoveAbove(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.RemoveAbove(10)
	is.Equal(s.AllValues(), []int{4, 5, 6, 7, 8, 9, 10})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 10)
	is.Equal(valueOf(s.Min()), 4)
	is.Equal(s.Size, 7)
	is.True(!s.Contains(12))
	is.True(s.Contains(7))
}

func TestRemoveAboveAll(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.RemoveAbove(3)
	is.Equal(s.AllValues(), []int{})
	is.True(errorOf(s.Max()) != nil)
	is.True(errorOf(s.Min()) != nil)
	is.Equal(s.Size, 0)
	is.True(!s.Contains(12))
	is.True(!s.Contains(7))
}

func TestRemoveBelow(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.RemoveBelow(10)
	is.Equal(s.AllValues(), []int{10, 11, 12, 13})
	is.NoErr(errorOf(s.Max()))
	is.NoErr(errorOf(s.Min()))
	is.Equal(valueOf(s.Max()), 13)
	is.Equal(valueOf(s.Min()), 10)
	is.Equal(s.Size, 4)
	is.True(s.Contains(12))
	is.True(!s.Contains(7))
}

func TestRemoveBelowAll(t *testing.T) {
	is := is.New(t)

	s := NewSet(10, 4)
	s.RemoveBelow(15)
	is.Equal(s.AllValues(), []int{})
	is.True(errorOf(s.Max()) != nil)
	is.True(errorOf(s.Min()) != nil)
	is.Equal(s.Size, 0)
	is.True(!s.Contains(12))
	is.True(!s.Contains(7))
}

func errorOf[T any](_ T, err error) error {
	return err
}

func valueOf[T any](value T, _ error) T {
	return value
}
