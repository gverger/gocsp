package tinycsp

import (
	log "log/slog"
	"math"

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

type VarSelection func(csp TinyCsp) (variables.Variable, bool)

func (csp TinyCsp) Solve(selectVar VarSelection, onSolution func(solution []int) bool) bool {
	log.Debug("Selecting variable")
	variable, foundBranchingVar := selectVar(csp)
	if !foundBranchingVar {
		return exportSolution(csp.Variables, onSolution)
	}

	log.Debug("variable chosen", log.String("variable", variable.Name()))
	val := variable.Dom().Min()

	backup := csp.backupDomains()

	log.Debug("%s <- %d", variable.Name(), val)
	if variable.Dom().Fix(val) {
		log.Debug("Fixing point")
		err := csp.fixPoint()
		if err == nil {
			log.Debug("Go deeper")
			if !csp.Solve(selectVar, onSolution) {
				return false
			}
		}
	}

	log.Debug("Restore")
	csp.restoreDomains(backup)

	log.Debug("value removed", log.Int("value", val), log.String("variable", variable.Name()))
	variable.Dom().Remove(val)
	if variable.Dom().Empty() {
		log.Debug("Domain is empty")
		return true
	}
	if err := csp.fixPoint(); err != nil {
		log.Debug("Fix point inconsistent")
		return true
	}
	return csp.Solve(selectVar, onSolution)
}

func exportSolution(variables []variables.Variable, callback func(solution []int) bool) bool {
	solution := make([]int, len(variables))
	for i := 0; i < len(variables); i++ {
		solution[i] = variables[i].Dom().Min()
	}
	return callback(solution)
}

func FirstVar(csp TinyCsp) (variables.Variable, bool) {
	for _, variable := range csp.Variables {
		if !variable.Dom().Fixed() {
			return variable, true
		}
	}
	return nil, false
}

func MaxDomVar(csp TinyCsp) (variables.Variable, bool) {
	maxDom := 0
	var bestVar variables.Variable

	for _, variable := range csp.Variables {
		if variable.Dom().Fixed() {
			continue
		}
		domSize := variable.Dom().Size()
		if domSize > maxDom {
			maxDom = domSize
			bestVar = variable
		}
	}
	return bestVar, bestVar != nil
}

func MinDomVar(csp TinyCsp) (variables.Variable, bool) {
	minDom := math.MaxInt
	var bestVar variables.Variable

	for _, variable := range csp.Variables {
		if variable.Dom().Fixed() {
			continue
		}
		domSize := variable.Dom().Size()
		if domSize < minDom {
			minDom = domSize
			bestVar = variable
		}
	}
	return bestVar, bestVar != nil
}

func DomOverDegVar(csp TinyCsp) (variables.Variable, bool) {
	minDomOverDeg := math.MaxFloat32
	var bestVar variables.Variable

	for _, variable := range csp.Variables {
		if variable.Dom().Fixed() {
			continue
		}
		domOverDeg := float64(variable.Dom().Size()) / float64(variable.NbConstraints())
		if domOverDeg < minDomOverDeg {
			minDomOverDeg = domOverDeg
			bestVar = variable
		}
	}
	return bestVar, bestVar != nil
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
	v1.ConstraintAdded(c)
	v2.ConstraintAdded(c)

	return c
}
