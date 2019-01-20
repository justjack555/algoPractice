package main

import (
	"log"

	"github.com/justjack555/algoPractice/pkg/graph"
	"github.com/justjack555/algoPractice/pkg/dijkstra"

)

func main(){
	log.Println("Doing Dijkstra...")
	g := graph.New()
	log.Println("MAIN(): New graph has initial vertex: ", g.V)

	g.AddUndirectedEdge(g.V, &graph.Vertex{
		Val : 0,
		Name : "1",
		Edges : make([]*graph.Edge, 0, 0),
	}, 3)

	paths, err := dijkstra.Dijkstra(g)
	if err != nil {
		log.Fatal("ERR: Dijkstra() failed with error: ", err)
	}

	for _ , path := range paths {
		path.Print()
	}

}
