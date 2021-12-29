package binaryTree

import "fmt"

type Tree interface {
	Insert(v int) Tree
	Delete(v int)
	Search(v int) bool
	Traverse(f func(val int))
}

type node struct {
	val   int
	left  *node
	right *node
}

type nodeTree struct {
	root *node
}

func MakeTree() Tree {
	return &nodeTree{nil}
}

func (n *node) insert(v int) {
	newNode := node{v, nil, nil}

	if v < n.val {
		if n.left == nil {
			n.left = &newNode
		} else {
			n.left.insert(v)
		}
	} else if v > n.val {
		if n.right == nil {
			n.right = &newNode
		} else {
			n.right.insert(v)
		}
	}
}

func (n *node) delete(val int) *node {
	if n == nil {
		fmt.Println("val to delete not found")
		return nil
	} else if val < n.val {
		n.left = n.left.delete(val)
		return n
	} else if val > n.val {
		n.right = n.right.delete(val)
		return n
  } else {
    if n.left == nil {
      return n.right
    } else if n.right == nil {
      return n.left
    } else {
      n.right = lift(n.right, n)
      return n
    }
  }
}

func (n *node) traverseLeft() *node {
	if n.left == nil {
		return n
	} else {
		return n.left.traverseLeft()
	}
}

func (n *node) traverseRight() *node {
	if n.right == nil {
		return n
	} else {
		return n.right.traverseRight()
	}
}

func (n *node) successor() *node {
	if n == nil {
		panic("No successor of nil node")
	}

	if n.left == nil && n.right == nil {
		panic("No successor of node with no children")
	}

	if n.right == nil {
		return n.left
	}

	return n.right.traverseLeft()
}

func (n *node) search(v int) *node {
	if n == nil || v == n.val {
		return n
	} else if v < n.val {
		return n.left.search(v)
	} else {
		return n.right.search(v)
	}
}

func (t *nodeTree) Delete(val int) {
	t.root.delete(val)
}

func lift(n, toDelete *node) *node {
	if n == nil {
		return nil
	} else if n.left != nil {
		n.left = lift(n.left, toDelete)
		return n
	} else {
		toDelete.val = n.val
		return n.right
	}
}

func (t nodeTree) Traverse(f func(val int)) {
	t.root.traverse(f)
}

func (n *node) traverse(f func(val int)) {
	if n == nil {
		return
	}

	n.left.traverse(f)
	f(n.val)
	n.right.traverse(f)
}


func (t *nodeTree) Search(v int) bool {
	n := t.root.search(v)
	return n != nil
}

func (t *nodeTree) Insert(v int) Tree {
	if t.root == nil {
		t.root = &node{v, nil, nil}
	} else {
		t.root.insert(v)
	}

	return t
}
