package libpff

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version() != "20190904" {
		t.Fail()
	}
}
