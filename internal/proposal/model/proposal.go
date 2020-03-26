package model

import (
	"fmt"
	"sync/atomic"
)

type Proposal struct {
	ViewCount int64  `json:"view_count"`
	Name      string `json:"name"`
}

func NewProposal(name string) *Proposal {
	return &Proposal{ViewCount: 0, Name: name}
}

func (p *Proposal) Increment() int64 {
	return atomic.AddInt64(&p.ViewCount, 1)
}

func (p *Proposal) GetCount() int64 {
	return atomic.LoadInt64(&p.ViewCount)
}

func (p *Proposal) String() string {
	return fmt.Sprintf("%d - %s", p.GetCount(), p.Name)
}
