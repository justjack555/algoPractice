package graph

import "fmt"

/**
	Vertex structure
 */
type Vertex struct {
	Val interface{}
	Name string
	Edges []*Edge
}

/**
	Edge structure 
	
**/
type Edge struct {
	v *Vertex
	weight int
}

type Path []*Edge

/**
	Graph contains a vertex
	to start at and the
	number of nodes in the
	graph
 */
type Graph struct {
	V *Vertex
}

func New() *Graph {
	g := new(Graph)
	g.V = nil
	return g
}

/**
	Add a a new Vertex to the graph.

	If the graph is empty, simply add the edge at the
	root and return

	Otherwise, add the undirected Edges to each vertex's list
	and update the graph size
 */
func (g *Graph) AddUndirectedEdge(source *Vertex, newVert *Vertex, weight int) {
	if g.V == nil {
		g.V = newVert
		return
	}
	
	newEdgeSource := &Edge {
		v : newVert,
		weight : weight,
	}
	newEdgeDest := &Edge {
		v : source,
		weight : weight,
	}
	
	source.Edges = append(source.Edges, newEdgeSource)
	newVert.Edges = append(newVert.Edges, newEdgeDest)
}

func (p *Path) Print() {
	len := 0
	fmt.Printf("Path: ")
	for _ , e := range *p {
		fmt.Printf("%s, ", e.v.Name)
		len += e.weight
	}
	fmt.Printf(" with length: %d", len)
}
