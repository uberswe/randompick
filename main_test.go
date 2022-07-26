package randompick

import (
	"math/rand"
	"testing"
	"time"
)

func TestResult(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	i := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	res := Result(i)

	found := false
	for k := range i {
		if res == k {
			found = true
		}
	}

	if found == false {
		t.Errorf("Result not found in map")
	}
}
