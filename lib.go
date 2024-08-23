package par

import (
	"runtime"
	"sync"
)

type pair[T any, U any] struct {
	a T
	b U
}

func ForN[T any](n int, objs []T, f func(T)) {
	ch := make(chan T)
	wg := new(sync.WaitGroup)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for obj := range ch {
				f(obj)
			}
		}()
	}
	for _, obj := range objs {
		ch <- obj
	}
	close(ch)
	wg.Wait()
}

func MapN[T any, U any, K comparable](n int, objs []T, keyFunc func(T) K, f func(T) U) []U {
	ch := make(chan pair[K, T])
	wg := new(sync.WaitGroup)
	chRes := make(chan pair[K, U], len(objs))
	mp := make(map[K]int)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for obj := range ch {
				chRes <- pair[K, U]{obj.a, f(obj.b)}
			}
		}()
	}
	for i, obj := range objs {
		mp[keyFunc(obj)] = i
		ch <- pair[K, T]{keyFunc(obj), obj}
	}
	close(ch)
	res := make([]U, len(objs))
	for i := 0; i < len(objs); i++ {
		r := <-chRes
		res[mp[r.a]] = r.b
	}
	close(chRes)
	wg.Wait()
	return res
}

func For[T any](objs []T, f func(T)) {
	ForN(runtime.NumCPU(), objs, f)
}

func Map[T any, U any, K comparable](objs []T, keyFunc func(T) K, f func(T) U) []U {
	return MapN(runtime.NumCPU(), objs, keyFunc, f)
}
