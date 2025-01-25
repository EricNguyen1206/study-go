package DSA

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 100. Same Tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

type Node struct {
	Val      int
	Children []*Node
}

// 559. Maximum Depth of N-ary Tree
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 0
	for _, child := range root.Children {
		depth := maxDepth(child)
		if depth > max {
			max = depth
		}
	}
	return max + 1
}

// 515. Find Largest Value in Each Tree Row
type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func largestValues(root *TreeNode1) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	var maxValues func(node *TreeNode1, depth int)
	maxValues = func(node *TreeNode1, depth int) {
		if node == nil {
			return
		}
		if depth >= len(result) {
			result = append(result, node.Val)
		} else if node.Val > result[depth] {
			result[depth] = node.Val
		}
		maxValues(node.Left, depth+1)
		maxValues(node.Right, depth+1)
	}
	maxValues(root, 0)
	return result
}

// 101. Symmetric Tree
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// 226. Invert Binary Tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 110. Balanced Binary Tree
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height(node.Left), height(node.Right)) + 1
}

// 257. Binary Tree Paths
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var result []string
	var paths func(node *TreeNode, path string)
	paths = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			result = append(result, path+strconv.Itoa(node.Val))
		}
		paths(node.Left, path+strconv.Itoa(node.Val)+"->")
		paths(node.Right, path+strconv.Itoa(node.Val)+"->")
	}
	paths(root, "")
	return result
}
