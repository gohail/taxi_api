package handlers

import (
	"github.com/gohail/taxi_api/internal/proposal/cache"
	"io"
	"net/http"
)

type ProposalHandler struct {
	cacheStore *cache.ProposalCache
}

func NewProposalHandler(cacheStore *cache.ProposalCache) *ProposalHandler {
	return &ProposalHandler{cacheStore: cacheStore}
}

func (h *ProposalHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	p := h.cacheStore.GetRandomProposal()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, p.Name)
}

func (h *ProposalHandler) GetList(w http.ResponseWriter, r *http.Request) {
	p := h.cacheStore.GetProposalList()

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	var resData string

	for _, v := range p {
		resData += v.String() + "\n"
	}
	io.WriteString(w, resData)
}
