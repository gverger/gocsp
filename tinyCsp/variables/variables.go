package variables

type BoundConstraint interface{
	IsDone() bool
}

type Variable interface {
	Name() string
	Dom() Domain
	SetDom(Domain)
	NbConstraints() int
	ConstraintAdded(BoundConstraint)
	Constraints()
}

type Domain interface {
	Empty() bool
	Size() int
	Fixed() bool
	Min() int
	Remove(value int) bool // true if removed
	Fix(value int) bool // true if possible
	Clone() Domain
}
