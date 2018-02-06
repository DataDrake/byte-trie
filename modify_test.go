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

import (
	"bytes"
	"testing"
	"strconv"
)

// TestPut ensures that Put behaves as expected
func TestPut(t *testing.T) {
	n := NewNode()
	k1 := []byte("a")
	v1 := []byte("b")
	n.Put(k1, v1)
	v, ok := n.Get(k1)
	if !ok {
		t.Log("Should have found matching node")
		t.FailNow()
	}
	if !bytes.Equal(v, v1) {
		t.Logf("Should have been '%s', found '%s'", string(v1), string(v))
		t.FailNow()
	}
	k2 := []byte("b")
	v2 := []byte("b")
	n.Put(k2, v2)
	v, ok = n.Get(k2)
	if !ok {
		t.Log("Should have found matching node")
		t.FailNow()
	}
	if !bytes.Equal(v, v2) {
		t.Logf("Should have been '%s', found '%s'", string(v2), string(v))
		t.FailNow()
	}
	k3 := []byte("cd")
	v3 := []byte("e")
	n.Put(k3, v3)
	v, ok = n.Get(k3)
	if !ok {
		t.Log("Should have found matching node")
		t.FailNow()
	}
	if !bytes.Equal(v, v3) {
		t.Logf("Should have been '%s', found '%s'", string(v3), string(v))
		t.FailNow()
	}
	k4 := []byte("c")
	v, ok = n.Get(k4)
	if ok {
		t.Log("Should not have found matching node")
		t.FailNow()
	}
	if v != nil {
		t.Errorf("Should have been nil, found '%s'", string(v))
	}
}

// TestDelete ensures that deletions work correctly
func TestDelete(t *testing.T) {
	n := NewNode()
	k1 := []byte("a")
	k2 := []byte("ab")
	k3 := []byte("abc")
	k4 := []byte("abcd")
	v := []byte("mark")
	n.Put(k1, v)
	n.Put(k3, v)
	n.Put(k4, v)

	n.Delete(k1)
	_, ok := n.Get(k1)
	if ok {
		t.Log("k1 should not still exist")
		t.FailNow()
	}
	_, ok = n.Get(k2)
	if ok {
		t.Log("k2 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k3)
	if !ok {
		t.Log("k3 should still exist")
		t.FailNow()
	}
	_, ok = n.Get(k4)
	if !ok {
		t.Log("k4 should still exist")
		t.FailNow()
	}

	n.Delete(k2)
	_, ok = n.Get(k1)
	if ok {
		t.Log("k1 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k2)
	if ok {
		t.Log("k2 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k3)
	if !ok {
		t.Log("k3 should still exist")
		t.FailNow()
	}
	_, ok = n.Get(k4)
	if !ok {
		t.Log("k4 should still exist")
		t.FailNow()
	}

	n.Delete(k4)
	_, ok = n.Get(k1)
	if ok {
		t.Log("k1 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k2)
	if ok {
		t.Log("k2 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k3)
	if !ok {
		t.Log("k3 should still exist")
		t.FailNow()
	}
	_, ok = n.Get(k4)
	if ok {
		t.Log("k4 should not exist")
		t.FailNow()
	}

	n.Delete(k3)
	_, ok = n.Get(k1)
	if ok {
		t.Log("k1 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k2)
	if ok {
		t.Log("k2 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k3)
	if ok {
		t.Log("k3 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k4)
	if ok {
		t.Log("k4 should not exist")
		t.FailNow()
	}

	n.Delete(k3)
	_, ok = n.Get(k1)
	if ok {
		t.Log("k1 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k2)
	if ok {
		t.Log("k2 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k3)
	if ok {
		t.Log("k3 should not exist")
		t.FailNow()
	}
	_, ok = n.Get(k4)
	if ok {
		t.Log("k4 should not exist")
		t.FailNow()
	}

}

// BenchmarkPut finds the average insertion time of a number of records
func BenchmarkPut(b *testing.B) {
	n := NewNode()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		k := []byte(strconv.Itoa(i))
		b.StartTimer()
		n.Put(k,k)
	}
}