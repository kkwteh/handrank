package evaluator_test

import (
	"testing"

	"github.com/kkwteh/handrank/internal/app/evaluator"
)

func TestRoyalFlushHandScore(t *testing.T) {
	res := evaluator.HandScore([]string{"Ac", "Kc", "Qc", "Jc", "Tc", "2d", "3h"})
	if res != 1 {
		t.Errorf("Got %v", res)
	}
}

func TestWorstHandScore(t *testing.T) {
	res := evaluator.HandScore([]string{"7c", "5d", "4c", "3h", "2c"})
	if res != 7462 {
		t.Errorf("Got %v", res)
	}
}
