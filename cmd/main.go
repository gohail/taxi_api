package main

import (
	"fmt"
	"github.com/gohail/taxi_api/internal/proposal/cache"
	"github.com/gohail/taxi_api/internal/proposal/handlers"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Start taxi_api server")

	proposalCache := cache.NewProposalCache(50)
	proposalCache.FillCacheData()

	handler := handlers.NewProposalHandler(proposalCache)

	http.HandleFunc("/request", handler.GetOne)
	http.HandleFunc("/admin/requests", handler.GetList)

	ticker := time.NewTicker(200 * time.Millisecond)
	go func() {
		for range ticker.C {
			proposalCache.Substitute()
		}
	}()

	fmt.Println("started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
