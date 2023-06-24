package progpuzzlib

// Tree represents a binary tree in which each node contains a
// single generic value
type Tree struct {
  Value interface{}
  Left *Tree
  Right *Tree
}

// FindNextUntraversedPath either returns a path with at least one untraversed node
// or returns the empty path.  Assuming we are traversing left to right.
func FindNextUntraversedPath(currPath []*Tree, currSum int) ([]*Tree, int) {
    // Handle the empty tree edge case
    if len(currPath) == 0 {
        return currPath, 0
    }
    // Otherwise there must be at least one node
    lastNode := currPath[len(currPath)-1]
    // Try descending
    if lastNode.Left != nil {
        return append(currPath, lastNode.Left), currSum + lastNode.Left.Value.(int)
    }
    if lastNode.Right != nil {
        return append(currPath, lastNode.Right), currSum + lastNode.Right.Value.(int)
    }
    // Backtrack
    for {
        // Pre-conditions:
        // 1. lastNode is a dead-end
        // 2. len(currPath) > 0 (since it must at least contain lastNode)
        currPath = currPath[0:len(currPath)-1]
        currSum -= lastNode.Value.(int)
        if len(currPath) == 0 {
            // This loop must finish at some point because eventually we run out of tree nodes on the path
            return currPath, 0
        }
        parentNode := currPath[len(currPath)-1]
        if parentNode.Left == lastNode && parentNode.Right != nil {
            // We found a node that we haven't searched before!
            return append(currPath, parentNode.Right), currSum + parentNode.Right.Value.(int)
        }
        // Gotta keep going
        lastNode = parentNode
    }
}
