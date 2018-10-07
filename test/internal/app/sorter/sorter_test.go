package sorter_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/kkwteh/handrank/internal/app/sorter"
)

func TestSortRange(t *testing.T) {
	suitedConnectors := sorter.HoleCards{"5h", "6h"}
	bigSlick := sorter.HoleCards{"Ac", "Kc"}
	handRange := []sorter.HoleCards{suitedConnectors, bigSlick}
	bigSlickBoard := []string{"7c", "Ad", "Kd"}
	suitedConnectorsBoard := []string{"5s", "6s", "7c"}
	res := sorter.SortRange(handRange, bigSlickBoard)
	if res[1] != bigSlick {
		t.Errorf("Got %v for res", res)
	}

	res2 := sorter.SortRange(handRange, suitedConnectorsBoard)
	if res2[1] != suitedConnectors {
		t.Errorf("Got %v for res2", res2)
	}
}

func TestSortFullDeck(t *testing.T) {
	allHoleCards := make([]sorter.HoleCards, 0, len(sorter.FullDeck())*len(sorter.FullDeck()))
	FullDeckList := []string{
		"As", "Ah", "Ad", "Ac",
		"Ks", "Kh", "Kd", "Kc",
		"Qs", "Qh", "Qd", "Qc",
		"Js", "Jh", "Jd", "Jc",
		"Ts", "Th", "Td", "Tc",
		"9s", "9h", "9d", "9c",
		"8s", "8h", "8d", "8c",
		"7s", "7h", "7d", "7c",
		"6s", "6h", "6d", "6c",
		"5s", "5h", "5d", "5c",
		"4s", "4h", "4d", "4c",
		"3s", "3h", "3d", "3c",
		"2s", "2h", "2d", "2c",
	}

	boardCards := make(map[string]bool)
	boardCards["As"] = true
	boardCards["Ks"] = true
	boardCards["Qs"] = true

	for i, cardi := range FullDeckList {
		for j, cardj := range FullDeckList {
			if i < j && boardCards[cardi] == false && boardCards[cardj] == false {
				allHoleCards = append(allHoleCards, sorter.HoleCards{cardi, cardj})
			}
		}
	}
	res := sorter.SortRange(allHoleCards, []string{"As", "Ks", "Qs"})
	if len(res) != 1176 {
		t.Errorf("Got %v", len(res))
	}
}

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
