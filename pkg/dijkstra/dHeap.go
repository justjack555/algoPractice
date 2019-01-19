package dijkstra

import (
	"log"
	"container/heap"
)


// A PriorityQueue implements heap.Interface and holds Dijkstra
// vertices.
type dQueue []*DVertex

func (dq dQueue) Len() int { return len(dq) }

func (dq dQueue) Less(i, j int) bool {
	return dq[i].dist < dq[j].dist
}

func (dq dQueue) Swap(i, j int) {
	dq[i], dq[j] = dq[j], dq[i]
	dq[i].index = i
	dq[j].index = j
}

func (dq *dQueue) Push(x interface{}) {
	n := len(*dq)
	item := x.(*DVertex)
	item.index = n
	*dq = append(*dq, item)
}

func (dq *dQueue) Pop() interface{} {
	old := *dq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*dq = old[0 : n-1]
	return item
}

func (dq *dQueue) update(dv *DVertex, newDist int) {
	dv.dist = newDist
	heap.Fix(dq, dv.index)
}
