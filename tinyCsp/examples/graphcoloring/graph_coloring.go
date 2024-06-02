package graphcoloring

import (
	"bufio"
	"fmt"
	log "log/slog"
	"os"
	"strconv"
	"strings"

	tinycsp "gverger.com/csp/tinyCsp"
	"gverger.com/csp/tinyCsp/variables"
)

type Edge struct {
	x int
	y int
}

type Instance struct {
	maxColors int
	nbNodes   int
	edges     []Edge
}

func Solve(instance Instance) ([]int, bool) {

	csp := tinycsp.NewTinyCsp()

	nodes := make([]variables.Variable, instance.nbNodes)
	for i := range nodes {
		nodes[i] = csp.MakeVariable(fmt.Sprint("x", i), instance.maxColors)
	}

	for _, e := range instance.edges {
		csp.NotEqual(nodes[e.x], nodes[e.y], 0)
	}

	var sol []int
	csp.Solve(tinycsp.FirstVar, func(solution []int) bool {
		sol = solution
		return true
	})

	return sol, true
}

func readFile(filename string) Instance {
	file, err := os.Open(filename)
	if err != nil {
		log.Error("cannot open file", log.String("file", filename), log.String("errror", err.Error()))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	header := scanner.Text()

	numbers := strings.SplitN(header, " ", 3)

	n := must(strconv.Atoi(numbers[0]))
	e := must(strconv.Atoi(numbers[1]))
	nbCol := must(strconv.Atoi(numbers[2]))

	// l.Infof("Layer 1: %d nodes", layer1Size)
	// l.Infof("Layer 2: %d nodes", layer2Size)
	// l.Infof("Edges: %d", edgesCount)

	log.Info("reading instance", log.Int("n", n), log.Int("edges", e), log.Int("colors", nbCol))

	g := Instance{
		maxColors: nbCol,
		nbNodes:   n,
		edges:     make([]Edge, 0, e),
	}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		log.Info(scanner.Text())
		nodes := strings.SplitN(scanner.Text(), " ", 2)
		a := must(strconv.Atoi(nodes[0]))
		b := must(strconv.Atoi(nodes[1]))
		g.edges = append(g.edges, Edge{a, b})
	}

	return g
}

func must[T any](value T, err error) T {
	if err != nil {
		log.Error("failure", log.String("error", err.Error()))
		panic(err)
	}
	return value
}
