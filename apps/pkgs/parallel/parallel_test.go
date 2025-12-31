package parallel

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
)

func TestParallel2Success(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }

		v1, v2, err := Parallel2(ctx, f1, f2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if v1 != 1 || v2 != "b" {
			t.Errorf("unexpected values: %v, %q", v1, v2)
		}
	})
}

func TestParallel2Error(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("error in f1")
		f1 := func() (int, error) { return 0, expectedErr }
		f2 := func() (string, error) { return "b", nil }

		_, _, err := Parallel2(ctx, f1, f2)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestParallel3Success(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }
		f3 := func() (bool, error) { return true, nil }

		v1, v2, v3, err := Parallel3(ctx, f1, f2, f3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if v1 != 1 || v2 != "b" || v3 != true {
			t.Errorf("unexpected values: %v, %q, %v", v1, v2, v3)
		}
	})
}

func TestParallel3Error(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("error in f3")
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }
		f3 := func() (bool, error) { return false, expectedErr }

		_, _, _, err := Parallel3(ctx, f1, f2, f3)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestParallel4Success(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }
		f3 := func() (bool, error) { return true, nil }
		f4 := func() (float64, error) { return 3.14, nil }

		v1, v2, v3, v4, err := Parallel4(ctx, f1, f2, f3, f4)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if v1 != 1 || v2 != "b" || v3 != true || v4 != 3.14 {
			t.Errorf("unexpected values: %v, %q, %v, %v", v1, v2, v3, v4)
		}
	})
}

func TestParallel4Error(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("error in f2")
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "", expectedErr }
		f3 := func() (bool, error) { return true, nil }
		f4 := func() (float64, error) { return 3.14, nil }

		_, _, _, _, err := Parallel4(ctx, f1, f2, f3, f4)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})
}

func TestParallel5Success(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }
		f3 := func() (bool, error) { return true, nil }
		f4 := func() (float64, error) { return 3.14, nil }
		f5 := func() (byte, error) { return 'a', nil }

		v1, v2, v3, v4, v5, err := Parallel5(ctx, f1, f2, f3, f4, f5)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if v1 != 1 || v2 != "b" || v3 != true || v4 != 3.14 || v5 != 'a' {
			t.Errorf("unexpected values: %v, %q, %v, %v, %v", v1, v2, v3, v4, v5)
		}
	})
}

func TestParallel5Error(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx := context.Background()
		expectedErr := errors.New("error in f5")
		f1 := func() (int, error) { return 1, nil }
		f2 := func() (string, error) { return "b", nil }
		f3 := func() (bool, error) { return true, nil }
		f4 := func() (float64, error) { return 3.14, nil }
		f5 := func() (byte, error) { return 0, expectedErr }

		_, _, _, _, _, err := Parallel5(ctx, f1, f2, f3, f4, f5)
		if err != expectedErr {
			t.Errorf("expected error %v, got %v", expectedErr, err)
		}
	})
}
