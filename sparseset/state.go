package sparseset

type SetState struct {
	min  int
	max  int
	size int
}

func (s Set) SaveState() SetState {
	return SetState{
		min:  s.min,
		max:  s.max,
		size: s.Size,
	}
}

func (s *Set) RestoreState(state SetState) {
	s.min = state.min
	s.max = state.max
	s.Size = state.size
}
