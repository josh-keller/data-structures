package binaryTree

import (
	"fmt"
	"testing"
)

func TestMakeTree(t *testing.T) {
	got := MakeTree()
	want := NodeTree{}

	tree, ok := got.(*NodeTree)

	if !ok || *tree != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestInsertEmpty(t *testing.T) {
	tree := NodeTree{nil}
	tree.Insert(5)
	fmt.Println(tree.root)
	got := tree.root.val
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestInsertValLessThanRoot(t *testing.T) {
	tree := NodeTree{nil}
	tree.Insert(5)
	tree.Insert(3)
	got := tree.root.left.val
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestInsertValGreaterThanRoot(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(5)
	tree.Insert(7)
	got := tree.root.right.val
	want := 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestInsertTwoLevels(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(4).Insert(2).Insert(1).Insert(3).Insert(6).Insert(5).Insert(7)

	got := tree.root.left.left.val
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = tree.root.left.right.val
	want = 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = tree.root.right.val
	want = 6

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = tree.root.right.left.val
	want = 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = tree.root.right.right.val
	want = 7

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestSearch(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(4).Insert(2).Insert(1).Insert(3).Insert(6).Insert(5).Insert(7)

	got := tree.Search(3)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(1)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(7)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(8)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(0)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestSuccessorNode(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(50)
	tree.Insert(25).Insert(75)
	tree.Insert(10).Insert(33).Insert(56).Insert(89)
	tree.Insert(4).Insert(11).Insert(30).Insert(40)
	tree.Insert(52).Insert(61).Insert(82).Insert(95)

	got := tree.root.search(56).successor()
	want := tree.root.search(61)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = tree.root.successor()
	want = tree.root.search(52)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = tree.root.search(25).successor()
	want = tree.root.search(30)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	tree.Insert(55)

	got = tree.root.successor()
	want = tree.root.search(52)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDelete(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(50)
	tree.Insert(25).Insert(75)
	tree.Insert(10).Insert(33).Insert(56).Insert(89)
	tree.Insert(4).Insert(11).Insert(30).Insert(40)
	tree.Insert(52).Insert(61).Insert(82).Insert(95)

	tree.Delete(4)
	got := tree.Search(4)
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(10)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(11)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	tree.Delete(10)
	got = tree.Search(11)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(10)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	tree.Delete(56)
	got = tree.Search(56)
	want = false

	got = tree.Search(61)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = tree.Search(52)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	tree.Insert(55)
	tree.Delete(50)

	got = tree.Search(50)
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	gotVal := tree.root.val
	wantVal := 52

	if got != want {
		t.Errorf("got %d, wanted %d", gotVal, wantVal)
	}

	got = tree.Search(55)
	want = true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestTraverse(t *testing.T) {
	tree := NodeTree{}
	tree.Insert(50)
	tree.Insert(25).Insert(75)
	tree.Insert(10).Insert(33).Insert(56).Insert(89)
	tree.Insert(4).Insert(11).Insert(30).Insert(40)
	tree.Insert(52).Insert(61).Insert(82).Insert(95)

	tree.Traverse(func(v int) {
		fmt.Println(v)
	})
}
