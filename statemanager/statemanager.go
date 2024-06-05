package statemanager

type StateInt struct {
	Value int
}

type State struct {
	ints []int
}

type StateManager struct {
	stateInts []*StateInt
}

func New() *StateManager {
	return &StateManager{
		stateInts: make([]*StateInt, 0),
	}
}

func (m *StateManager) NewInt(value int) *StateInt {
	si := &StateInt{Value: value}
	m.stateInts = append(m.stateInts, si)

	return si
}

func (m StateManager) Save() State {
	ints := make([]int, len(m.stateInts))
	for i, s := range m.stateInts {
		ints[i] = s.Value
	}

	return State{
		ints: ints,
	}
}

func (m *StateManager) Restore(s State) {
	for i, value := range s.ints {
		m.stateInts[i].Value = value
	}
}
