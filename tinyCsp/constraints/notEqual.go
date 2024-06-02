package constraints

import (
	"errors"

	"gverger.com/csp/tinyCsp/variables"
)

type notEqual struct {
	v1     variables.Variable
	v2     variables.Variable
	offset int
}

// IsDone implements Constraint.
func (n notEqual) IsDone() bool {
	return n.v1.Dom().Fixed() || n.v2.Dom().Fixed()
}

var errEmptyDomain = errors.New("empty domain")

// Propagate implements Constraint.
// Todo: error to be implemented
func (n notEqual) Propagate() (bool, error) {
	if n.v1.Dom().Fixed() && n.v2.Dom().Fixed() && n.v1.Dom().Min() == n.v2.Dom().Min()+n.offset {
		return true, errEmptyDomain
	}

	if n.v1.Dom().Fixed() {
		return n.v2.Dom().Remove(n.v1.Dom().Min() - n.offset), nil
	}

	if n.v2.Dom().Fixed() {
		return n.v1.Dom().Remove(n.v2.Dom().Min() + n.offset), nil
	}

	return false, nil
}

func NotEqual(v1, v2 variables.Variable) notEqual {
	return notEqual{
		v1:     v1,
		v2:     v2,
		offset: 0,
	}
}

func NotEqualOffset(v1, v2 variables.Variable, offset int) notEqual {
	return notEqual{
		v1:     v1,
		v2:     v2,
		offset: offset,
	}
}

var _ Constraint = notEqual{}
