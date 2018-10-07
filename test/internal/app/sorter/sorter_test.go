package sorter_test

import (
	"math/rand"
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

func TestClassifyHands(t *testing.T) {
	allHands := []sorter.HoleCards{{Cards: [2]string{"Ac", "Ah"}}}
	boardCards := []string{"Ad", "3c", "Qd", "7h", "6s"}
	res := sorter.ClassifyHands(allHands, boardCards)
	if len(res) != 1 {
		t.Errorf("Got %v", len(res))
	}

	if res[0] != "Three Of A Kind" {
		t.Errorf("Got %v", res)
	}
}
