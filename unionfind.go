package unionfind

type Item interface {
	~uint32 | ~uint64
}

func Add[T Item](set []T, i T) {
	if !Contains(set, i) {
		set[i] = i
	}
}

func Contains[T Item](set []T, i T) bool {
	return set[i] != 0
}

func Union[T Item](set []T, A, B T) {
	rootA := Find(set, A)
	rootB := Find(set, B)

	switch {
	case rootA < rootB:
		set[rootB] = rootA
	case rootA > rootB:
		set[rootA] = rootB
	}
}

func Find[T Item](set []T, i T) T {
	// use the path splitting approach
	for {
		parent := set[i]
		if parent == i {
			return parent
		}
		grandparent := set[parent]
		set[i] = grandparent
		if grandparent == i {
			return grandparent
		}
		i = grandparent
	}
}
