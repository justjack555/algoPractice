package graph

/**
	Vertex structure
 */
type Vertex struct {
	val interface{}
	edges []*Edge
}

/**
	Edge structure 
	
**/
type Edge struct {
	v *Vertex
	weight int
}

type Path []*Edge

type Graph *Vertex
