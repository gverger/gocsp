package graphcoloring

import (
	"fmt"
	"path"
	"testing"

	"github.com/matryer/is"
)

var data = "../../../data/graph_coloring"

func TestGraphColoring(t *testing.T) {

	filenames := make([]string, 10)
	for i := 0; i < len(filenames); i++ {
		filenames[i] = path.Join(data, fmt.Sprintf("gc_15_30_%d", i))
	}

	for _, filename := range filenames {
		t.Run(filename, func(t *testing.T) {
			is := is.New(t)
			instance := readFile(filename)

			solution, found := Solve(instance)

			is.True(found)                            // solution was found
			is.Equal(len(solution), instance.nbNodes) // solution nb of nodes

			for _, e := range instance.edges {
				is.True(solution[e.x] != solution[e.y]) // adjacent nodes have the same color
			}
		})
	}

}
