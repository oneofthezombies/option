package option

type Option[T any] struct {
	holder []T
	ok     bool
}

func (o Option[T]) Value() (T, bool) {
	return o.holder[0], o.ok
}

func Some[T any](value T) Option[T] {
	holder := make([]T, 1)
	holder[0] = value

	return Option[T]{
		ok:     true,
		holder: holder,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		ok:     false,
		holder: make([]T, 1),
	}
}
