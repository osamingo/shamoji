package filter_test

import (
	"testing"

	"github.com/osamingo/shamoji"
	"github.com/osamingo/shamoji/filter"
)

func TestNewCuckooFilter(t *testing.T) {
	t.Parallel()

	cf := filter.NewCuckooFilter("死ね", "教えて", "LINE")

	if cf == nil || cf.Cuckoo == nil {
		t.Error("should not be nil")
	} else if cf.Cuckoo.Count() != 3 {
		t.Error("should be 3")
	}

	var i interface{} = cf
	if _, ok := i.(shamoji.Filter); !ok {
		t.Error("should be implements shamoji.Filter")
	}
}

func TestCuckooFilter_Test(t *testing.T) {
	t.Parallel()

	denyList := []string{"死ね", "教えて", "LINE"}

	cf := filter.NewCuckooFilter(denyList...)
	for i := range denyList {
		cf.Cuckoo.Insert([]byte(denyList[i]))
	}

	for i := range denyList {
		if !cf.Test([]byte(denyList[i])) {
			t.Error("should be true")
		}
	}
}
