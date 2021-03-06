package expression

import "testing"

func TestChecker(t *testing.T) {
	node := Parse("$(loop range=1,10 => i)").Prune()
	if !(&ExpChecker{Node: node}).IsLoop() {
		t.Error("Error parsing loop.")
	}

	node = Parse("$(pool)").Prune()
	if !(&ExpChecker{Node: node}).IsPool() {
		t.Error("Error parsing pool.")
	}
}
