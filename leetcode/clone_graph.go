package leetcode

import "github.com/emirpasic/gods/v2/maps/hashmap"

type CloneGraphNode struct {
	Val       int
	Neighbors []*CloneGraphNode
}

// Medium: https://leetcode.com/problems/clone-graph
func CloneGraph(root *CloneGraphNode) *CloneGraphNode {
	if root == nil {
		return nil
	}
	nodes := hashmap.New[int, *CloneGraphNode]()
	createNodes(root, nodes)
	visited := hashmap.New[int, bool]()
	visited.Put(root.Val, true)
	return clone(root, nodes, visited)
}

func clone(root *CloneGraphNode, nodes *hashmap.Map[int, *CloneGraphNode], visited *hashmap.Map[int, bool]) *CloneGraphNode {
	if root == nil {
		return nil
	}

	currNode, _ := nodes.Get(root.Val)
	for _, neighbor := range root.Neighbors {
		nbrNode, _ := nodes.Get(neighbor.Val)
		currNode.Neighbors = append(currNode.Neighbors, nbrNode)
	}

	for _, neighbor := range root.Neighbors {
		value, _ := visited.Get(neighbor.Val)
		if !value {
			visited.Put(neighbor.Val, true)
			clone(neighbor, nodes, visited)
		}
	}

	return currNode
}

func createNodes(root *CloneGraphNode, nodes *hashmap.Map[int, *CloneGraphNode]) {
	if root == nil {
		return
	}

	_, found := nodes.Get(root.Val)
	if found {
		return
	}

	newNode := &CloneGraphNode{
		Val:       root.Val,
		Neighbors: make([]*CloneGraphNode, 0),
	}

	nodes.Put(newNode.Val, newNode)

	for _, neighbor := range root.Neighbors {
		createNodes(neighbor, nodes)
	}
}

func CloneGraphV2(root *CloneGraphNode) *CloneGraphNode {
	if root == nil {
		return nil
	}
	nodeMap := hashmap.New[int, *CloneGraphNode]()
	return cloneDfs(root, nodeMap)
}

func cloneDfs(root *CloneGraphNode, nodeMap *hashmap.Map[int, *CloneGraphNode]) *CloneGraphNode {
	if clonedNode, found := nodeMap.Get(root.Val); found {
		return clonedNode
	}

	clonedNode := &CloneGraphNode{
		Val:       root.Val,
		Neighbors: make([]*CloneGraphNode, 0),
	}
	nodeMap.Put(clonedNode.Val, clonedNode)

	for _, neighbor := range root.Neighbors {
		clonedNeighbour := cloneDfs(neighbor, nodeMap)
		clonedNode.Neighbors = append(clonedNode.Neighbors, clonedNeighbour)
	}

	return clonedNode
}
