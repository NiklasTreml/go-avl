package tree

import (
	"fmt"
	"reflect"
)

type TreePrinter interface {
	Left() TreePrinter
	Right() TreePrinter
	Value() int
}

func printTree(t TreePrinter, indent string, last bool) {
	if t != nil && !reflect.ValueOf(t).IsNil() {
		fmt.Print(indent)
		if last {
			fmt.Printf("L----")
			indent += "      "
		} else {
			fmt.Printf("R----")
			indent += "|   "
		}

		fmt.Println(t.Value())
		printTree(t.Right(), indent, false)
		printTree(t.Left(), indent, true)
	}
}

func PrintTree(t TreePrinter) {
	printTree(t, "", true)
}

func DFS(t TreePrinter) []int {
	return dfs(t, []int{})
}

func dfs(t TreePrinter, path []int) []int {
	if t == nil || reflect.ValueOf(t).IsNil() {
		return path
	}
	path = dfs(t.Left(), path)
	path = append(path, t.Value())
	path = dfs(t.Right(), path)
	return path
}
