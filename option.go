package option

type Option[T any] struct {
	value T
	ok    bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		ok:    true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		ok: false,
	}
}

func (o Option[T]) IsSome() bool {
	return o.ok
}

func (o Option[T]) IsNone() bool {
	return !o.ok
}

// Map doesn't seem to fit idiomatic Go IMO
// func Map[T, K any](option Option[T], mapFn func(value T) K) Option[K] {
// 	if IsNone(option) {
// 		return None[K]()
// 	}
// 	return Some(mapFn(option.value))
// }

func (o Option[T]) UnwrapOr(or T) T {
	if o.IsSome() {
		return o.value
	}
	return or
}

func (o Option[T]) Unwrap() (T, bool) {
	if o.IsSome() {
		return o.value, true
	}
	var t T
	return t, false
}
