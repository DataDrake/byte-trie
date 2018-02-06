//
// Copyright 2018 Bryan T. Meyers <bmeyers@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package trie

// Node is an element in the trie structure
type Node struct {
	Children map[byte]*Node
	Value    []byte
}

// NewNode returns a pointer to an empty Node
func NewNode() *Node {
	return &Node{make(map[byte]*Node), nil}
}

// IsLeaf checks if this Node is at the end of a branch
func (n *Node) IsLeaf() bool {
	return n.Value == nil || len(n.Value) > 0
}
