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

import "testing"

// TestNewNode ensures that NewNode returns a valid Node
func TestNewNode(t *testing.T) {
	n := NewNode()
	if n == nil {
		t.Log("Node should not be nil")
		t.FailNow()
	}
	if len(n.Children) != 0 {
		t.Errorf("Node should not have children, found: %d", len(n.Children))
	}
	if n.Value != nil {
		t.Errorf("Value should be nil, found: '%v'", n.Value)
	}
}

// TestIsLeaf ensures that a new node is a leaf
func TestIsLeaf(t *testing.T) {
	n := NewNode()
	if !n.IsLeaf() {
		t.Error("New Node should be a leaf")
	}
}
