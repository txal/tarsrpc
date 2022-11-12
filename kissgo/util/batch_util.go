package util

import (
	"sync"
)

type BatchFn func() error

func BatchWork(f ...BatchFn) (err error) {
	if len(f) == 0 {
		return
	}

	var wg sync.WaitGroup
	errs := make([]error, len(f))

	wg.Add(len(f))
	for i := 0; i < len(f); i++ {
		j := i
		go func() {
			defer wg.Done()
			errs[j] = f[j]()
		}()
	}
	wg.Wait()

	for _, v := range errs {
		if v != nil {
			return v
		}
	}
	return
}
