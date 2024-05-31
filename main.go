package main

import (
	"fmt"
	"log"

	tinycsp "gverger.com/constraints/tinyCsp"
	"gverger.com/constraints/tinyCsp/variables"
)

func specializedQueens() {
	solver := NewQueenSolver(10)

	nbSolutions := 0
	solver.Solve(func(positions []int) {
		log.Println("Solution found")
		fmt.Println(DisplayBoard(positions))
		nbSolutions++
	})

	fmt.Println("Solutions: ", nbSolutions)
}

func tinyCspQueens() {
	solver := tinycsp.NewTinyCsp()

	n := 10
	vars := make([]variables.Variable, 0, n)

	for i := 0; i < n; i++ {
		vars = append(vars, solver.MakeVariable(n))
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			solver.NotEqual(vars[i], vars[j], 0)
		}
	}

	solver.Solve(func(solution []int) {
		fmt.Print(DisplayBoard(solution))
	})
}

func main() {
	tinyCspQueens()
}
