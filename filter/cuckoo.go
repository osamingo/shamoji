package filter

import "github.com/irfansharif/cfilter"

// CuckooFilter filters by cuckoo.
type CuckooFilter struct {
	Cuckoo *cfilter.CFilter
}

// Test implements shamoji.Filter interface.
func (f *CuckooFilter) Test(src []byte) bool {
	return f.Cuckoo.Lookup(src)
}

// NewCuckooFilter generates new CuckooFilter.
func NewCuckooFilter(blacklist ...string) *CuckooFilter {
	cf := cfilter.New(cfilter.Size(uint(len(blacklist))))
	for i := range blacklist {
		cf.Insert([]byte(blacklist[i]))
	}
	return &CuckooFilter{
		Cuckoo: cf,
	}
}
