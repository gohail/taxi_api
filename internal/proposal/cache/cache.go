package cache

import (
	"fmt"
	"github.com/gohail/taxi_api/internal/proposal/model"
)

type ProposalCache struct {
	activeStr ProposalStorage
	adminList ProposalStorage
	gen       *Generator
	cacheSize int
}

func NewProposalCache(cacheSize int) *ProposalCache {
	return &ProposalCache{activeStr: NewProposalStorage(), adminList: NewProposalStorage(), gen: NewGenerator(cacheSize), cacheSize: cacheSize}
}

func (c *ProposalCache) GetRandomProposal() model.Proposal {
	rand := c.gen.RandomId()
	p := c.activeStr.GetByIndex(rand)
	ref := c.activeStr.GetReferByIndex(rand)
	if ref.GetCount() == 1 {
		c.adminList.AddNew(ref)
	}
	return p
}

func (c *ProposalCache) GetProposalList() []*model.Proposal {
	return c.adminList.GetList()
}

func (c *ProposalCache) FillCacheData() {
	for i := 0; i < c.cacheSize; i++ {
		c.activeStr.AddNew(model.NewProposal(c.gen.GenName()))
	}
	fmt.Println("fill cache data: ", len(c.activeStr.list))
}

// Replace old Proposal on a new random with random index
func (c *ProposalCache) Substitute() {
	rand := c.gen.RandomId()
	c.activeStr.SubstituteByIndex(rand, model.NewProposal(c.gen.GenName()))
	//fmt.Println("TIK-TOK")
}
