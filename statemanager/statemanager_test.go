package statemanager_test

import (
	"testing"

	"github.com/matryer/is"
	"gverger.com/csp/statemanager"
)

func TestStateManager(t *testing.T) {
	is := is.New(t)

	m := statemanager.New()

	a := m.NewInt(12)
	is.Equal(a.Value, 12)

	b := m.NewInt(43)
	is.Equal(b.Value, 43)

	c := m.NewInt(-3)
	is.Equal(c.Value, -3)

	a.Value = 11

	stateAfterA11 := m.Save()

	b.Value = 34

	stateAfterB34 := m.Save()

	c.Value = 55

	is.Equal(a.Value, 11)
	is.Equal(b.Value, 34)
	is.Equal(c.Value, 55)

	m.Restore(stateAfterB34)

	is.Equal(a.Value, 11)
	is.Equal(b.Value, 34)
	is.Equal(c.Value, -3)

	m.Restore(stateAfterA11)

	is.Equal(a.Value, 11)
	is.Equal(b.Value, 43)
	is.Equal(c.Value, -3)
}
