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
	"strconv"
	"testing"
)

func getNum(b *testing.B, num int) {
	b.StopTimer()
	n := NewNode()
	for i := 0; i < num; i++ {
		k := []byte(strconv.Itoa(i))
		n.Put(k, k)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		k := []byte(strconv.Itoa(i % num))
		n.Get(k)
	}
}

// BenchmarkGet10000 finds the average search time for 10000 records
func BenchmarkGet10000(b *testing.B) {
	getNum(b, 10000)
}

// BenchmarkGet100000 finds the average search time for 100000 records
func BenchmarkGet100000(b *testing.B) {
	getNum(b, 100000)
}

// BenchmarkGet1000000 finds the average search time for 1000000 records
func BenchmarkGet1000000(b *testing.B) {
	getNum(b, 1000000)
}

// TestFuzzyGet makes sure that fuzzy searches work correctly
func TestFuzzyGet(t *testing.T) {
	n := NewNode()
	k1 := []byte("a")
	v1 := []byte("1")
	n.Put(k1, v1)
	k, v, ok := n.FuzzyGet(k1)
	if !ok {
		t.Log("Should have been ok")
		t.FailNow()
	}
	if !bytes.Equal(k1, k) {
		t.Errorf("Should have been '%s', found: '%s'", string(k1), string(k))
	}
	if !bytes.Equal(v1, v) {
		t.Errorf("Should have been '%s', found: '%s'", string(v1), string(v))
	}
	k2 := []byte("ab")
	k, v, ok = n.FuzzyGet(k2)
	if !ok {
		t.Log("Should have been ok")
		t.FailNow()
	}
	if !bytes.Equal(k1, k) {
		t.Errorf("Should have been '%s', found: '%s'", string(k1), string(k))
	}
	if !bytes.Equal(v1, v) {
		t.Errorf("Should have been '%s', found: '%s'", string(v1), string(v))
	}
}

func fuzzyGetNum(b *testing.B, num int) {
	b.StopTimer()
	n := NewNode()
	for i := 0; i < num; i++ {
		k := []byte(strconv.Itoa(i))
		n.Put(k, k)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		k := []byte(strconv.Itoa(i))
		n.FuzzyGet(k)
	}
}

// BenchmarkFuzzyGet10000 finds the average fuzzy search time for 10000 records
func BenchmarkFuzzyGet10000(b *testing.B) {
	fuzzyGetNum(b, 10000)
}

// BenchmarkFuzzyGet100000 finds the average fuzzy search time for 100000 records
func BenchmarkFuzzyGet100000(b *testing.B) {
	fuzzyGetNum(b, 100000)
}

// BenchmarkFuzzyGet1000000 finds the average fuzzy search time for 1000000 records
func BenchmarkFuzzyGet1000000(b *testing.B) {
	fuzzyGetNum(b, 1000000)
}

