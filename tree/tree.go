package tree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	List *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.List.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	if e := stack.List.Back(); e != nil {
		stack.List.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.List.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}

// 先序遍历
func PreTravesal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	s := stack.NewStack()
	s.push(root)
	for !s.Empty() {
		cur := s.Pop().(*Node)
		res = append(res, cur.Val)
		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
	return res
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	for root == nil {
		return nil
	}
	var res []int
	s := NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}
		if s.Empty() {
			break
		}
		cur = s.Pop().(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// 后序遍历
func PostTravesal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	s := stack.NewStack()
	out := stack.NewStack()
	s.Push(root)
	for !s.Empty() {
		cur := s.Pop().(*Node)
		out.Push(cur)

		if cur.Left != nil {
			s.Push(cur.Left)
		}

		if cur.Right != nil {
			s.Push(cur.Right)
		}
	}

	for !out.Empty() {
		cur := out.Pop().(*Node)
		res = append(res, cur.Val)
	}
	return res
}
