package dijkstra

import (
	"log"
	"github.com/justjack555/algoPractice/pkg/graph"
)

/**
	Store:
	1. Pointer to the vertex itself
	2. The distance from the vertex to source vertex
	3. The edge containing the previous vertex en route to
	the source
	4. Index position in the heap
**/
type DVertex struct {
	v *Vertex
	dist int
	prev *Edge
	index int
}

/**
	Dijkstra returns the shortest path for
	each vertex to the source.
	
	Assumption is that source is the vertex pointed
	to by g
	
	Setup adds items to the heap, then they are removed,
	and distances/previous vertices on path are updated
	if applicable
**/
func (g *Graph) Dijkstra() ([]Path, error) {

} 
