package variables

import (
	"fmt"
	log "log/slog"
	"os"
)

type IntVariable struct {
	name        string
	domain      Domain
	constraints []BoundConstraint
}

// Constraints implements Variable.
func (v *IntVariable) Constraints() {
}

// ConstraintAdded implements Variable.
func (v *IntVariable) ConstraintAdded(c BoundConstraint) {
	v.constraints = append(v.constraints, c)
}

// NbConstraints implements Variable.
func (v IntVariable) NbConstraints() int {
	n := 0
	for _, c := range v.constraints {
		if !c.IsDone() {
			n++
		}
	}
	return n
}

// Name implements Variable.
func (v IntVariable) Name() string {
	return v.name
}

// Dom implements Variable.
func (v IntVariable) Dom() Domain {
	return v.domain
}

func (v *IntVariable) SetDom(d Domain) {
	v.domain = d
}

type EnumeratedDomain struct {
	Values []int
}

// Size implements Domain.
func (d *EnumeratedDomain) Size() int {
	return len(d.Values)
}

// Fix implements Domain.
func (d *EnumeratedDomain) Fix(value int) bool {
	for _, val := range d.Values {
		if val == value {
			d.Values = []int{value}
			return true
		}
	}
	return false
}

// Clone implements Domain.
func (d *EnumeratedDomain) Clone() Domain {
	values := make([]int, len(d.Values))
	copy(values, d.Values)
	return &EnumeratedDomain{
		Values: values,
	}
}

func newEnumeratedDomain(domSize int) *EnumeratedDomain {
	domain := make([]int, domSize)
	for i := 0; i < domSize; i++ {
		domain[i] = i
	}
	return &EnumeratedDomain{
		Values: domain,
	}
}

// Fixed implements Domain.
func (d EnumeratedDomain) Fixed() bool {
	return len(d.Values) == 1
}

// Max implements Domain.
func (d EnumeratedDomain) Max() int {
	if d.Empty() {
		log.Error("Empty domain in Max")
		os.Exit(1)
	}

	maxVal := d.Values[0]
	for _, val := range d.Values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// Min implements Domain.
func (d EnumeratedDomain) Min() int {
	if d.Empty() {
		log.Error("Empty domain in Min")
		os.Exit(1)
	}

	minVal := d.Values[0]
	for _, val := range d.Values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

// Remove implements Domain.
func (d *EnumeratedDomain) Remove(value int) bool {
	for i := 0; i < len(d.Values); i++ {
		if d.Values[i] == value {
			d.Values[i] = d.Values[len(d.Values)-1]
			d.Values = d.Values[:len(d.Values)-1]
			return true
		}
	}
	return false
}

func (d EnumeratedDomain) Empty() bool {
	return len(d.Values) == 0
}

func (d EnumeratedDomain) String() string {
	return fmt.Sprintf("%#v", d.Values)
}

func NewIntVariable(name string, domSize int) IntVariable {
	return IntVariable{
		name:   name,
		domain: newEnumeratedDomain(domSize),
	}
}

var _ Variable = &IntVariable{}

var _ Domain = &EnumeratedDomain{}
