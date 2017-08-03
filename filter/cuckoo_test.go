package filter

import (
	"testing"

	"github.com/osamingo/shamoji"
	"github.com/stretchr/testify/assert"
)

var blacklist = []string{
	"死ね",
	"教えて",
	"LINE",
}

func TestNewCuckooFilter(t *testing.T) {

	cf := NewCuckooFilter(blacklist...)

	assert.NotNil(t, cf)
	assert.NotNil(t, cf.Cuckoo)
	assert.Equal(t, 3, int(cf.Cuckoo.Count()))
	assert.Implements(t, (*shamoji.Filter)(nil), cf)
}

func TestCuckooFilter_Test(t *testing.T) {

	cf := NewCuckooFilter(blacklist...)

	for i := range blacklist {
		cf.Cuckoo.Insert([]byte(blacklist[i]))
	}
	for i := range blacklist {
		assert.True(t, cf.Test([]byte(blacklist[i])))
	}
}
