package sorter_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/kkwteh/handrank/internal/app/sorter"
	"github.com/kkwteh/joker/hand"
	"github.com/kkwteh/joker/jokertest"
)

func TestRankingHigh(t *testing.T) {
	handResult := hand.New(jokertest.Cards("Ks", "Qs", "Ac", "2s", "3c"))
	if handResult.Ranking() != hand.HighCard {
		t.Errorf("Ranking was incorrect, got: %v, want: %v.", handResult.Ranking(), hand.HighCard)
	}
}

func TestRankingFlush(t *testing.T) {
	handResult := hand.New(jokertest.Cards("Ks", "Qs", "Ac", "Js", "Tc", "2s", "4s"))
	if handResult.Ranking() != hand.Flush {
		t.Errorf("Ranking was incorrect, got: %v, want: %v.", handResult.Ranking(), hand.Flush)
	}
}

func TestDescription(t *testing.T) {
	handResult := hand.New(jokertest.Cards("Ks", "Qs", "Ac", "Js", "Tc", "2s", "4s"))
	if handResult.Description() != "flush king high" {
		t.Errorf("Description was incorrect, got: %v, want: %v.", handResult.Description(), "flush king high")
	}
}

func TestSort(t *testing.T) {
	hand1 := hand.New(jokertest.Cards("Ks", "Qs", "Ac", "Js", "Tc", "2s", "4s"))
	hand2 := hand.New(jokertest.Cards("Ks", "Qs", "Ac", "Js", "3c", "2d", "4h"))
	res := hand.Sort(hand.SortingHigh, hand.DESC, hand1, hand2)
	if res[0] != hand1 {
		t.Errorf("sort result %v", res)
	}
}

func TestFullDeck(t *testing.T) {
	res := sorter.FullDeck()
	if len(res) != 52 {
		t.Errorf("full deck does not have 52 cards")
	}
}

func TestCardString(t *testing.T) {
	if hand.Card(0).String() != "2♠" {
		t.Errorf("Expected 2♠, got %v", hand.Card(0).String())
	}
}

func TestCardOrder(t *testing.T) {
	holeCards := &sorter.HoleCards{Cards: [2]hand.Card{hand.Card(51), hand.Card(0)}}
	sort.Sort(holeCards)
	if holeCards.Cards[0] != hand.Card(0) {
		t.Errorf("Got unexpected hole cards sort order")
	}
}

func TestHoleCardsLen(t *testing.T) {
	holeCards := &sorter.HoleCards{Cards: [2]hand.Card{hand.Card(51), hand.Card(0)}}
	if holeCards.Len() != 2 {
		t.Errorf("Expected 2, got %v", holeCards.Len())
	}
}

func TestCardLess(t *testing.T) {
	holeCards := &sorter.HoleCards{Cards: [2]hand.Card{hand.Card(51), hand.Card(0)}}
	if holeCards.Less(0, 1) {
		t.Errorf("Expected false, got %v", holeCards.Less(0, 1))
	}
}

func TestRandomRunout(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	res := sorter.RandomRunout([]hand.Card{1, 2}, r)
	if len(res) != 3 {
		t.Errorf("runout was wrong length. Expected 5, got %v", len(res))
	}
}

func TestClassifyHands(t *testing.T) {
	pocketDeuces := sorter.HoleCards{Cards: [2]hand.Card{hand.TwoSpades, hand.TwoHearts}}
	bigSlickSuited := sorter.HoleCards{Cards: [2]hand.Card{hand.AceClubs, hand.KingClubs}}
	res := sorter.ClassifyHands([]sorter.HoleCards{pocketDeuces, bigSlickSuited},
		[]hand.Card{hand.TwoClubs, hand.FourClubs, hand.SixClubs})
	if res[0] != "ThreeOfAKind" {
		t.Errorf("Expected ThreeOfAKind, got %v", res[0])
	}

	if res[1] != "Flush" {
		t.Errorf("Expected Flush, got %v", res[1])
	}
}
