package medium

import "math"

/**
https://leetcode.com/problems/most-profitable-path-in-a-tree/description/
2467. Most Profitable Path in a Tree

There is an undirected tree with n nodes labeled from 0 to n - 1, rooted at node 0. You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

At every node i, there is a gate. You are also given an array of even integers amount, where amount[i] represents:

the price needed to open the gate at node i, if amount[i] is negative, or,
the cash reward obtained on opening the gate at node i, otherwise.
The game goes on as follows:

Initially, Alice is at node 0 and Bob is at node bob.
At every second, Alice and Bob each move to an adjacent node. Alice moves towards some leaf node, while Bob moves towards node 0.
For every node along their path, Alice and Bob either spend money to open the gate at that node, or accept the reward. Note that:
If the gate is already open, no price will be required, nor will there be any cash reward.
If Alice and Bob reach the node simultaneously, they share the price/reward for opening the gate there. In other words, if the price to open the gate is c, then both Alice and Bob pay c / 2 each. Similarly, if the reward at the gate is c, both of them receive c / 2 each.
If Alice reaches a leaf node, she stops moving. Similarly, if Bob reaches node 0, he stops moving. Note that these events are independent of each other.
Return the maximum net income Alice can have if she travels towards the optimal leaf node.

*/

func MostProfitablePath(edges [][]int, bob int, amount []int) int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	bobNodeTimeMap := make(map[int]int)
	visited := make([]bool, len(edges)+1)
	// DFS from bob to 0 node, collect the time at which bob reached each node in its path
	bobDfs(bob, 0, bobNodeTimeMap, visited, adjList)

	maxAmount := math.MinInt

	// reset visited array for BFS
	for i := range visited {
		visited[i] = false
	}

	visited[0] = true
	// (node, time, amount)
	aliceBfsQueue := [][]int{{0, 0, 0}}

	for len(aliceBfsQueue) > 0 {
		front := aliceBfsQueue[0]
		aliceBfsQueue = aliceBfsQueue[1:]

		currNode := front[0]
		currTime := front[1]
		currAmount := front[2]

		amt := currAmount
		// check time when node was visited by bob
		bobTime, ok := bobNodeTimeMap[currNode]
		if !ok {
			// not visited by bob
			amt += amount[currNode]
		} else if currTime < bobTime {
			// alice visited first
			amt += amount[currNode]
		} else if currTime == bobTime {
			// both visit this node together
			amt += amount[currNode] / 2
		} else {
			// bob visit first
			amt += 0
		}

		// if curr node is a leaf node
		if len(adjList[currNode]) == 1 && currNode != 0 {
			maxAmount = max(maxAmount, amt)
		}

		for _, nbr := range adjList[currNode] {
			if !visited[nbr] {
				aliceBfsQueue = append(aliceBfsQueue, []int{nbr, currTime + 1, amt})
			}
		}

		// mark and visit the curr node
		visited[currNode] = true
	}

	return maxAmount
}

func bobDfs(curr int, time int, nodeTimeMap map[int]int, visited []bool, adjList map[int][]int) bool {
	// base cond
	nodeTimeMap[curr] = time
	visited[curr] = true
	if curr == 0 {
		return true
	}

	for _, nbr := range adjList[curr] {
		if !visited[nbr] {
			if bobDfs(nbr, time+1, nodeTimeMap, visited, adjList) {
				return true
			}
		}
	}

	// backtrack
	delete(nodeTimeMap, curr)
	return false
}
