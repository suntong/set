////////////////////////////////////////////////////////////////////////////
// Porgram: set
// Purpose: the SET data structure
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/NvwxAJgD3KI
////////////////////////////////////////////////////////////////////////////

package set

import (
	"sort"
	"strings"
)

type Set interface {
	Add(string)
	Has(string) bool
	Del(string)
	Len() int
	Get() []string
}

// set is case-sensitive SET
type set map[string]struct{}

// Add adds the given string to the set.
func (s set) Add(a string) { s[a] = struct{}{} }

// Del deletes the given string to the set. CAUTION! No key validation!
func (s set) Del(a string) { delete(s, a) }

// Has tells if the given string is in the set.
func (s set) Has(a string) bool { _, ok := s[a]; return ok }

// Len returns the number of elements in the set.
func (s set) Len() int { return len(s) }

// Get returns the keys of the set
func (s set) Get() []string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// NewSet returns a case-sensitive SET
func NewSet() set {
	s := make(set)
	return s
}

// NewSetFrom returns a case-sensitive SET from the given string slice
func NewSetFrom(i []string) set {
	s := NewSet()
	for _, v := range i {
		s.Add(v)
	}
	return s
}

// set0 is case-insensitive SET
type set0 map[string]struct{}

// Add adds the given string to the set.
func (s set0) Add(a string) { a = strings.ToLower(a); s[a] = struct{}{} }

// Del deletes the given string to the set. CAUTION! No key validation!
func (s set0) Del(a string) { a = strings.ToLower(a); delete(s, a) }

// Has tells if the given string is in the set.
func (s set0) Has(a string) bool { a = strings.ToLower(a); _, ok := s[a]; return ok }

// Len returns the number of elements in the set.
func (s set0) Len() int { return len(s) }

// Get returns the keys of the set
func (s set0) Get() []string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

// NewSet0 returns a case-sensitive SET
func NewSet0() set0 {
	s := make(set0)
	return s
}

// NewSet0From returns a case-sensitive SET from the given string slice
func NewSet0From(i []string) set0 {
	s := NewSet0()
	for _, v := range i {
		s.Add(v)
	}
	return s
}
