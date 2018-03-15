package toolbox

import (
	"testing"
)

func TestNow(t *testing.T) {
	ts := NewTimeService()
	t.Log(ts.Now().String())
}
