package binarysearchtree

// TreeNode is the basic structure for our tree
type TreeNode struct {
	left   *TreeNode
	right  *TreeNode
	parent *TreeNode
	Count  int
	Data   TreeData
}

// GetDataAndCount Returns the data and count from a TreeNode
func (tNode *TreeNode) GetDataAndCount() (TreeData, int) {
	return tNode.Data, tNode.Count
}

//TreeData Defines the interface for allow reusability of the BST
type TreeData interface {
	ValueString() string
	GetKey() int
}

// BinarySearchTree is the basic structure for mainting our BST
type BinarySearchTree struct {
	Root *TreeNode
}

// Insert a TreeNode into the BST
func (bst *BinarySearchTree) Insert(tData TreeData) {
	nodeKey := tData.GetKey()
	node := &TreeNode{left: nil, right: nil, parent: nil, Data: tData, Count: 1}
	if bst.Root == nil {
		bst.Root = node
	} else {
		for currentNode := bst.Root; currentNode != node; {
			switch currentKey := currentNode.Data.GetKey(); {
			case nodeKey < currentKey:
				if !hasLeftChild(currentNode) {
					node.parent = currentNode
					currentNode.left = node
				}
				currentNode = currentNode.left
			case nodeKey > currentKey:
				if !hasRightChild(currentNode) {
					node.parent = currentNode
					currentNode.right = node
				}
				currentNode = currentNode.right
			default:
				currentNode.Count++
				return
			}
		}
	}
}

// Find Node returns node if found else returns error and nil for *TreeNode
func (bst *BinarySearchTree) Find(key int) *TreeNode {
	currentNode := bst.Root
	for currentNode != nil {
		switch currentKey := currentNode.Data.GetKey(); {
		case key < currentKey:
			currentNode = currentNode.left
		case key > currentKey:
			currentNode = currentNode.right
		default:
			return currentNode
		}
	}
	return currentNode
}

// Delete node if it exists.
func (bst *BinarySearchTree) Delete(key int) {
	foundNode := bst.Find(key)
	if foundNode != nil {
		if foundNode.Count > 1 {
			foundNode.Count--
		} else {
			if isRoot(foundNode) && !hasChildren(foundNode) {
				bst.Root = nil
			} else {
				deleteFromSubtree(foundNode)
			}
		}
	}
}

// InOrderSlice creates a slice of the tree data in ascending order based on key value
func (bst *BinarySearchTree) InOrderSlice() []TreeData {
	keySlice := []TreeData{}
	nodeStack := []*TreeNode{}
	current := bst.Root

	for current != nil || len(nodeStack) > 0 {
		for current != nil {
			nodeStack = append(nodeStack, current)
			current = current.left
		}
		if len(nodeStack) > 0 {
			current = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
		}
		keySlice = append(keySlice, current.Data)
		current = current.right
	}
	return keySlice
}

// PreOrderSlice creates a slice of the tree data in bst preorder
func (bst *BinarySearchTree) PreOrderSlice() []TreeData {
	keySlice := []TreeData{}
	nodeStack := []*TreeNode{}
	current := bst.Root
	for current != nil || len(nodeStack) > 0 {
		for current != nil {
			keySlice = append(keySlice, current.Data)
			if hasRightChild(current) {
				nodeStack = append(nodeStack, current.right)
			}
			current = current.left
		}
		if len(nodeStack) > 0 {
			current = nodeStack[len(nodeStack)-1]
			nodeStack = nodeStack[:len(nodeStack)-1]
		}
	}
	return keySlice
}

// PostOrderSlice creates a slice of the tree data in bst post order
func (bst *BinarySearchTree) PostOrderSlice() []TreeData {
	var previous *TreeNode = nil
	keySlice := []TreeData{}
	current := bst.Root
	nodeStack := []*TreeNode{}
	for current != nil || len(nodeStack) > 0 {
		for current != nil {
			nodeStack = append(nodeStack, current)
			current = current.left
		}
		current = nodeStack[len(nodeStack)-1]
		if !hasRightChild(current) || current.right == previous {
			keySlice = append(keySlice, current.Data)
			nodeStack = nodeStack[:len(nodeStack)-1]
			previous = current
			current = nil
		} else {
			current = current.right
		}
	}
	return keySlice
}

// LowestCommonAncestor returns the lowest node that has both as ancestors.
func (bst *BinarySearchTree) LowestCommonAncestor(lhKey int, rhKey int) *TreeNode {
	ancestor := bst.Root
	for ancestor != nil {
		switch currentKey := ancestor.Data.GetKey(); {
		case lhKey < currentKey && rhKey < currentKey:
			ancestor = ancestor.left
		case lhKey > currentKey && rhKey > currentKey:
			ancestor = ancestor.right
		default:
			return ancestor
		}
	}
	return ancestor
}

func deleteFromSubtree(root *TreeNode) {
	switch {
	case !hasChildren(root):
		replaceInParent(root, nil)
	case !hasLeftChild(root):
		replaceInParent(root, root.right)
	case !hasRightChild(root):
		replaceInParent(root, root.left)
	default:
		minNode := findMin(root.right)
		root.Data, root.Count = minNode.GetDataAndCount()
		deleteFromSubtree(minNode)
	}
}

func replaceInParent(node *TreeNode, replacementValue *TreeNode) {
	if node.parent.left == node {
		node.parent.left = replacementValue
	} else {
		node.parent.right = replacementValue
	}
}

func findMin(subTree *TreeNode) *TreeNode {
	minNode := subTree
	for minNode.left != nil {
		minNode = minNode.left
	}
	return minNode
}

func isRoot(tNode *TreeNode) bool {
	return tNode.parent == nil
}

func hasChildren(tNode *TreeNode) bool {
	return hasLeftChild(tNode) || hasRightChild(tNode)
}

func hasRightChild(tNode *TreeNode) bool {
	return tNode.right != nil
}

func hasLeftChild(tNode *TreeNode) bool {
	return tNode.left != nil
}
