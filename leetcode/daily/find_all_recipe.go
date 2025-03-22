package leetcode

import (
	"container/list"
	"github.com/emirpasic/gods/v2/sets/hashset"
)

// Medium: https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies
func FindAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesSet := hashset.New[string]()
	for _, supply := range supplies {
		suppliesSet.Add(supply)
	}

	recipeQue := make([]int, 0)
	for i := range recipes {
		recipeQue = append(recipeQue, i)
	}

	lastSuppliesSize := -1
	var ans []string
	for suppliesSet.Size() > lastSuppliesSize {
		lastSuppliesSize = suppliesSet.Size()
		tempQueue := make([]int, 0)

		for _, recipeInd := range recipeQue {
			recipe := recipes[recipeInd]
			recipeIngredients := ingredients[recipeInd]
			canMake := true

			for _, ingredient := range recipeIngredients {
				if !suppliesSet.Contains(ingredient) {
					canMake = false
					break
				}
			}

			if canMake {
				ans = append(ans, recipe)
				suppliesSet.Add(recipe)
			} else {
				tempQueue = append(tempQueue, recipeInd)
			}
		}

		recipeQue = tempQueue
	}

	return ans

}

// Topological Sort
func FindAllRecipesV2(recipes []string, ingredients [][]string, supplies []string) []string {
	suppliesSet := hashset.New[string]()
	for _, supply := range supplies {
		suppliesSet.Add(supply)
	}

	recipeIndexMap := make(map[string]int)
	for i, recipe := range recipes {
		recipeIndexMap[recipe] = i
	}

	indegree := make([]int, len(recipes))
	dependencyGraph := make(map[string][]string)

	for i := 0; i < len(recipes); i++ {
		for _, ingredient := range ingredients[i] {
			if !suppliesSet.Contains(ingredient) {
				// Add edge: ingredient -> recipe
				dependencyGraph[ingredient] = append(dependencyGraph[ingredient], recipes[i])
				indegree[i]++
			}
		}
	}

	queue := list.New()
	for i, deg := range indegree {
		if deg == 0 {
			queue.PushBack(i)
		}
	}

	var ans []string
	for queue.Len() > 0 {
		recipeIndex := queue.Remove(queue.Front()).(int)
		recipe := recipes[recipeIndex]
		ans = append(ans, recipe)

		for _, rc := range dependencyGraph[recipe] {
			indegree[recipeIndexMap[rc]]--

			if indegree[recipeIndexMap[rc]] == 0 {
				queue.PushBack(recipeIndexMap[rc])
			}
		}
	}

	return ans
}
