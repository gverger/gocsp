package variables

type Variable interface {
	Dom() Domain
	SetDom(Domain)
}

type Domain interface {
	Empty() bool
	Fixed() bool
	Min() int
	Max() int
	Remove(value int) bool // true if removed
	Fix(value int) bool // true if possible
	Clone() Domain
}
