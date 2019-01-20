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
	v *graph.Vertex
	dist int
	prev *graph.Edge
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
func Dijkstra(g *graph.Graph) ([]*graph.Path, error) {
	paths := make([]*graph.Path, 0, 0)
	newPath := make(graph.Path, 0)

//	newPath = append(newPath, g.V.Edges[0])
	paths = append(paths, &newPath)
	log.Println("DIJKSTRA(): path ", paths)
	return paths, nil
} 
