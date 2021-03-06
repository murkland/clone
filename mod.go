package clone

// Cloner is what anything cloneable must implement.
type Cloner[T any] interface {
	Clone() T
}

// Slice clones a slice of cloneable values.
func Slice[T Cloner[T]](xs []T) []T {
	cloned := make([]T, len(xs), cap(xs))
	for i, x := range xs {
		cloned[i] = x.Clone()
	}
	return cloned
}

// Map clones a map of cloneable values.
func Map[T Cloner[T], U comparable](xs map[U]T) map[U]T {
	cloned := make(map[U]T, len(xs))
	for k, v := range xs {
		cloned[k] = v.Clone()
	}
	return cloned
}

// P creates a value pointer from a value. Use with care.
func P[T any](x T) *T {
	return &x
}

// Shallow shallow clones (copies) a value.
func Shallow[T any](x *T) *T {
	if x == nil {
		return nil
	}
	return P(*x)
}

// ValuePointer clones a value pointer.
func ValuePointer[V any, T interface {
	Cloner[*V]
	*V
}](x T) T {
	if x == nil {
		return nil
	}
	return x.Clone()
}

// Interface clones an interface pointer.
func Interface[T Cloner[T]](x Cloner[T]) T {
	if x == nil {
		return *new(T)
	}
	return x.Clone()
}
