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

// Put either inserts a new value or updates an existing value
func (n *Node) Put(key, value []byte) {
	if len(key) == 0 {
		n.Value = value
		return
	}
	if len(key) > 0 {
		next := n.Children[key[0]]
		if next == nil {
			next = NewNode()
			n.Children[key[0]] = next
		}
		next.Put(key[1:], value)
	}
	return
}

// Delete removes an existing value if found, removing parents if they are childless
func (n *Node) Delete(key []byte) {
	if len(key) == 1 {
		delete(n.Children, key[0])
		return
	}
	if len(key) > 1 {
		next := n.Children[key[0]]
		if next == nil {
			return
		}
		next.Delete(key[1:])
		if len(next.Children) == 0 && next.Value == nil {
			delete(n.Children, key[0])
		}
	}
	return
}
