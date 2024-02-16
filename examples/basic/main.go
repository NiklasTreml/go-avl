package main

import (
	"avl/tree"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// This example shows how an avl tree stays balanced even when inserting an array of numbers that only increases
// which, when using a normal binary, would be considered a worst case scenario, since it essentially just creates
// a fancy LinkedList
// The AVL tree prevents this from happening by rebalancing the tree on every insertion while still ensuring the 
// requirements for a Binary Search Tree are satisfied. This keeps lookup times low, 
// even in a suboptimal scenario like this one.
func main() {
	var t *tree.Tree
	values := NValues(25)
	delay := 250 * time.Millisecond
	clearScreen()

	// Insert random values
	for i, v := range values {
		t = tree.InsertAVL(t, v)
		clearScreen()
		fmt.Printf("######################## Iteration %d ##########################################\n", i)
		tree.PrintTree(t)
		path := tree.DFS(t)
		fmt.Println("In order DFS Path: ", path)
		time.Sleep(delay)
	}

	time.Sleep(delay * 3)
	// Delete the same items in random order
	shuffled := shuffle(values)
	for _, v := range shuffled {
		t = tree.DeleteAVL(t, v)
		clearScreen()
		fmt.Printf("######################## Delete %d ###########################################\n", v)
		tree.PrintTree(t)
		path := tree.DFS(t)
		fmt.Println("In order DFS Path: ", path)
		time.Sleep(delay)
	}

}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func NValues(n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

func shuffle(array []int) []int {
	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(len(array) - 1)
		tmp := array[i]
		array[i] = array[j]
		array[j] = tmp
	}
	return array
}
