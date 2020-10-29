package inttreedata

import (
	"encoding/json"
	"fmt"
	"goalgs/testutils"
	"testing"
)

func TestIntNodeValueString(t *testing.T) {
	t.Run("It returns the formatted json string of the struct", func(t *testing.T) {
		tNodeData := IntTreeData{10}
		result := tNodeData.ValueString()
		expectedBytes, _ := json.MarshalIndent(tNodeData, "", " ")
		expected := string(expectedBytes)
		testutils.Expect(
			t,
			expected,
			result,
			expected == result,
		)
	})

}

func TestGetKey(t *testing.T) {
	t.Run("It returns the Value as its key", func(t *testing.T) {
		intNodeCurrent := IntTreeData{10}
		expected := intNodeCurrent.Value
		result := intNodeCurrent.GetKey()
		testutils.Expect(
			t,
			fmt.Sprint(expected),
			fmt.Sprint(result),
			result == expected,
		)
	})
}
