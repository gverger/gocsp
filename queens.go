package main

import (
	"strings"
)

type QueensSolver struct {
	boardSize int
	positions []int
}

func NewQueenSolver(boardSize int) QueensSolver {
	return QueensSolver{
		boardSize: boardSize,
		positions: make([]int, boardSize),
	}
}

func (s QueensSolver) Solve(onSolution func(positions []int)) {
	s.dfs(0, onSolution)
}

func (s QueensSolver) dfs(depth int, onSolution func(positions []int)) {
	if depth >= s.boardSize {
		solution := make([]int, len(s.positions))
		copy(solution, s.positions)
		onSolution(solution)
	} else {
		for i := 0; i < s.boardSize; i++ {
			s.positions[depth] = i
			if constraintHolds(s.positions, depth) {
				s.dfs(depth+1, onSolution)
			}
		}
	}
}

func constraintHolds(positions []int, decisionIdx int) bool {
	for i := 0; i < decisionIdx; i++ {
		if positions[i] == positions[decisionIdx] {
			return false
		}
		abs := positions[i] - positions[decisionIdx]
		if abs < 0 {
			abs = -abs
		}
		if abs == decisionIdx-i {
			return false
		}
	}

	return true
}

func DisplayBoard(positions []int) string {
	s := strings.Builder{}
	for _, pos := range positions {
		for i := 0; i < pos; i++ {
			s.WriteString(". ")
		}
		s.WriteString("#")
		for i := pos + 1; i < len(positions); i++ {
			s.WriteString(" .")
		}
		s.WriteString("\n")
	}

	return s.String()
}
