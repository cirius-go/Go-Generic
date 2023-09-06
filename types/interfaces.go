package types

// MergingHandler represents the merging handler.
type MergingHandler[M any] interface {
	Merge(next M) M
}
