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
	"testing"
	"strconv"
)

func getNum(b *testing.B, num int) {
    b.StopTimer()
	n := NewNode()
	for i := 0; i < num; i++ {
		k := []byte(strconv.Itoa(i))
		n.Put(k,k)
	}
    b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		k := []byte(strconv.Itoa(i%num))
		b.StartTimer()
		n.Get(k)
	}
}

// BenchmarkGet10000 finds the average search time for 10000 records
func BenchmarkGet10000(b *testing.B){
    getNum(b, 10000)
}

// BenchmarkGet100000 finds the average search time for 100000 records
func BenchmarkGet100000(b *testing.B){
    getNum(b, 100000)
}

// BenchmarkGet1000000 finds the average search time for 1000000 records
func BenchmarkGet1000000(b *testing.B){
    getNum(b, 1000000)
}