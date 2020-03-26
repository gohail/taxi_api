package cache

import (
	"github.com/gohail/taxi_api/internal/proposal/model"
	"sync"
)

type ProposalStorage struct {
	mx   sync.RWMutex
	list []*model.Proposal
}

func NewProposalStorage() ProposalStorage {
	return ProposalStorage{}
}

func (ps *ProposalStorage) GetByIndex(index int) model.Proposal {
	ps.mx.Lock()
	defer ps.mx.Unlock()
	p := *ps.list[index]
	ps.list[index].Increment()
	return p
}

func (ps *ProposalStorage) GetReferByIndex(index int) *model.Proposal {
	ps.mx.RLock()
	defer ps.mx.RUnlock()
	return ps.list[index]
}

func (ps *ProposalStorage) SubstituteByIndex(index int, p *model.Proposal) {
	ps.mx.Lock()
	ps.list[index] = p
	ps.mx.Unlock()
}

func (ps *ProposalStorage) GetList() []*model.Proposal {
	ps.mx.RLock()
	defer ps.mx.RUnlock()
	return ps.list
}

func (ps *ProposalStorage) AddNew(p *model.Proposal) {
	ps.mx.Lock()
	ps.list = append(ps.list, p)
	ps.mx.Unlock()
}
