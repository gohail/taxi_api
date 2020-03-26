package main

import (
	"github.com/gohail/taxi_api/internal/proposal/cache"
	"sync"
	"testing"
)

// go test -race race_test.go
func Test_Race(t *testing.T) {
	c := cache.NewProposalCache(50)
	c.FillCacheData()

	var wg sync.WaitGroup
	wg.Add(12)
	for i := 0; i < 4; i++ {
		go func() {
			for i := 0; i < 200000; i++ {
				c.GetRandomProposal()
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 200000; i++ {
				c.GetProposalList()
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < 200000; i++ {
				c.Substitute()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
