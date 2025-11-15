package types

// Result型
type Result[T any, E any] struct {
	value *T
	err   *E
}

// コンストラクタ
func Ok[T any, E any](value T) Result[T, E] {
	return Result[T, E]{value: &value}
}

func Err[T any, E any](err E) Result[T, E] {
	return Result[T, E]{err: &err}
}

// 判定メソッド
func (r Result[T, E]) IsOk() bool {
	return r.err == nil
}

func (r Result[T, E]) IsErr() bool {
	return r.err != nil
}

// Map系メソッド
func Map[T, U, E any](r Result[T, E], fn func(T) U) Result[U, E] {
	if r.err != nil {
		return Err[U](*r.err)
	}
	return Ok[U, E](fn(*r.value))
}

func MapErr[T, E, F any](r Result[T, E], fn func(E) F) Result[T, F] {
	if r.err != nil {
		return Err[T](fn(*r.err))
	}
	return Ok[T, F](*r.value)
}

func FlatMap[T, U, E any](r Result[T, E], fn func(T) Result[U, E]) Result[U, E] {
	if r.err != nil {
		return Err[U](*r.err)
	}
	return fn(*r.value)
}

// AndThenMap - 型変換を伴うAndThen
func AndThen[T, U, E any](r Result[T, E], fn func(T) Result[U, E]) Result[U, E] {
	return FlatMap(r, fn)
}

// Match
func (r Result[T, E]) Match(onOk func(T), onErr func(E)) {
	if r.err != nil {
		onErr(*r.err)
	} else {
		onOk(*r.value)
	}
}

// Combine - 複数のResultを結合
func Combine[T, E any](results ...Result[T, E]) Result[[]T, E] {
	var values []T
	for _, r := range results {
		if r.err != nil {
			return Err[[]T](*r.err)
		}
		values = append(values, *r.value)
	}
	return Ok[[]T, E](values)
}
