package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	tinycsp "gverger.com/csp/tinyCsp"
	"gverger.com/csp/tinyCsp/variables"
)

func specializedQueens(n int) {
	solver := NewQueenSolver(n)

	nbSolutions := 0
	solver.Solve(func(positions []int) {
		// log.Info("Solution found")
		// fmt.Println(DisplayBoard(positions))
		nbSolutions++
	})

	fmt.Println("Solutions: ", nbSolutions)
}

func tinyCspQueens(n int) {
	solver := tinycsp.NewTinyCsp()

	vars := make([]variables.Variable, 0, n)

	for i := 0; i < n; i++ {
		vars = append(vars, solver.MakeVariable(fmt.Sprintf("x_%d", i), n))
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			solver.NotEqual(vars[i], vars[j], 0)
			solver.NotEqual(vars[i], vars[j], j-i)
			solver.NotEqual(vars[i], vars[j], i-j)
		}
	}

	nbSolutions := 0

	solver.Solve(func(solution []int) bool {
		nbSolutions++
		return true
	})

	fmt.Println("Solutions: ", nbSolutions)
}

func main() {
	// Start profiling
	f, err := os.Create("myprogram.prof")
	if err != nil {

		fmt.Println(err)
		return

	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Run your program here
	tinyCspQueens(12)

	// specializedQueens(12)
}
