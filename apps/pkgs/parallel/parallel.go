package parallel

import "context"

type result[T any] struct {
	val T
	err error
}

type parallelFunc[T any] struct {
	ctx context.Context
	ch  chan result[T]
	f   func() (T, error)
}

func newParallelFunc[T any](ctx context.Context, f func() (T, error)) *parallelFunc[T] {
	return &parallelFunc[T]{ctx: ctx, ch: make(chan result[T], 1), f: f}
}

func (p *parallelFunc[T]) run() {
	defer close(p.ch)
	v, err := p.f()
	p.ch <- result[T]{val: v, err: err}
}

func Parallel2[T1, T2 any](ctx context.Context, f1 func() (T1, error), f2 func() (T2, error)) (T1, T2, error) {
	func1 := newParallelFunc[T1](ctx, f1)
	func2 := newParallelFunc[T2](ctx, f2)

	go func1.run()
	go func2.run()

	var (
		res1 T1
		res2 T2
	)

	// We need to collect 2 results
	// This is a naive implementation that waits for one, then the other, but checks for errors eagerly?
	// Actually better to just use a loop or nested selects carefully?
	// Given the number of permutations, nested selects are verbose but correct if fully expanded.
	// But simply: select on all. Once one returns, select on the others.

	select {
	case <-ctx.Done():
		return res1, res2, ctx.Err()
	case r1 := <-func1.ch:
		if r1.err != nil {
			return res1, res2, r1.err
		}
		res1 = r1.val
		// wait for func2
		select {
		case <-ctx.Done():
			return res1, res2, ctx.Err()
		case r2 := <-func2.ch:
			if r2.err != nil {
				return res1, res2, r2.err
			}
			res2 = r2.val
			return res1, res2, nil
		}
	case r2 := <-func2.ch:
		if r2.err != nil {
			return res1, res2, r2.err
		}
		res2 = r2.val
		// wait for func1
		select {
		case <-ctx.Done():
			return res1, res2, ctx.Err()
		case r1 := <-func1.ch:
			if r1.err != nil {
				return res1, res2, r1.err
			}
			res1 = r1.val
			return res1, res2, nil
		}
	}
}

func Parallel3[T1, T2, T3 any](ctx context.Context, f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error)) (T1, T2, T3, error) {
	func1 := newParallelFunc[T1](ctx, f1)
	func2 := newParallelFunc[T2](ctx, f2)
	func3 := newParallelFunc[T3](ctx, f3)

	go func1.run()
	go func2.run()
	go func3.run()

	var (
		res1 T1
		res2 T2
		res3 T3
	)

	// To avoid combinatorial explosion of nested selects, we can use a simpler approach?
	// But we want to return on FIRST error.
	// We can aggregate channels?
	// Or just stick to the pattern used in Parallel2 but expanded.
	// Expanding for 3 is 3 branches, each having 2 inner branches. 3! permutations? No.
	// We can track completion state.

	// Let's stick to the generated code style but correct.
	// OR: Use a loop with reflect? No reflect.
	// Manual expansion is safest for fixed arity.

	// Helper to wait for the remaining 2
	wait23 := func() (T2, T3, error) {
		select {
		case <-ctx.Done():
			return res2, res3, ctx.Err()
		case r2 := <-func2.ch:
			if r2.err != nil {
				return res2, res3, r2.err
			}
			res2 = r2.val
			// wait for 3
			select {
			case <-ctx.Done():
				return res2, res3, ctx.Err()
			case r3 := <-func3.ch:
				if r3.err != nil {
					return res2, res3, r3.err
				}
				res3 = r3.val
				return res2, res3, nil
			}
		case r3 := <-func3.ch:
			if r3.err != nil {
				return res2, res3, r3.err
			}
			res3 = r3.val
			// wait for 2
			select {
			case <-ctx.Done():
				return res2, res3, ctx.Err()
			case r2 := <-func2.ch:
				if r2.err != nil {
					return res2, res3, r2.err
				}
				res2 = r2.val
				return res2, res3, nil
			}
		}
	}

	// This is getting verbose. Is there a cleaner way?
	// Maybe:
	// We need to consume exactly one item from each channel.
	// If any is error, return error.

	// Since we are inside specific functions (Parallel3), we know the types.

	// Correct implementation for Parallel3:
	select {
	case <-ctx.Done():
		return res1, res2, res3, ctx.Err()
	case r1 := <-func1.ch:
		if r1.err != nil {
			return res1, res2, res3, r1.err
		}
		res1 = r1.val
		res2, res3, err := wait23()
		return res1, res2, res3, err
	case r2 := <-func2.ch:
		if r2.err != nil {
			return res1, res2, res3, r2.err
		}
		res2 = r2.val
		// wait for 1 and 3
		// We can reuse logic?
		// Let's simple write it out.
		select {
		case <-ctx.Done():
			return res1, res2, res3, ctx.Err()
		case r1 := <-func1.ch:
			if r1.err != nil {
				return res1, res2, res3, r1.err
			}
			res1 = r1.val
			select {
			case <-ctx.Done():
				return res1, res2, res3, ctx.Err()
			case r3 := <-func3.ch:
				if r3.err != nil {
					return res1, res2, res3, r3.err
				}
				res3 = r3.val
				return res1, res2, res3, nil
			}
		case r3 := <-func3.ch:
			if r3.err != nil {
				return res1, res2, res3, r3.err
			}
			res3 = r3.val
			select {
			case <-ctx.Done():
				return res1, res2, res3, ctx.Err()
			case r1 := <-func1.ch:
				if r1.err != nil {
					return res1, res2, res3, r1.err
				}
				res1 = r1.val
				return res1, res2, res3, nil
			}
		}
	case r3 := <-func3.ch:
		if r3.err != nil {
			return res1, res2, res3, r3.err
		}
		res3 = r3.val
		// wait for 1 and 2
		select {
		case <-ctx.Done():
			return res1, res2, res3, ctx.Err()
		case r1 := <-func1.ch:
			if r1.err != nil {
				return res1, res2, res3, r1.err
			}
			res1 = r1.val
			select {
			case <-ctx.Done():
				return res1, res2, res3, ctx.Err()
			case r2 := <-func2.ch:
				if r2.err != nil {
					return res1, res2, res3, r2.err
				}
				res2 = r2.val
				return res1, res2, res3, nil
			}
		case r2 := <-func2.ch:
			if r2.err != nil {
				return res1, res2, res3, r2.err
			}
			res2 = r2.val
			select {
			case <-ctx.Done():
				return res1, res2, res3, ctx.Err()
			case r1 := <-func1.ch:
				if r1.err != nil {
					return res1, res2, res3, r1.err
				}
				res1 = r1.val
				return res1, res2, res3, nil
			}
		}
	}
}

