package gorocksdb

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/facebookgo/ensure"
)

func TestMemoryUsage(t *testing.T) {
	// create database with cache
	cache := NewLRUCache(8 * 1024 * 1024)
	bbto := NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(cache)
	defer cache.Destroy()

	applyOpts := func(opts *Options) {
		opts.SetBlockBasedTableFactory(bbto)
	}

	db := newTestDB(t, "TestMemoryUsage", applyOpts)
	defer db.Close()

	// take first memory usage snapshot
	mu1, err := GetApproximateMemoryUsageByType([]*DB{db}, []*Cache{cache})
	ensure.Nil(t, err)
	t.Log("mu1:", mu1.MemTableTotal, mu1.MemTableUnflushed, mu1.CacheTotal, mu1.MemTableReadersTotal)

	// perform IO operations that will affect in-memory tables (and maybe cache as well)
	wo := NewDefaultWriteOptions()
	defer wo.Destroy()
	ro := NewDefaultReadOptions()
	defer ro.Destroy()

	key := []byte("key")
	value := make([]byte, mu1.MemTableTotal+1)
	_, err = rand.Read(value)
	ensure.Nil(t, err)

	err = db.Put(wo, key, value)
	ensure.Nil(t, err)
	var value2 *Slice
	value2, err = db.Get(ro, key)
	ensure.Nil(t, err)
	assert.Equal(t, value, value2.Data())

	// take second memory usage snapshot
	mu2, err := GetApproximateMemoryUsageByType([]*DB{db}, []*Cache{cache})
	ensure.Nil(t, err)
	t.Log("mu2:", mu2.MemTableTotal, mu2.MemTableUnflushed, mu2.CacheTotal, mu2.MemTableReadersTotal)

	// the amount of memory used by memtables should increase after write/read;
	// cache memory usage is not likely to be changed, perhaps because requested key is kept by memtable
	assert.True(t, mu2.MemTableTotal > mu1.MemTableTotal)
	assert.True(t, mu2.MemTableUnflushed > mu1.MemTableUnflushed)
	assert.True(t, mu2.CacheTotal >= mu1.CacheTotal)
	assert.True(t, mu2.MemTableReadersTotal >= mu1.MemTableReadersTotal)
}
