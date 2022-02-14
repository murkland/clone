package clone

// Interface is what anything cloneable must implement.
type Interface[T any] interface {
	Clone() T
}

// Slice clones a slice of cloneable values.
func Slice[T Interface[T]](xs []T) []T {
	cloned := make([]T, len(xs), cap(xs))
	for i, x := range xs {
		cloned[i] = x.Clone()
	}
	return cloned
}

// Map clones a map of cloneable values.
func Map[T Interface[T], U comparable](xs map[U]T) map[U]T {
	cloned := make(map[U]T, len(xs))
	for k, v := range xs {
		cloned[k] = v.Clone()
	}
	return cloned
}

type basic interface {
	~bool |
		~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// P creates a value pointer from a value. Use with care.
func P[T any](x T) *T {
	return &x
}

// Basic clones (copies) a basic value.
func Basic[T basic](x *T) *T {
	if x == nil {
		return nil
	}
	return P(*x)
}

// ValuePointer clones a value pointer.
func ValuePointer[V any, T interface {
	Interface[*V]
	*V
}](x T) T {
	if x == nil {
		return nil
	}
	return x.Clone()
}
