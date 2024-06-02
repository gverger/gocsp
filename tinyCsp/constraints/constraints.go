package constraints

type Constraint interface {
	Propagate() (bool, error) // true if changes have been made, error if inconsistent
	IsDone() bool
}
