package algorithm

type DisjointSet[T comparable] struct {
	parent map[T]T
	rank   map[T]int
}

func NewDisjointSet[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		parent: make(map[T]T),
		rank:   make(map[T]int),
	}
}

func (dSet *DisjointSet[T]) MakeSet(x T) {
	dSet.parent[x] = x
	dSet.rank[x] = 0
}

func (dSet *DisjointSet[T]) Find(x T) T {
	if dSet.parent[x] != x {
		dSet.parent[x] = dSet.Find(dSet.parent[x])
	}

	return dSet.parent[x]
}

func (dSet *DisjointSet[T]) Union(x, y T) {
	rootX := dSet.Find(x)
	rootY := dSet.Find(y)

	if rootX != rootY {
		switch {
		case dSet.rank[rootX] > dSet.rank[rootY]:
			dSet.parent[rootY] = rootX
		case dSet.rank[rootX] < dSet.rank[rootY]:
			dSet.parent[rootX] = rootY
		default:
			dSet.parent[rootY] = rootX
			dSet.rank[rootX]++
		}
	}
}
