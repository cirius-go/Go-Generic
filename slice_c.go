package generic

type B struct {
}

// SliceC represents the slice of comparable items.
// Waiting for https://github.com/vincenzopalazzo/cln4go/issues/12
type SliceC[T comparable] SliceA[T]
