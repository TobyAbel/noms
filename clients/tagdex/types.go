// This file was generated by nomgen.
// To regenerate, run `go generate` in this package.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

// RemotePhoto

type RemotePhoto struct {
	m types.Map
}

func NewRemotePhoto() RemotePhoto {
	return RemotePhoto{
		types.NewMap(types.NewString("$name"), types.NewString("RemotePhoto")),
	}
}

func RemotePhotoFromVal(v types.Value) RemotePhoto {
	return RemotePhoto{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s RemotePhoto) NomsValue() types.Map {
	return s.m
}

func (s RemotePhoto) Equals(p RemotePhoto) bool {
	return s.m.Equals(p.m)
}

func (s RemotePhoto) Ref() ref.Ref {
	return s.m.Ref()
}

func (s RemotePhoto) Sizes() MapOfSizeToString {
	return MapOfSizeToStringFromVal(s.m.Get(types.NewString("sizes")))
}

func (s RemotePhoto) SetSizes(p MapOfSizeToString) RemotePhoto {
	return RemotePhotoFromVal(s.m.Set(types.NewString("sizes"), p.NomsValue()))
}

func (s RemotePhoto) Title() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("title")))
}

func (s RemotePhoto) SetTitle(p types.String) RemotePhoto {
	return RemotePhotoFromVal(s.m.Set(types.NewString("title"), p))
}

func (s RemotePhoto) Id() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("id")))
}

func (s RemotePhoto) SetId(p types.String) RemotePhoto {
	return RemotePhotoFromVal(s.m.Set(types.NewString("id"), p))
}

func (s RemotePhoto) Tags() SetOfString {
	return SetOfStringFromVal(s.m.Get(types.NewString("tags")))
}

func (s RemotePhoto) SetTags(p SetOfString) RemotePhoto {
	return RemotePhotoFromVal(s.m.Set(types.NewString("tags"), p.NomsValue()))
}

func (s RemotePhoto) Url() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("url")))
}

func (s RemotePhoto) SetUrl(p types.String) RemotePhoto {
	return RemotePhotoFromVal(s.m.Set(types.NewString("url"), p))
}

// SetOfString

type SetOfString struct {
	s types.Set
}

type SetOfStringIterCallback (func(p types.String) (stop bool))

func NewSetOfString() SetOfString {
	return SetOfString{types.NewSet()}
}

func SetOfStringFromVal(p types.Value) SetOfString {
	return SetOfString{p.(types.Set)}
}

func (s SetOfString) NomsValue() types.Set {
	return s.s
}

func (s SetOfString) Equals(p SetOfString) bool {
	return s.s.Equals(p.s)
}

func (s SetOfString) Ref() ref.Ref {
	return s.s.Ref()
}

func (s SetOfString) Empty() bool {
	return s.s.Empty()
}

func (s SetOfString) Len() uint64 {
	return s.s.Len()
}

func (s SetOfString) Has(p types.String) bool {
	return s.s.Has(p)
}

func (s SetOfString) Iter(cb SetOfStringIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(types.StringFromVal(v))
	})
}

func (s SetOfString) Insert(p ...types.String) SetOfString {
	return SetOfString{s.s.Insert(s.fromElemSlice(p)...)}
}

func (s SetOfString) Remove(p ...types.String) SetOfString {
	return SetOfString{s.s.Remove(s.fromElemSlice(p)...)}
}

func (s SetOfString) Union(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Union(s.fromStructSlice(others)...)}
}

func (s SetOfString) Subtract(others ...SetOfString) SetOfString {
	return SetOfString{s.s.Subtract(s.fromStructSlice(others)...)}
}

func (s SetOfString) Any() types.String {
	return types.StringFromVal(s.s.Any())
}

func (s SetOfString) fromStructSlice(p []SetOfString) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfString) fromElemSlice(p []types.String) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

// Photo

type Photo struct {
	m types.Map
}

func NewPhoto() Photo {
	return Photo{
		types.NewMap(types.NewString("$name"), types.NewString("Photo")),
	}
}

func PhotoFromVal(v types.Value) Photo {
	return Photo{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Photo) NomsValue() types.Map {
	return s.m
}

func (s Photo) Equals(p Photo) bool {
	return s.m.Equals(p.m)
}

func (s Photo) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Photo) Width() types.UInt32 {
	return types.UInt32FromVal(s.m.Get(types.NewString("width")))
}

func (s Photo) SetWidth(p types.UInt32) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("width"), p))
}

func (s Photo) Title() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("title")))
}

func (s Photo) SetTitle(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("title"), p))
}

func (s Photo) Id() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("id")))
}

func (s Photo) SetId(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("id"), p))
}

func (s Photo) Tags() SetOfString {
	return SetOfStringFromVal(s.m.Get(types.NewString("tags")))
}

