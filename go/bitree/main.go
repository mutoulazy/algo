package main

import "fmt"

/*
	二叉树树的每一个节点数据结构
	https://blog.csdn.net/qq_36027670/article/details/105950572
*/
type Node struct {
	Data interface{} /*树的数据*/

	Left *Node /*树的左节点*/

	Right *Node /*树的右节点*/
}

type BiTree interface {
	Inorder(tree *Node) /*中序遍历*/

	Preorder(tree *Node) /*先序遍历*/

	Postorder(tree *Node) /*后续遍历*/

	Bforder(tree *Node) /*广度优先遍历*/

	Dforder(tree *Node) /*深度优先遍历*/

	Layers() int
}

// 先序遍历
func (n Node) Preorder(tree *Node) {
	if tree == nil {
		return
	}
	fmt.Print(tree.Data, " ")
	n.Preorder(tree.Left)
	n.Preorder(tree.Right)
}

// 中序遍历
func (n Node) Inorder(tree *Node) {
	if tree == nil {
		return
	}
	n.Inorder(tree.Left)
	fmt.Print(tree.Data, " ")
	n.Inorder(tree.Right)
}

// 后续遍历
func (n Node) Postorder(tree *Node) {
	if tree == nil {
		return
	}
	n.Postorder(tree.Left)
	n.Postorder(tree.Right)
	fmt.Print(tree.Data, " ")
}

// 广度优先遍历 queue
func (n Node) Bforder(tree *Node) {
	if tree == nil {
		return
	}
	result := []interface{}{}
	nodes := []*Node{tree}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode.Data)
		if curNode.Left != nil {
			nodes = append(nodes, curNode.Left)
		}
		if curNode.Right != nil {
			nodes = append(nodes, curNode.Right)
		}
	}
	for _, v := range result {
		fmt.Print(v, " ")
	}
}

// 深度度优先遍历 等同于先序遍历 或者使用 Stack
func (n Node) Dforder(tree *Node) {
	if tree == nil {
		return
	}
	//element := make([]interface{},0)
	//stack := stack2.NewStack()
	//stack.Push(&n)
	//for !stack.IsEmpty() {
	//	node := stack.Pop().(*Node)
	//	element = append(element,node.Data)
	//	if node.Right != nil {
	//		stack.Push(node.Right)
	//	}
	//	if  node.Left != nil  {
	//		stack.Push(node.Left)
	//	}
	//}
	//for i :=0 ;i < len(element); i++ {
	//	fmt.Print(element[i]," ")
	//}
}

// 求深度 递归
func (n *Node) Layers() int {
	if n == nil {
		return 0
	}
	leftLayers := n.Left.Layers()
	rightLayers := n.Right.Layers()
	if leftLayers > rightLayers {
		return leftLayers + 1
	} else {
		return rightLayers + 1
	}
}

func main() {
	tree := &Node{
		Data: "A",
		Left: &Node{
			Data: "B",
			Left: &Node{
				Data:  "D",
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Data: "C",
				Left: nil,
				Right: &Node{
					Data:  "H",
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &Node{
			Data: "E",
			Left: &Node{
				Data:  "G",
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Data:  "F",
				Left:  nil,
				Right: nil,
			},
		},
	}
	/*前序遍历*/
	tree.Preorder(tree)
	fmt.Println()
	/*中序遍历*/
	tree.Inorder(tree)
	fmt.Println()
	/*后续遍历*/
	tree.Postorder(tree)
	fmt.Println()
	/*广度优先*/
	tree.Bforder(tree)
	fmt.Println()
	/*深度优先*/
	//tree.Dforder(tree)
	//fmt.Println()
	/*求深度*/
	fmt.Println(tree.Layers())
}
