package main

import (
	"errors"
	"fmt"
	"math"
)

type Queue struct {
	items []interface{}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func createNode(val int) *BSTNode {
	return &BSTNode{value: val, left: nil, right: nil}
}

func (node *BSTNode) Add(val int) {
	if node != nil {
		if val < node.value {
			if node.left == nil {
				node.left = createNode(val)
			} else {
				node.left.Add(val)
			}
		} else {
			if node.right == nil {
				node.right = createNode(val)
			} else {
				node.right.Add(val)
			}
		}
	}
}

func (node *BSTNode) Remove(val int) *BSTNode {
	if node == nil {
		return nil
	}
	if val < node.value {
		node.left = node.left.Remove(val)
	} else if val > node.value {
		node.right = node.right.Remove(val)
	} else { //encontrou nó a ser removido
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.left == nil && node.right != nil {
			return node.right
		} else {
			max, _ := node.left.Max()
			node.value = max
			node.left = node.left.Remove(max)
		}
	}
	return node
}

func (node *BSTNode) Search(val int) bool {
	if node == nil {
		return false // Value not found in an empty subtree
	}
	if val == node.value {
		return true // Value found
	} else if val < node.value {
		return node.left.Search(val) // Search in the left subtree
	} else {
		return node.right.Search(val) // Search in the right subtree
	}
}

func (node *BSTNode) Min() (int, error) {
	if node == nil {
		return 0, errors.New("empty tree")
	} else {
		if node.left == nil {
			return node.value, nil
		} else {
			return node.left.Min()
		}
	}
} // Para encontrar o máximo para substituir left por right

func (node *BSTNode) Max() (int, error) {
	if node == nil {
		return 0, errors.New("empty tree")
	} else {
		if node.right == nil {
			return node.value, nil
		} else {
			return node.right.Max()
		}
	}
}

func (node *BSTNode) PreOrder() {
	if node != nil {
		fmt.Print(node.value, ", ")
		node.left.PreOrder()  // Caso fosse o EmOrdem, trocaria a ordem entre essa linha e o print
		node.right.PreOrder() // Caso fosse o PósOrdem, trocaria a ordem entre essa linha e o print
	}
}

func (node *BSTNode) InOrder() {
	if node != nil {
		node.left.InOrder()
		fmt.Print(node.value, ", ")
		node.right.InOrder()
	}
}

func (node *BSTNode) PosOrder() {
	if node != nil {
		node.left.PosOrder()
		node.right.PosOrder()
		fmt.Print(node.value, ", ")
	}
}

func (node *BSTNode) LevelOrderNav() {
	queue := &Queue{}
	queue.Enqueue(node)
	for !queue.IsEmpty() {
		item := queue.Dequeue().(*BSTNode)
		fmt.Print(item.value, ", ")
		if item.left != nil {
			queue.Enqueue(item.left)
		}
		if item.right != nil {
			queue.Enqueue(item.right)
		}
	}
}


func (node *BSTNode) height() int {
	if node == nil {
		return -1
	}
	hl := node.left.height()
	hr := node.right.height()
	return int(math.Max(float64(hl), float64(hr))) + 1
}

func (bstNode *BSTNode) IsBst() bool {
	bst := true
	if bstNode.left != nil {
		if bstNode.left.value > bstNode.value {
			return false
		} else {
			bst = bst && bstNode.left.IsBst()
		}
	}
	if bstNode.right != nil {
		if bstNode.right.value < bstNode.value {
			return false
		} else {
			bst = bst && bstNode.right.IsBst()
		}
	}
	return bst
}


func (node *BSTNode) Size() int {
	size := 1
	if node.left != nil {
		size += node.left.Size()
	}

	if node.right != nil {
		size += node.right.Size()
	}
	return size
}

func convertToBalancedBst(v []int, ini int, fim int) *BSTNode {
	//v estah ordenado
	if ini <= fim {
		meio := (ini + fim) / 2
		raiz := &BSTNode{value: v[meio]}
		raiz.left = convertToBalancedBst(v, ini, meio-1)
		raiz.right = convertToBalancedBst(v, meio+1, fim)
		return raiz
	} else {
		return nil
	}
}

func main() {
	/*
	bs := createNode(6)

	bs.Add(3)
	bs.Add(10)
	bs.Add(5)
	bs.Add(7)
	bs.Add(15)
	bs.Add(2)

	fmt.Println("Height:", bs.height())
	fmt.Print("Minimum value:")
	fmt.Println(bs.Min())
	fmt.Print("Maximum value:")
	fmt.Println(bs.Max())

	bs.PreOrder()
	fmt.Println()
	bs.InOrder()
	fmt.Println()
	bs.PosOrder()
	fmt.Println()
	bs.LevelOrderNav()
	fmt.Println()
	fmt.Println(bs.IsBst())
	fmt.Println(bs.Size())
	*/

	v := []int{50, 30, 45, 32, 10, 90, 55, 49, -5, 88, 110, 40}

	root := createNode(50)
	for i := 1; i<len(v)-1; i++{
		root.Add(v[i])
	}
	//root := convertToBalancedBst(v,0,len(v)-1)
	fmt.Println(root.IsBst()) // true
	fmt.Println(root.height())
	root.PreOrder()
	fmt.Println()
	root.PosOrder()


}
