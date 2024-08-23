package par_test

import (
	"sync/atomic"
	"testing"

	"github.com/nekoite/par"
	"github.com/stretchr/testify/assert"
)

func TestFor(t *testing.T) {
	assert := assert.New(t)
	objs := make([]int64, 50000)
	for i := range objs {
		objs[i] = int64(i)
	}
	sum := int64(0)
	par.For(objs, func(obj int64) {
		atomic.AddInt64(&sum, obj)
	})
	assert.EqualValues(1249975000, sum)
}

func TestMap(t *testing.T) {
	assert := assert.New(t)
	objs := make([]int64, 5000)
	for i := range objs {
		objs[i] = int64(i)
	}
	res := par.Map(objs, func(obj int64) int64 { return obj }, func(obj int64) int64 { return obj * obj })
	for i, obj := range objs {
		assert.Equal(obj*obj, res[i])
	}
}

func TestForN(t *testing.T) {
	assert := assert.New(t)
	objs := make([]int64, 50000)
	for i := range objs {
		objs[i] = int64(i)
	}
	sum := int64(0)
	par.ForN(4, objs, func(obj int64) {
		atomic.AddInt64(&sum, obj)
	})
	assert.EqualValues(1249975000, sum)
}

func TestMapN(t *testing.T) {
	assert := assert.New(t)
	objs := make([]int64, 5000)
	for i := range objs {
		objs[i] = int64(i)
	}
	res := par.MapN(4, objs, func(obj int64) int64 { return obj }, func(obj int64) int64 { return obj * obj })
	for i, obj := range objs {
		assert.Equal(obj*obj, res[i])
	}
}
