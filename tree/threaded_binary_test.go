package tree

import (
	"fmt"
	"testing"
)

func TestThreadedBinary_PreThreading(t *testing.T) {
	fmt.Println(t.Name())
	tb := newThreadedBinary()

	tb.PreThreading()

	fmt.Println(tb.PreThreadedTraverse().string())
}

func TestThreadedBinary_InThreading(t *testing.T) {
	fmt.Println(t.Name())
	tb := newThreadedBinary()

	tb.InThreading()

	fmt.Println(tb.InThreadedTraverse().string())
}

func TestThreadedBinary_PostThreading(t *testing.T) {
	fmt.Println(t.Name())
	tb := newThreadedBinary()

	tb.PostThreading()

	fmt.Println(tb.PostThreadedTraverse().string())
}

func newThreadedBinary() *threadedBinary {
	node11 := &threadedNode{
		data:       11,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: nil,
	}

	node10 := &threadedNode{
		data:       10,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: nil,
	}

	node9 := &threadedNode{
		data:       9,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: nil,
	}

	node8 := &threadedNode{
		data:       8,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: node11,
	}

	node7 := &threadedNode{
		data:       7,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: node10,
	}

	node6 := &threadedNode{
		data:       6,
		leftTag:    false,
		rightTag:   false,
		leftChild:  node9,
		rightChild: nil,
	}

	node5 := &threadedNode{
		data:       5,
		leftTag:    false,
		rightTag:   false,
		leftChild:  nil,
		rightChild: nil,
	}

	node4 := &threadedNode{
		data:       4,
		leftTag:    false,
		rightTag:   false,
		leftChild:  node8,
		rightChild: nil,
	}

	node3 := &threadedNode{
		data:       3,
		leftTag:    false,
		rightTag:   false,
		leftChild:  node6,
		rightChild: node7,
	}

	node2 := &threadedNode{
		data:       2,
		leftTag:    false,
		rightTag:   false,
		leftChild:  node4,
		rightChild: node5,
	}

	node1 := &threadedNode{
		data:       1,
		leftTag:    false,
		rightTag:   false,
		leftChild:  node2,
		rightChild: node3,
	}

	return &threadedBinary{head: node1}
}
