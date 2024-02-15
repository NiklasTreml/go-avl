package main

import (
	"avl/tree"
	"fmt"
	"math/rand"
)

func main() {
	var t *tree.Tree
	values := []int{1, 2, 3, 4, 5, 6, 7}
	// values := NValues(100)

	for i, v := range values {
		fmt.Printf("######################## Iteration %d ##########################################\n", i)
		t = tree.InsertAVL(t, v)
		tree.PrintTree(t)
		path := tree.DFS(t)
		fmt.Println("In order DFS Path: ", path)
	}
	fmt.Println("######################## Deletion #############################################")
	t = tree.DeleteAVL(t, 4)
	tree.PrintTree(t)
	path := tree.DFS(t)
	fmt.Println("In order DFS Path: ", path)

}

func NValues(n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, rand.Intn(200)-100)
	}
	return result
}
