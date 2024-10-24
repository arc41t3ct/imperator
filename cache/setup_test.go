package cache

import (
	"log"
	"os"
	"testing"
	"time"

	miniredis "github.com/alicebob/miniredis/v2"
	badger "github.com/dgraph-io/badger/v4"
	"github.com/gomodule/redigo/redis"
)

var testRedisCache RedisCache
var testBadgerCache BadgerCache

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	pool := redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", s.Addr())
		},
	}

	testRedisCache.Conn = &pool
	testRedisCache.Prefix = "imperator_tests"
	defer testRedisCache.Conn.Close()

	_ = os.RemoveAll("./testdata/tmp/badger")
	// create a badger db
	if _, err := os.Stat("./testdata/tmp"); os.IsNotExist(err) {
		if err := os.Mkdir("./testdata/tmp", 0755); err != nil {
			log.Fatal(err)
		}
	}
	if err := os.Mkdir("./testdata/tmp/badger", 0755); err != nil {
		log.Fatal(err)
	}

	db, _ := badger.Open(badger.DefaultOptions("./testdata/tmp/badger"))
	testBadgerCache.Conn = db

	os.Exit(m.Run())
}
