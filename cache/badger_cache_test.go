package cache

import "testing"

func TestBadgerCache_Has(t *testing.T) {
	if err := testBadgerCache.Forget("foo"); err != nil {
		t.Error(err)
	}
	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache althout it should not be there")
	}

	_ = testBadgerCache.Set("foo", "bar")
	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found although it should be there")
	}
}

func TestBadgerCache_Get(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testBadgerCache.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if x != "bar" {
		t.Error("expected bar to be returned")
	}
}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}
	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("foo found although it should not be there")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("alpha", "delta")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}
	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("alpha found although it should not be there")
	}
}

func TestBadgerCache_EmptyMatching(t *testing.T) {
	err := testBadgerCache.Set("alpha", "delta")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Set("alpha2", "delta2")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Set("theta", "zeta")
	if err != nil {
		t.Error(err)
	}
	err = testBadgerCache.Set("theta2", "zeta2")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyMatching("a")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("alpha found although it should not be there")
	}

	inCache, err = testBadgerCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}
	if inCache {
		t.Error("alpha2 found although it should not be there")
	}
	inCache, err = testBadgerCache.Has("theta")
	if err != nil {
		t.Error(err)
	}
	if !inCache {
		t.Error("theta not found although it should be there")
	}
	inCache, err = testBadgerCache.Has("theta2")
	if err != nil {
		t.Error(err)
	}
	if !inCache {
		t.Error("theta2 not found although it should be there")
	}
}
