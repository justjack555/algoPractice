package graph

/**
	Vertex structure
 */
type Vertex struct {
	val interface{}
	edges []*Vertex
}

type Graph *Vertex