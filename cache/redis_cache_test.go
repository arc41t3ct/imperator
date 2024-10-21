package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	if err := testRedisCache.Forget("foo"); err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache although it should not have been")
	}

	if err = testRedisCache.Set("foo", "bar"); err != nil {
		t.Error(err)
	}

	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache although it should not have been")
	}
}

func TestRedisCache_Get(t *testing.T) {
	if err := testRedisCache.Set("foo", "bar"); err != nil {
		t.Error(err)
	}

	item, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if item != "bar" {
		t.Error("did not get correct value from cache")
	}
}

func TestRedisCache_Forget(t *testing.T) {
	if err := testRedisCache.Set("alpha", "beta"); err != nil {
		t.Error(err)
	}

	if err := testRedisCache.Forget("alpha"); err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha found in cache although it should not have been")
	}
}

func TestRedisCache_Empty(t *testing.T) {
	if err := testRedisCache.Set("alpha", "beta"); err != nil {
		t.Error(err)
	}

	if err := testRedisCache.Empty(); err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("alpha was found in cache although it should not have been")
	}
}

func TestRedisCache_EmptyMatching(t *testing.T) {
	if err := testRedisCache.Set("alpha", "beta"); err != nil {
		t.Error(err)
	}
	if err := testRedisCache.Set("alpha2", "beta2"); err != nil {
		t.Error(err)
	}
	if err := testRedisCache.Set("omega", "zeta"); err != nil {
		t.Error(err)
	}
	if err := testRedisCache.Set("omega2", "zeta2"); err != nil {
		t.Error(err)
	}

	if err := testRedisCache.EmptyMatching("alpha"); err != nil {
		t.Error(err)
	}

	inCacheAlpha, err := testRedisCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}
	if inCacheAlpha {
		t.Error("alpha was found in cache although it should not have been")
	}

	inCacheAlpha2, err := testRedisCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if inCacheAlpha2 {
		t.Error("alpha2 was found in cache although it should not have been")
	}

	inCacheOmega, err := testRedisCache.Has("omega")
	if err != nil {
		t.Error(err)
	}

	if !inCacheOmega {
		t.Error("omega was not found in cache although it should have been")
	}

	inCacheOmega2, err := testRedisCache.Has("omega2")
	if err != nil {
		t.Error(err)
	}

	if !inCacheOmega2 {
		t.Error("omega2 was not found in cache although it should have been")
	}
}

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"
	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}
	_, err = decode(string(bytes))
	if err != nil {
		t.Error(err)
	}
}
