package binarysearchtree

import (
	"fmt"
	itd "goalgs/binarysearchtree/inttreedata"
	"goalgs/testutils"
	"reflect"
	"testing"
)

func TestBSTInsert(t *testing.T) {
	t.Run("It sets the inserted node to root if no root node exists", func(t *testing.T) {
		intNode := itd.IntTreeData{Value: 1}
		intTree := BinarySearchTree{}
		intTree.Insert(intNode)
		expected := intNode
		actual := intTree.Root.Data

		testutils.Expect(
			t,
			fmt.Sprintf("%+v", expected),
			fmt.Sprintf("%+v", actual),
			reflect.DeepEqual(expected, actual),
		)
	})

	t.Run("It inserts Data to the left of the node if the new value is less than the current node value", func(t *testing.T) {
		intTree := BinarySearchTree{}
		intNode := itd.IntTreeData{Value: 0}
		intTree.Insert(itd.IntTreeData{Value: 2})
		intTree.Insert(itd.IntTreeData{Value: 1})
		intTree.Insert(intNode)
		expected := intNode
		actual := intTree.Root.left.left.Data
		testutils.Expect(
			t,
			fmt.Sprintf("%+v", expected),
			fmt.Sprintf("%+v", actual),
			reflect.DeepEqual(expected, actual),
		)
	})

	t.Run("It increments the count if the current key already exists", func(t *testing.T) {
		intTree := BinarySearchTree{}
		intNode := itd.IntTreeData{Value: 0}
		intTree.Insert(itd.IntTreeData{Value: 1})
		intTree.Insert(intNode)
		intTree.Insert(intNode)
		expected := 2
		actual := intTree.Root.left.Count
		testutils.Expect(
			t,
			fmt.Sprintf("%+v", expected),
			fmt.Sprintf("%+v", actual),
			expected == actual,
		)
	})

	t.Run("It inserts to the right if the new value is greater than the current value", func(t *testing.T) {
		intTree := BinarySearchTree{}
		intNode := itd.IntTreeData{Value: 3}
		intTree.Insert(itd.IntTreeData{Value: 1})
		intTree.Insert(itd.IntTreeData{Value: 2})
		intTree.Insert(intNode)
		expected := intNode
		actual := intTree.Root.right.right.Data
		testutils.Expect(
			t,
			fmt.Sprintf("%+v", expected),
			fmt.Sprintf("%+v", actual),
			reflect.DeepEqual(expected, actual),
		)
	})
}

func TestBSTFind(t *testing.T) {
	t.Run("It returns nil if the node is note found", func(t *testing.T) {
		bst := BinarySearchTree{}
		var expected *TreeNode
		actual := bst.Find(1)
		testutils.Expect(
			t,
			fmt.Sprintf("%v", expected),
			fmt.Sprintf("%v", actual),
			reflect.DeepEqual(expected, actual),
		)

	})

	t.Run("It returns the current node if the keys match", func(t *testing.T) {
		bst := BinarySearchTree{}
		itdData := itd.IntTreeData{Value: 10}
		bst.Insert(itdData)
		expected := itdData
		actual := bst.Find(itdData.GetKey())
		testutils.Expect(
			t,
			expected.ValueString(),
			actual.Data.ValueString(),
			expected.ValueString() == actual.Data.ValueString(),
		)
	})

	t.Run("It searches left if the value to find is less than current value", func(t *testing.T) {
		bst := BinarySearchTree{}
		intData := itd.IntTreeData{Value: 10}
		intData2 := itd.IntTreeData{Value: 50}
		bst.Insert(intData2)
		bst.Insert(intData)
		expectedData := bst.Root.left.Data
		expectedCount := 1
		actualNode := bst.Find(intData.GetKey())
		testutils.Expect(
			t,
			fmt.Sprintf("%+v, %d", expectedData, expectedCount),
			fmt.Sprintf("%+v, %d", actualNode.Data, actualNode.Count),
			expectedCount == actualNode.Count && expectedData.ValueString() == actualNode.Data.ValueString(),
		)
	})

	t.Run("It searches right if the value to find is greater than current value", func(t *testing.T) {
		bst := BinarySearchTree{}
		intData := itd.IntTreeData{Value: 50}
		intData2 := itd.IntTreeData{Value: 10}
		bst.Insert(intData2)
		bst.Insert(intData)
		expectedData := bst.Root.right.Data
		expectedCount := 1
		actualNode := bst.Find(intData.GetKey())
		testutils.Expect(
			t,
			fmt.Sprintf("%+v, %d", expectedData, expectedCount),
			fmt.Sprintf("%+v, %d", actualNode.Data, actualNode.Count),
			expectedCount == actualNode.Count && expectedData.ValueString() == actualNode.Data.ValueString(),
		)
	})
}

