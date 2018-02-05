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

// Search gets the value of a node only if an exact match exists
func (n *Node) Search(key []byte) []byte {
	if len(key) == 0 {
		return n.Value
	}
	if len(key) > 0 {
		next := n.Children[key[0]]
		if next == nil {
			return nil
		}
		return next.Search(key[1:])
	}
	return nil
}
