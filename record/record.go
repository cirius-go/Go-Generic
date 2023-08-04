package record

// Record represents the record that contains data as key & value.
type Record[K comparable, V any] map[K]V
