package sorter_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/kkwteh/handrank/internal/app/sorter"
)

func TestFullDeck(t *testing.T) {
	res := sorter.FullDeck()
	if len(res) != 52 {
		t.Errorf("full deck does not have 52 cards")
	}
}

func TestRandomRunout(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	res := sorter.RandomRunout([]string{"Ac", "Ad"}, r)
	if len(res) != 3 {
		t.Errorf("runout was wrong length. Expected 3, got %v", len(res))
	}
}

func TestScoreHoleCards(t *testing.T) {
	suitedConnectors := sorter.HoleCards{"5h", "6h"}
	unexcludedRange := []sorter.HoleCards{suitedConnectors}
	boardCards := []string{"5c", "6c", "Ac"}
	runout := map[string]bool{"Qd": true, "6d": true}
	res := sorter.ScoreHoleCards(unexcludedRange, boardCards, runout)

	if res[0].Score != 271 {
		t.Errorf("Got %v", res)
	}

	if res[0].Cards != [2]string{"5h", "6h"} {
		t.Errorf("Got %v", res)
	}
}

func TestSortScoredHoleCards(t *testing.T) {
	handA := sorter.ScoredHoleCards{
		Cards: sorter.HoleCards{"5h", "6h"},
		Score: 162,
	}
	handB := sorter.ScoredHoleCards{
		Cards: sorter.HoleCards{"7h", "2s"},
		Score: 5000,
	}
	handC := sorter.ScoredHoleCards{
		Cards: sorter.HoleCards{"Ac", "Kc"},
		Score: 1,
	}
	handRange := sorter.ScoredRange{handA, handB, handC}
	sort.Sort(handRange)
	if handRange[0].Score != 5000 && handRange[1].Score != 162 && handRange[2].Score != 1 {
		t.Errorf("Got%v", handRange)
	}
}

func TestUnexcludedRange(t *testing.T) {
	suitedConnectors := sorter.HoleCards{"5h", "6h"}
	bigSlick := sorter.HoleCards{"Ac", "Kc"}
	handRange := []sorter.HoleCards{suitedConnectors, bigSlick}
	runout := map[string]bool{"Kc": true}
	res := sorter.UnexcludedRange(handRange, runout)
	if len(res) != 1 && res[0] != suitedConnectors {
		t.Errorf("Got %v", res)
	}
}

func TestClassifyHands(t *testing.T) {
	allHands := []sorter.HoleCards{{"Ac", "Ah"}}
	boardCards := []string{"Ad", "3c", "Qd", "7h", "6s"}
	res := sorter.ClassifyHands(allHands, boardCards)
	if len(res) != 1 {
		t.Errorf("Got %v", len(res))
	}

	if res[0] != "Three Of A Kind" {
		t.Errorf("Got %v", res)
	}
}
