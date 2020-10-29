package inttreedata

import (
	"encoding/json"
)

// IntTreeData is the integer implementation of TreeDataNode
type IntTreeData struct {
	Value int `json:"value"`
}

// ValueString returns the string form of the nodes Value
func (intNode IntTreeData) ValueString() string {
	jsonString, _ := json.MarshalIndent(intNode, "", " ")
	return string(jsonString)
}

// GetKey returns the integer key value for the node
func (intNode IntTreeData) GetKey() int {
	return intNode.Value
}
