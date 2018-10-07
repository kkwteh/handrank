package sorter_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/kkwteh/handrank/internal/app/sorter"
	"github.com/kkwteh/joker/hand"
	"github.com/kkwteh/joker/jokertest"
)

func TestSorter(t *testing.T) {
	_ = sorter.Foo()
}

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

func TestRandomRunout(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	res := sorter.RandomRunout([]hand.Card{1, 2}, r)
	if len(res) != 3 {
		t.Errorf("runout was wrong length. Expected 5, got %v", len(res))
	}
}

func TestClassifyHands(t *testing.T) {
	// ClassifyHands(allHands [][]hand.Card, boardCards []hand.Card) []string {
	res := sorter.ClassifyHands([][]hand.Card{
		{hand.TwoSpades, hand.TwoHearts},
		{hand.AceClubs, hand.KingClubs}},
		[]hand.Card{hand.TwoClubs, hand.FourClubs, hand.SixClubs})
	if res[0] != "ThreeOfAKind" {
		t.Errorf("Expected ThreeOfAKind, got %v", res[0])
	}

	if res[1] != "Flush" {
		t.Errorf("Expected Flush, got %v", res[1])
	}
}

// func TestCards(t *testing.T) {
// 	handResult := sorter.Foo()
// 	if handResult.Ranking() != hand.Ranking(1) {
// 		t.Errorf("Sum was incorrect, got: %v, want: %v.", handResult.Ranking(), hand.Ranking(1))
// 	}
// }
