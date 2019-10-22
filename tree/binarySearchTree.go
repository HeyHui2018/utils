package tree

/*
Insert(val)         // 在二叉搜索树中插入一个元素
Search(val)         // 在二叉搜索树中搜索一个元素
Parent(val)         // 返回当前节点的父节点
Contains(val)       // 在二叉搜索中是否包含这个元素
Min()               // 返回二叉搜索树中的最小的元素
Max()               // 返回二叉搜索树中的最大的元素
Remove(val)         // 删除二叉搜索树中的一个元素
PreOrder()          // 前序遍历二叉搜索树
InOrder()           // 中序遍历二叉搜索树
PostOrder()         // 后续遍历二叉搜索树
*/

type BinarySearchTree struct {
	root *TreeNode
}

// 在二叉搜索树中插入一个元素
func (tree *BinarySearchTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &TreeNode{val, nil, nil}
	} else {
		insertNode(val, tree.root)
	}
}

func insertNode(val int, root *TreeNode) {
	if val < root.val {
		if root.left == nil {
			root.left = &TreeNode{val, nil, nil}
		} else {
			insertNode(val, root.left)
		}
	} else {
		if root.right == nil {
			root.right = &TreeNode{val, nil, nil}
		} else {
			insertNode(val, root.right)
		}
	}
}

// 在二叉搜索树中搜索一个元素
func (tree *BinarySearchTree) Search(val int) *TreeNode {
	if tree.root != nil {
		return findNode(val, tree.root)
	}
	return nil
}

func findNode(val int, root *TreeNode) *TreeNode {
	if val == root.val {
		return root
	} else if val < root.val {
		if root.left != nil {
			return findNode(val, root.left)
		}
	} else {
		if root.right != nil {
			return findNode(val, root.right)
		}
	}
	return nil
}

// 返回当前节点的父节点
func (tree *BinarySearchTree) Parent(val int) *TreeNode {
	if tree.root != nil {
		return findParent(val, tree.root)
	}
	return nil
}

func findParent(val int, root *TreeNode) *TreeNode {
	if val == root.val {
		return nil
	} else if val < root.val {
		if root.left != nil {
			if val == root.left.val {
				return root
			} else {
				return findParent(val, root.left)
			}
		}
	} else {
		if root.right != nil {
			if val == root.right.val {
				return root
			} else {
				return findParent(val, root.right)
			}
		}
	}
	return nil
}

// 在二叉搜索中是否包含这个元素
func (tree *BinarySearchTree) Contains(val int) bool {
	if tree.root != nil {
		return containNode(val, tree.root)
	}
	return false
}

func containNode(val int, root *TreeNode) bool {
	if val == root.val {
		return true
	} else if val < root.val {
		if root.left == nil {
			return false
		} else {
			return containNode(val, root.left)
		}
	} else {
		if root.right == nil {
			return false
		} else {
			return containNode(val, root.right)
		}
	}
}

// 返回二叉搜索树中的最小的元素
func (tree *BinarySearchTree) Min() *TreeNode {
	if tree.root != nil {
		return findMin(tree.root)
	}
	return nil
}

func findMin(root *TreeNode) *TreeNode {
	if root.left == nil {
		return root
	} else {
		return findMin(root.left)
	}
}

// 返回二叉搜索树中的最大的元素
func (tree *BinarySearchTree) Max() *TreeNode {
	if tree.root != nil {
		return findMax(tree.root)
	}
	return nil
}

func findMax(root *TreeNode) *TreeNode {
	if root.right == nil {
		return root
	} else {
		return findMax(root.right)
	}
}

// 删除二叉搜索树中的一个元素
func (tree *BinarySearchTree) Remove(val int) bool {
	nodeToRemoved := tree.Search(val)
	if nodeToRemoved == nil {
		return false
	}

	parent := tree.Parent(val)
	if tree.root.left == nil && tree.root.right == nil {
		tree.root = nil
		return true
	} else if nodeToRemoved.left == nil && nodeToRemoved.right == nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else if nodeToRemoved.left == nil && nodeToRemoved.right != nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nodeToRemoved.right
		} else {
			parent.right = nodeToRemoved.right
		}
	} else if nodeToRemoved.left != nil && nodeToRemoved.right == nil {
		if nodeToRemoved.val < parent.val {
			parent.left = nodeToRemoved.left
		} else {
			parent.right = nodeToRemoved.left
		}
	} else {
		largestValue := nodeToRemoved.left
		for largestValue.right != nil {
			largestValue = largestValue.right
		}
		tree.Remove(largestValue.val)
		nodeToRemoved.val = largestValue.val
	}

	return true
}

// 前序遍历二叉搜索树
func (tree *BinarySearchTree) PreOrder() {
	preOrder(tree.root)
}

func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.val, "->")
		preOrder(root.left)
		preOrder(root.right)
	}
}

// 中序遍历二叉搜索树
func (tree *BinarySearchTree) InOrder() {
	inOrder(tree.root)
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.left)
		fmt.Print(root.val, "->")
		inOrder(root.right)
	}
}

// 后续遍历二叉搜索树
func (tree *BinarySearchTree) PostOrder() {
	postOrder(tree.root)
}

func postOrder(root *TreeNode) {
	if root != nil {
		postOrder(root.right)
		postOrder(root.left)
		fmt.Print(root.val, "->")
	}
}