func (s Photo) SetTags(p SetOfString) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("tags"), p.NomsValue()))
}

func (s Photo) Height() types.UInt32 {
	return types.UInt32FromVal(s.m.Get(types.NewString("height")))
}

func (s Photo) SetHeight(p types.UInt32) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("height"), p))
}

func (s Photo) Image() types.Blob {
	return types.BlobFromVal(s.m.Get(types.NewString("image")))
}

func (s Photo) SetImage(p types.Blob) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("image"), p))
}

func (s Photo) Url() types.String {
	return types.StringFromVal(s.m.Get(types.NewString("url")))
}

func (s Photo) SetUrl(p types.String) Photo {
	return PhotoFromVal(s.m.Set(types.NewString("url"), p))
}

// MapOfSizeToString

type MapOfSizeToString struct {
	m types.Map
}

type MapOfSizeToStringIterCallback (func(k Size, v types.String) (stop bool))

func NewMapOfSizeToString() MapOfSizeToString {
	return MapOfSizeToString{types.NewMap()}
}

func MapOfSizeToStringFromVal(p types.Value) MapOfSizeToString {
	return MapOfSizeToString{p.(types.Map)}
}

func (m MapOfSizeToString) NomsValue() types.Map {
	return m.m
}

func (m MapOfSizeToString) Equals(p MapOfSizeToString) bool {
	return m.m.Equals(p.m)
}

func (m MapOfSizeToString) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfSizeToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfSizeToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfSizeToString) Has(p Size) bool {
	return m.m.Has(p.NomsValue())
}

func (m MapOfSizeToString) Get(p Size) types.String {
	return types.StringFromVal(m.m.Get(p.NomsValue()))
}

func (m MapOfSizeToString) Set(k Size, v types.String) MapOfSizeToString {
	return MapOfSizeToStringFromVal(m.m.Set(k.NomsValue(), v))
}

// TODO: Implement SetM?

func (m MapOfSizeToString) Remove(p Size) MapOfSizeToString {
	return MapOfSizeToStringFromVal(m.m.Remove(p.NomsValue()))
}

func (m MapOfSizeToString) Iter(cb MapOfSizeToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(SizeFromVal(k), types.StringFromVal(v))
	})
}

// MapOfStringToSet

type MapOfStringToSet struct {
	m types.Map
}

type MapOfStringToSetIterCallback (func(k types.String, v types.Set) (stop bool))

func NewMapOfStringToSet() MapOfStringToSet {
	return MapOfStringToSet{types.NewMap()}
}

func MapOfStringToSetFromVal(p types.Value) MapOfStringToSet {
	return MapOfStringToSet{p.(types.Map)}
}

func (m MapOfStringToSet) NomsValue() types.Map {
	return m.m
}

func (m MapOfStringToSet) Equals(p MapOfStringToSet) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToSet) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToSet) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToSet) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToSet) Has(p types.String) bool {
	return m.m.Has(p)
}

func (m MapOfStringToSet) Get(p types.String) types.Set {
	return types.SetFromVal(m.m.Get(p))
}

func (m MapOfStringToSet) Set(k types.String, v types.Set) MapOfStringToSet {
	return MapOfStringToSetFromVal(m.m.Set(k, v))
}

// TODO: Implement SetM?

func (m MapOfStringToSet) Remove(p types.String) MapOfStringToSet {
	return MapOfStringToSetFromVal(m.m.Remove(p))
}

func (m MapOfStringToSet) Iter(cb MapOfStringToSetIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(types.StringFromVal(k), types.SetFromVal(v))
	})
}

// Size

type Size struct {
	m types.Map
}

func NewSize() Size {
	return Size{
		types.NewMap(types.NewString("$name"), types.NewString("Size")),
	}
}

func SizeFromVal(v types.Value) Size {
	return Size{v.(types.Map)}
}

// TODO: This was going to be called Value() but it collides with root.value. We need some other place to put the built-in fields like Value() and Equals().
func (s Size) NomsValue() types.Map {
	return s.m
}

func (s Size) Equals(p Size) bool {
	return s.m.Equals(p.m)
}

func (s Size) Ref() ref.Ref {
	return s.m.Ref()
}

func (s Size) Width() types.UInt32 {
	return types.UInt32FromVal(s.m.Get(types.NewString("width")))
}

func (s Size) SetWidth(p types.UInt32) Size {
	return SizeFromVal(s.m.Set(types.NewString("width"), p))
}

func (s Size) Height() types.UInt32 {
	return types.UInt32FromVal(s.m.Get(types.NewString("height")))
}

func (s Size) SetHeight(p types.UInt32) Size {
	return SizeFromVal(s.m.Set(types.NewString("height"), p))
}

