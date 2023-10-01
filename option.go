package option

type Option[T any] struct {
	value *T
}

func (o Option[T]) Value() T {
	if o.value == nil {
		panic("value is nil. use Has() to check if value is present")
	}

	return *o.value
}

func (o Option[T]) Has() bool {
	return o.value != nil
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
	}
}
