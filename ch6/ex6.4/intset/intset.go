// Copyright © 2017 Alan A. A. Donovan, Brian W. Kernighan & Vance E. Morris
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Elems returns a slice of integers that contains the elements of the IntSet
func (s *IntSet) Elems() []int {
	var result []int
	contents := s.String()
	// cut the { and }
	contents = contents[1 : len(contents)-1]

	if len(contents) > 0 {
		temp := strings.Split(contents, " ")
		for _, i := range temp {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			result = append(result, j)
		}
	}

	return result
}

// AddAll adds any number of integers to the IntSet
func (s *IntSet) AddAll(nums ...int) {
	for _, num := range nums {
		s.Add(num)
	}
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	result := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result++
			}
		}
	}
	return result
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] -= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var result IntSet
	result.words = make([]uint64, len(s.words))
	copy(result.words, s.words)
	return &result
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith sets to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	temp := s.Copy()
	temp.IntersectWith(t)
	for i := range s.words {
		s.words[i] ^= temp.words[i]
	}
}

// SymmetricDifference sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	intersect := s.Copy()
	intersect.IntersectWith(t)
	s.UnionWith(t)
	s.DifferenceWith(intersect)
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
