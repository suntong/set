////////////////////////////////////////////////////////////////////////////
// Porgram: set
// Purpose: the SET data structure
// https://groups.google.com/forum/?fromgroups=#!topic/golang-nuts/NvwxAJgD3KI
////////////////////////////////////////////////////////////////////////////

package set

import (
	"sort"
)

type Set interface {
	Add(string)
	Has(string) bool
	Del(string)
	Len() int
	Get() []string
}

type set map[string]struct{}

func (s set) Add(a string)      { s[a] = struct{}{} }
func (s set) Del(a string)      { delete(s, a) }
func (s set) Has(a string) bool { _, ok := s[a]; return ok }
func (s set) Len() int          { return len(s) }

func (s set) Get() []string {
	var keys []string
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func NewSet() set {
	s := make(set)
	return s
}

func NewSetFrom(i []string) set {
	s := NewSet()
	for _, v := range i {
		s.Add(v)
	}
	return s
}
