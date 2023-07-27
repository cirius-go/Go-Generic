package record

import (
	"github.com/cirius-go/generic/common"
	"github.com/cirius-go/generic/slice"
)

// Record represents the record that contains data as key & value.
type Record[K comparable, V any] map[K]V

// New create new record.
func New[K comparable, V any](maps ...Record[K, V]) Record[K, V] {
	var r = make(Record[K, V])

	return r.Apply(maps...)
}

func NewFromMap[K comparable, V any](maps ...map[K]V) Record[K, V] {
	var r = make(Record[K, V])

	return r.ApplyMap(maps...)
}

// FromMap converts m to Record.
func FromMap[K comparable, V any](m map[K]V) Record[K, V] {
	var r = make(Record[K, V])

	if m == nil {
		return r
	}

	for k, v := range m {
		r[k] = v
	}

	return r
}

func (r Record[K, V]) ApplyMap(maps ...map[K]V) Record[K, V] {
	convertedMaps := slice.Map(func(m map[K]V) Record[K, V] {
		return FromMap(m)
	}, maps...)

	return r.Apply(convertedMaps...)
}

// Apply multiple maps.
func (r Record[K, V]) Apply(maps ...Record[K, V]) Record[K, V] {
	for _, m := range maps {
		if m == nil {
			continue
		}

		for k, v := range m {
			r[k] = v
		}
	}

	return r
}

// Append or replace values.
func (r Record[K, V]) Append(name K, value V) Record[K, V] {
	if r == nil {
		return r
	}

	r[name] = value

	return r
}

// Delete values by names
func (r Record[K, V]) Delete(names ...K) Record[K, V] {
	for _, name := range names {
		delete(r, name)
	}

	return r
}

func (r Record[K, V]) SearchByKey(name K) (V, int) {
	if r == nil {
		return common.Zero[V](), -1
	}

	index := 0
	for k, v := range r {
		if k == name {
			return v, index
		}

		index++
	}

	return common.Zero[V](), -1
}

func (r Record[K, V]) DeleteByIndexes(indexes ...int) Record[K, V] {
	if r == nil {
		return r
	}

	for _, index := range indexes {
		count := 0
		for name := range r {
			if count == index {
				delete(r, name)
			}
			count++
		}
	}

	return r
}

func (r Record[K, V]) Filter(predicate func(K, V) bool) Record[K, V] {
	var (
		newR = make(Record[K, V])
	)

	for k, v := range r {
		if predicate(k, v) {
			newR[k] = v
		}
	}

	return newR
}

func (r Record[K, V]) Values() []V {
	var (
		a = []V{}
	)
	for _, v := range r {
		a = append(a, v)
	}

	return a
}
