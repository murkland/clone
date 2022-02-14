package clone

type Interface[T any] interface {
	Clone() T
}

func Slice[T Interface[T]](xs []T) []T {
	cloned := make([]T, len(xs), cap(xs))
	for i, x := range xs {
		cloned[i] = x.Clone()
	}
	return cloned
}

func Map[T Interface[T], U comparable](xs map[U]T) map[U]T {
	cloned := make(map[U]T, len(xs))
	for k, v := range xs {
		cloned[k] = v.Clone()
	}
	return cloned
}
