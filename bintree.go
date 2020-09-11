package exercises

type BinTreeNode struct {
	data        int
	left, right *BinTreeNode
}

type NodeFunc func(node *BinTreeNode)

func LMR(n *BinTreeNode, f NodeFunc) {
	if n.left != nil {
		LMR(n.left, f)
	}

	f(n)

	if n.right != nil {
		LMR(n.right, f)
	}
}

func LMR2(n *BinTreeNode, f NodeFunc) {
	stack := make([]*BinTreeNode, 0)
	// 因为父节点后面需要用，所以才入栈
	pop := func() *BinTreeNode {
		top := len(stack) - 1
		ret := stack[top]
		stack = stack[:top]
		return ret
	}
	push := func(node *BinTreeNode) {
		stack = append(stack, node)
	}

	for {
		for n.left != nil {
			push(n)
			n = n.left
		}
		f(n)

		for n.right == nil {
			if len(stack) == 0 {
				return
			}
			n = pop()
			f(n)
		}

		n = n.right
	}
}

func LMR3(n *BinTreeNode, f NodeFunc) {
	stack := make([]*BinTreeNode, 0)
	// 因为父节点后面需要用，所以才入栈
	pop := func() *BinTreeNode {
		top := len(stack) - 1
		ret := stack[top]
		stack = stack[:top]
		return ret
	}
	push := func(node *BinTreeNode) {
		stack = append(stack, node)
	}

	for {
		for n != nil {
			push(n)
			n = n.left
		}

		for {
			if len(stack) == 0 {
				return
			}
			n = pop()
			f(n)
			if n.right != nil {
				n = n.right
				break
			}
		}
	}
}

func MLR(n *BinTreeNode, f NodeFunc) {
	f(n)

	if n.left != nil {
		MLR(n.left, f)
	}

	if n.right != nil {
		MLR(n.right, f)
	}
}

func BS(n *BinTreeNode, f NodeFunc) {
	queue := NewQueue(1000)
	queue.EnQueue(n)

	for {
		if data, ok := queue.DeQueue(); !ok {
			break
		} else {
			n := data.(*BinTreeNode)
			f(n)
			if n.left != nil {
				if !queue.EnQueue(n.left) {
					panic("queue full")
				}
			}
			if n.right != nil {
				if !queue.EnQueue(n.right) {
					panic("queue full")
				}
			}
		}
	}
}
