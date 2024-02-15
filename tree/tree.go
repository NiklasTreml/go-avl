package tree

var _ TreePrinter = (*Tree)(nil)

type Tree struct {
	V      int
	L      *Tree
	R      *Tree
	Height int
}

func (t *Tree) Left() TreePrinter {
	return t.L
}
func (t *Tree) Right() TreePrinter {
	return t.R
}
func (t *Tree) Value() int {
	return t.V
}

func New(v int) *Tree {
	return &Tree{
		V:      v,
		L:      nil,
		R:      nil,
		Height: 1,
	}
}

func height(t *Tree) int {
	if t == nil {
		return 0
	}
	return t.Height
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getBalanceFactor(t *Tree) int {
	if t == nil {
		return 0
	}
	return height(t.L) - height(t.R)
}

func InsertAVL(t *Tree, v int) *Tree {
	if t == nil {
		return New(v)
	}
	if v < t.V {
		t.L = InsertAVL(t.L, v)
	} else if v > t.V {
		t.R = InsertAVL(t.R, v)
	} else {
		return t
	}

	t.Height = 1 + max(height(t.L), height(t.R))

	bFactor := getBalanceFactor(t)
	if bFactor > 1 {
		if v < t.L.V {
			return rightRotate(t)
		} else if v > t.L.V {
			t.L = leftRotate(t.L)
			return rightRotate(t)
		}
	}

	if bFactor < -1 {
		if v > t.R.V {
			return leftRotate(t)
		} else if v < t.R.V {
			t.R = rightRotate(t.R)
			return leftRotate(t)
		}

	}

	return t
}

func DeleteAVL(t *Tree, v int) *Tree {
	if t == nil {
		return t
	}

	if v < t.V {
		t.L = DeleteAVL(t.L, v)
	} else if v > t.V {
		t.R = DeleteAVL(t.R, v)
	} else {
		if t.L == nil || t.R == nil {
			var temp *Tree
			if t.L != nil {
				temp = t.L
			} else {
				temp = t.R
			}

			// if t is a leaf, set to nil to garbage collect
			if temp == nil {
				temp = t
				t = nil
			} else {
				// otherwise swap t with non nil child
				t = temp
			}

		} else {
			temp := nodeWithMinimumValue(t.R)
			t.V = temp.V
			t.R = DeleteAVL(t.R, temp.V)
		}
	}

	if t == nil {
		return t
	}

	// Update balance factors
	t.Height = 1 + max(height(t.L), height(t.R))
	bFactor := getBalanceFactor(t)

	if bFactor > 1 {
		if getBalanceFactor(t.L) >= 0 {
			return rightRotate(t)
		} else {
			t.L = leftRotate(t.L)
			return rightRotate(t)
		}
	}

	if bFactor < -1 {
		if getBalanceFactor(t.R) <= 0 {
			return leftRotate(t)
		} else {
			t.R = rightRotate(t.R)
			return leftRotate(t)
		}
	}

	return t

}

func nodeWithMinimumValue(node *Tree) *Tree {
	curr := node

	for curr.L != nil {
		curr = curr.L
	}
	return curr
}

func rightRotate(y *Tree) *Tree {
	x := y.L
	t2 := x.R
	x.R = y
	y.L = t2

	y.Height = max(height(y.L), height(y.R)) + 1
	x.Height = max(height(x.L), height(x.R)) + 1

	return x
}

func leftRotate(x *Tree) *Tree {
	y := x.R
	t2 := y.L
	y.L = x
	x.R = t2

	x.Height = max(height(x.L), height(x.R)) + 1
	y.Height = max(height(y.L), height(y.R)) + 1

	return y
}