// For Parallel4 and Parallel5, this manual expansion is hideous.
// Better to just have a generic way? But types are different.
// We can use flags to mark completion.

func Parallel4[T1, T2, T3, T4 any](ctx context.Context, f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error), f4 func() (T4, error)) (T1, T2, T3, T4, error) {
	func1 := newParallelFunc[T1](ctx, f1)
	func2 := newParallelFunc[T2](ctx, f2)
	func3 := newParallelFunc[T3](ctx, f3)
	func4 := newParallelFunc[T4](ctx, f4)

	go func1.run()
	go func2.run()
	go func3.run()
	go func4.run()

	var (
		res1 T1
		res2 T2
		res3 T3
		res4 T4
	)

	count := 4
	for count > 0 {
		select {
		case <-ctx.Done():
			return res1, res2, res3, res4, ctx.Err()
		case r1 := <-func1.ch:
			if r1.err != nil {
				return res1, res2, res3, res4, r1.err
			}
			res1 = r1.val
			func1.ch = nil // disable this case
			count--
		case r2 := <-func2.ch:
			if r2.err != nil {
				return res1, res2, res3, res4, r2.err
			}
			res2 = r2.val
			func2.ch = nil
			count--
		case r3 := <-func3.ch:
			if r3.err != nil {
				return res1, res2, res3, res4, r3.err
			}
			res3 = r3.val
			func3.ch = nil
			count--
		case r4 := <-func4.ch:
			if r4.err != nil {
				return res1, res2, res3, res4, r4.err
			}
			res4 = r4.val
			func4.ch = nil
			count--
		}
	}
	return res1, res2, res3, res4, nil
}

func Parallel5[T1, T2, T3, T4, T5 any](ctx context.Context, f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error), f4 func() (T4, error), f5 func() (T5, error)) (T1, T2, T3, T4, T5, error) {
	func1 := newParallelFunc[T1](ctx, f1)
	func2 := newParallelFunc[T2](ctx, f2)
	func3 := newParallelFunc[T3](ctx, f3)
	func4 := newParallelFunc[T4](ctx, f4)
	func5 := newParallelFunc[T5](ctx, f5)

	go func1.run()
	go func2.run()
	go func3.run()
	go func4.run()
	go func5.run()

	var (
		res1 T1
		res2 T2
		res3 T3
		res4 T4
		res5 T5
	)

	count := 5
	for count > 0 {
		select {
		case <-ctx.Done():
			return res1, res2, res3, res4, res5, ctx.Err()
		case r1 := <-func1.ch:
			if r1.err != nil {
				return res1, res2, res3, res4, res5, r1.err
			}
			res1 = r1.val
			func1.ch = nil
			count--
		case r2 := <-func2.ch:
			if r2.err != nil {
				return res1, res2, res3, res4, res5, r2.err
			}
			res2 = r2.val
			func2.ch = nil
			count--
		case r3 := <-func3.ch:
			if r3.err != nil {
				return res1, res2, res3, res4, res5, r3.err
			}
			res3 = r3.val
			func3.ch = nil
			count--
		case r4 := <-func4.ch:
			if r4.err != nil {
				return res1, res2, res3, res4, res5, r4.err
			}
			res4 = r4.val
			func4.ch = nil
			count--
		case r5 := <-func5.ch:
			if r5.err != nil {
				return res1, res2, res3, res4, res5, r5.err
			}
			res5 = r5.val
			func5.ch = nil
			count--
		}
	}
	return res1, res2, res3, res4, res5, nil
}
