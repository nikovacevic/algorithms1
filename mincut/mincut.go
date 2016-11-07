package mincut

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Vertex is a vertex in a Graph of ints
type Vertex int

// Graph is a graph representation, which maps each Vertex in the Graph
// to an array of its neighbors
type Graph map[Vertex][]Vertex

// Mincut ..
func (g *Graph) Mincut() int {
	min := maxInt
	for i := 0; i < 1000; i++ {
		f := g.duplicate()
		m := 0
		for len(*f) > 2 {
			f.contract(f.randEdge())
		}
		for _, edges := range *f {
			m += len(edges)
		}
		if m < min {
			min = m
		}
	}
	return min
}

// contract merges (or contracts) two Vertices in a Graph
func (g *Graph) contract(va Vertex, vb Vertex) error {
	// Contract edges into va
	(*g)[va] = append((*g)[va], (*g)[vb]...)
	// Delete vb from the graph
	delete((*g), vb)
	// Delete edges between va and vb
	for i := 0; i < len((*g)[va]); {
		if (*g)[va][i] == vb || (*g)[va][i] == va {
			if i == len((*g)[va])-1 {
				(*g)[va] = (*g)[va][:i]
			} else {
				(*g)[va] = append((*g)[va][:i], (*g)[va][i+1:]...)
			}
		} else {
			i++
		}
	}
	// Replace all instances of vb with va in other edges
	for v, neighbors := range *g {
		for i, n := range neighbors {
			if n == vb {
				(*g)[v][i] = va
			}
		}
	}
	return nil
}

func (g *Graph) duplicate() *Graph {
	f := Graph{}
	for v, e := range *g {
		f[v] = make([]Vertex, len(e))
		copy(f[v], e)
	}
	return &f
}

// randEdge selects a random starting Vertex, then a random ending Vertex
// from the starting Vertex, returning the pair
func (g *Graph) randEdge() (start Vertex, end Vertex) {
	i := 0
	n := r.Intn(len(*g))
	for s := range *g {
		if i < n {
			i++
			continue
		}
		start = s
	}
	end = (*g)[start][r.Intn(len((*g)[start]))]
	return
}

// String renders a Graph as a string
func (g *Graph) String() string {
	str := ""
	for vertex, neighbors := range *g {
		str += fmt.Sprintf("%v: %v\n", vertex, neighbors)
	}
	return str
}

// NewGraphFromFile takes a file path, read in the file as an adjacency list
// representation of a Graph, and returns the generated *Graph.
func NewGraphFromFile(path string) (*Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	G := Graph{}
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "\t")
		v, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}
		vertex := Vertex(v)
		neighbors := []Vertex{}
		values = values[1:]
		for _, v := range values {
			if len(v) == 0 {
				continue
			}
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			neighbors = append(neighbors, Vertex(n))
		}
		G[vertex] = neighbors
	}
	return &G, nil
}
