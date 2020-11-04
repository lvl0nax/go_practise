package main

import "fmt"

// BinaryTree node conatins value and pointer to the left and right node
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// BuildTree simple function to build a binary tree from the list of numbers
func BuildTree(values []int) *BinaryTree {
	rootNode := &BinaryTree{Value: values[0]}
	values = values[1:]

	nodes := []*BinaryTree{rootNode}

	for _, value := range values {
		newNode := &BinaryTree{Value: value}
		node := nodes[0]
		if node.Left == nil {
			node.Left = newNode
			nodes = append(nodes, newNode)
		} else if node.Right == nil {
			node.Right = newNode
			nodes = nodes[1:]
			nodes = append(nodes, newNode)
		}
	}

	return rootNode
}

// BranchSums recursive function which calculates branch sums
func BranchSums(node *BinaryTree, sum int, result []int) []int {
	sum += node.Value

	if node.Left != nil {
		result = BranchSums(node.Left, sum, result)
	}
	if node.Right != nil {
		result = BranchSums(node.Right, sum, result)
	}
	if node.Left == nil && node.Right == nil {
		result = append(result, sum)
	}
	return result
}

// BuildTreeAndCalculate takes list of numbers and returns list of branch sums
func BuildTreeAndCalculate(values ...int) []int {
	root := BuildTree(values)

	var result []int

	return BranchSums(root, 0, result)
}

func main() {
	result := BuildTreeAndCalculate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(result)
}
