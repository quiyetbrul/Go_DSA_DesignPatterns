package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (n *Node) search(value int) bool {
	if n == nil {
		return false
	}

	if value == n.value {
		return true
	}

	if value < n.value {
		if n.left == nil {
			return false
		}
		return n.left.search(value)
	}

	if n.right == nil {
		return false
	}

	return n.right.search(value)
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.root == nil {
		t.root = &Node{value: value}
	} else {
		t.root.insert(value)
	}

	return t
}

func printNode(n *Node) {
	if n == nil {
		return
	}

	printNode(n.left)
	println(n.value)
	printNode(n.right)
}

func main() {
	t := &Tree{}
	t.Insert(10)
	t.Insert(20)
	t.Insert(5)
	t.Insert(15).Insert(30).
		Insert(25).
		Insert(35)
	printNode(t.root)
	println(t.root.search(30))
	println(t.root.search(100))
}
