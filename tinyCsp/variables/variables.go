package variables

type Variable interface {
	Name() string
	Dom() Domain
	SetDom(Domain)
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
