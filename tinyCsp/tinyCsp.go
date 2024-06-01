package tinycsp

import (
	log "log/slog"

	"gverger.com/csp/tinyCsp/constraints"
	"gverger.com/csp/tinyCsp/variables"
)

type TinyCsp struct {
	Variables   []variables.Variable
	Constraints []constraints.Constraint
}

func NewTinyCsp() TinyCsp {

	variables := make([]variables.Variable, 0)
	constraints := make([]constraints.Constraint, 0)

	return TinyCsp{
		Variables:   variables,
		Constraints: constraints,
	}
}

func (csp TinyCsp) Solve(onSolution func(solution []int)) {
	log.Debug("Selecting variable")
	variable, foundBranchingVar := csp.selectVariable()
	if !foundBranchingVar {
		exportSolution(csp.Variables, onSolution)
		return
	}

	log.Debug("Chosen %s", variable.Name())
	val := variable.Dom().Min()

	backup := csp.backupDomains()

	log.Debug("%s <- %d", variable.Name(), val)
	if variable.Dom().Fix(val) {
		log.Debug("Fixing point")
		err := csp.fixPoint()
		if err == nil {
			log.Debug("Go deeper")
			csp.Solve(onSolution)
		}
	}

	log.Debug("Restore")
	csp.restoreDomains(backup)

	log.Debug("Remove Value %d for %s", val, variable.Name())
	variable.Dom().Remove(val)
	if variable.Dom().Empty() {
		log.Debug("Domain is empty")
		return
	}
	if err := csp.fixPoint(); err != nil {
		log.Debug("Fix point inconsistent")
		return
	}
	csp.Solve(onSolution)
}

func exportSolution(variables []variables.Variable, callback func(solution []int)) {
	solution := make([]int, len(variables))
	for i := 0; i < len(variables); i++ {
		solution[i] = variables[i].Dom().Min()
	}
	callback(solution)
}

func (csp TinyCsp) selectVariable() (variables.Variable, bool) {
	for _, variable := range csp.Variables {
		if !variable.Dom().Fixed() {
			return variable, true
		}
	}
	return nil, false
}

func (csp *TinyCsp) restoreDomains(backup []variables.Domain) {
	for i := 0; i < len(csp.Variables); i++ {
		csp.Variables[i].SetDom(backup[i])
	}
}

func (csp TinyCsp) backupDomains() []variables.Domain {
	backup := make([]variables.Domain, 0, len(csp.Variables))

	for i := 0; i < len(csp.Variables); i++ {
		backup = append(backup, csp.Variables[i].Dom().Clone())
	}

	return backup
}

func (csp *TinyCsp) fixPoint() error {
	fix := false

	for !fix {
		fix = true
		for _, c := range csp.Constraints {
			changed, err := c.Propagate()
			if err != nil {
				return err
			}
			fix = !changed && fix
		}
	}

	return nil
}

// upperBound is not included in the domain
func (csp *TinyCsp) MakeVariable(name string, domSize int) variables.Variable {
	v := variables.NewIntVariable(name, domSize)

	csp.Variables = append(csp.Variables, &v)

	return &v
}

// v1 != v2 + offset
func (csp *TinyCsp) NotEqual(v1, v2 variables.Variable, offset int) constraints.Constraint {
	c := constraints.NotEqualOffset(v1, v2, offset)

	csp.Constraints = append(csp.Constraints, c)

	return c
}
