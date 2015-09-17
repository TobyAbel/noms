// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package datas

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// Commit

type Commit struct {
	m types.Map
}

func NewCommit() Commit {
	return Commit{
		types.NewMap(types.NewString("$name"), types.NewString("Commit")),
	}
}

func CommitFromVal(v types.Value) Commit {
	return Commit{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Commit) NomsValue() types.Map {
	return s.m
}

func (s Commit) Equals(p Commit) bool {
	return s.m.Equals(p.m)
}

func (s Commit) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Commit) Parents() types.Set {
	return types.SetFromVal(s.m.Get(types.NewString("parents")))
}

func (s Commit) SetParents(p types.Set) Commit {
	return CommitFromVal(s.m.Set(types.NewString("parents"), p))
}

func (s Commit) Value() types.Value {
	return (s.m.Get(types.NewString("value")))
}

func (s Commit) SetValue(p types.Value) Commit {
	return CommitFromVal(s.m.Set(types.NewString("value"), p))
}

// SetOfCommit

type SetOfCommit struct {
	s types.Set
}

type SetOfCommitIterCallback (func(p Commit) (stop bool))

func NewSetOfCommit() SetOfCommit {
	return SetOfCommit{types.NewSet()}
}

func SetOfCommitFromVal(p types.Value) SetOfCommit {
	return SetOfCommit{p.(types.Set)}
}

func (s SetOfCommit) NomsValue() types.Set {
	return s.s
}

func (s SetOfCommit) Equals(p SetOfCommit) bool {
	return s.s.Equals(p.s)
}

func (s SetOfCommit) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfCommit) Empty() bool {
	return s.s.Empty()
}

func (s SetOfCommit) Len() uint64 {
	return s.s.Len()
}

func (s SetOfCommit) Has(p Commit) bool {
	return s.s.Has(p.NomsValue())
}

func (s SetOfCommit) Iter(cb SetOfCommitIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(CommitFromVal(v))
	})
}

func (s SetOfCommit) Insert(p ...Commit) SetOfCommit {
	return SetOfCommit{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfCommit) Remove(p ...Commit) SetOfCommit {
	return SetOfCommit{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfCommit) Union(others ...SetOfCommit) SetOfCommit {
	return SetOfCommit{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfCommit) Subtract(others ...SetOfCommit) SetOfCommit {
	return SetOfCommit{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfCommit) Any() Commit {
	return CommitFromVal(s.s.Any())
}

func (s SetOfCommit) fromStructSlice(p []SetOfCommit) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfCommit) fromElemSlice(p []Commit) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