func TestBSTDelete(t *testing.T) {
	t.Run("It decrements the nodes count if the nodes count is greater than 0", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 0})
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Delete(1)
		expected := 1
		actual := bst.Root.Count
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			expected == actual,
		)
	})

	t.Run("It decrements the roots count if the node has no parent and the count is greater than 0", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 0})
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Delete(1)
		expected := 1
		actual := bst.Root.Count
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			expected == actual,
		)
	})

	t.Run("It sets the root to nil if the node has no children, no parent, and the Count - 1 is zero", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Delete(1)
		var expected *TreeNode
		actual := bst.Root
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			actual == expected,
		)
	})

	t.Run("It sets the left child of parent node to nil if the node is the left child and has no children", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Insert(itd.IntTreeData{Value: 0})
		bst.Delete(0)
		var expected *TreeNode
		actual := bst.Root.left
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			actual == expected,
		)
	})

	t.Run("It sets the right child of parent node to nil if the node is the right child and has no children", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Insert(itd.IntTreeData{Value: 2})
		bst.Delete(2)
		var expected *TreeNode
		actual := bst.Root.right
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			actual == expected,
		)
	})

	t.Run("It sets the node to delete to the right child if no left child", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 5})
		bst.Insert(itd.IntTreeData{Value: 2})
		bst.Insert(itd.IntTreeData{Value: 4})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Delete(2)
		expectedReplacement := 4
		actualReplacement := bst.Root.left.Data.GetKey()

		testutils.Expect(
			t,
			fmt.Sprint(expectedReplacement),
			fmt.Sprint(actualReplacement),
			actualReplacement == expectedReplacement,
		)
	})

	t.Run("It replaces node in parent with the left child when there is no right child", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 5})
		bst.Insert(itd.IntTreeData{Value: 4})
		bst.Insert(itd.IntTreeData{Value: 2})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Delete(4)
		expectedReplacement := 2
		actualReplacement := bst.Root.left.Data.GetKey()

		testutils.Expect(
			t,
			fmt.Sprint(expectedReplacement),
			fmt.Sprint(actualReplacement),
			actualReplacement == expectedReplacement,
		)

	})

	t.Run("It replaces with the min of the right sub tree when the node has two children", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 6})
		bst.Insert(itd.IntTreeData{Value: 2})
		bst.Insert(itd.IntTreeData{Value: 1})
		bst.Insert(itd.IntTreeData{Value: 5})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Insert(itd.IntTreeData{Value: 4})
		bst.Delete(2)
		expectedReplacement := 3
		actualReplacement := bst.Root.left.Data.GetKey()
		expectedSecondReplacement := 4
		actualSecondReplacement := bst.Root.left.right.left.Data.GetKey()
		var expectedRemoved *TreeNode
		actualRemoved := bst.Root.left.right.left.right

		testutils.Expect(
			t,
			fmt.Sprint(expectedReplacement),
			fmt.Sprint(actualReplacement),
			actualReplacement == expectedReplacement,
		)
		testutils.Expect(
			t,
			fmt.Sprint(expectedSecondReplacement),
			fmt.Sprint(actualSecondReplacement),
			actualSecondReplacement == expectedSecondReplacement,
		)
		testutils.Expect(
			t,
			fmt.Sprint(expectedRemoved),
			fmt.Sprint(actualRemoved),
			expectedRemoved == actualRemoved,
		)
	})
}

func TestBSTInOrderSlice(t *testing.T) {
	t.Run("It returns a slice of the nodes in order based on the key", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 10})
		bst.Insert(itd.IntTreeData{Value: 7})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Insert(itd.IntTreeData{Value: 12})
		bst.Insert(itd.IntTreeData{Value: 43})
		bst.Insert(itd.IntTreeData{Value: 6})
		expected := []int{3, 6, 7, 10, 12, 43}
		actualData := bst.InOrderSlice()
		actual := []int{}
		for _, n := range actualData {
			actual = append(actual, n.GetKey())
		}
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(expected) == fmt.Sprint(actual),
		)
	})

	t.Run("It returns an empty slice if the root is nil", func(t *testing.T) {
		bst := BinarySearchTree{}
		expected := []TreeData{}
		actual := bst.InOrderSlice()
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(actual) == fmt.Sprint(expected),
		)
	})

}
func TestBSTPreOrderSlice(t *testing.T) {
	t.Run("It returns a slice of the nodes in order based on the key", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 10})
		bst.Insert(itd.IntTreeData{Value: 7})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Insert(itd.IntTreeData{Value: 12})
		bst.Insert(itd.IntTreeData{Value: 43})
		bst.Insert(itd.IntTreeData{Value: 6})
		expected := []int{10, 7, 3, 6, 12, 43}
		actualData := bst.PreOrderSlice()
		actual := []int{}
		for _, n := range actualData {
			actual = append(actual, n.GetKey())
		}
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(expected) == fmt.Sprint(actual),
		)
	})

	t.Run("It returns an empty slice if the root is nil", func(t *testing.T) {
		bst := BinarySearchTree{}
		expected := []TreeData{}
		actual := bst.PreOrderSlice()
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(actual) == fmt.Sprint(expected),
		)
	})
}

func TestBSTPostOrderSlice(t *testing.T) {
	t.Run("It returns a slice of the nodes in order based on the key", func(t *testing.T) {
		bst := BinarySearchTree{}
		bst.Insert(itd.IntTreeData{Value: 10})
		bst.Insert(itd.IntTreeData{Value: 7})
		bst.Insert(itd.IntTreeData{Value: 3})
		bst.Insert(itd.IntTreeData{Value: 12})
		bst.Insert(itd.IntTreeData{Value: 43})
		bst.Insert(itd.IntTreeData{Value: 6})
		expected := []int{6, 3, 7, 43, 12, 10}
		actualData := bst.PostOrderSlice()
		actual := []int{}
		for _, n := range actualData {
			actual = append(actual, n.GetKey())
		}
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(expected) == fmt.Sprint(actual),
		)
	})

	t.Run("It returns an empty slice if the root is nil", func(t *testing.T) {
		bst := BinarySearchTree{}
		expected := []TreeData{}
		actual := bst.PreOrderSlice()
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(actual),
			fmt.Sprint(actual) == fmt.Sprint(expected),
		)
	})
}
