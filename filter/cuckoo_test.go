package filter_test

import (
	"testing"

	"github.com/osamingo/shamoji"
	"github.com/osamingo/shamoji/filter"
	"github.com/stretchr/testify/assert"
)

func TestNewCuckooFilter(t *testing.T) {
	cf := filter.NewCuckooFilter("死ね", "教えて", "LINE")

	assert.NotNil(t, cf)
	assert.NotNil(t, cf.Cuckoo)
	assert.Equal(t, 3, int(cf.Cuckoo.Count()))
	assert.Implements(t, (*shamoji.Filter)(nil), cf)
}

func TestCuckooFilter_Test(t *testing.T) {
	blacklist := []string{"死ね", "教えて", "LINE"}

	cf := filter.NewCuckooFilter(blacklist...)

	for i := range blacklist {
		cf.Cuckoo.Insert([]byte(blacklist[i]))
	}
	for i := range blacklist {
		assert.True(t, cf.Test([]byte(blacklist[i])))
	}
}
