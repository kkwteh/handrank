package sorter

import (
	"math/rand"
	"time"

	"github.com/kkwteh/handrank/internal/app/evaluator"
)

type HoleCards [2]string

type ScoredHoleCards struct {
	Cards [2]string
	Score uint32
}

func SortRange(handRange []HoleCards, boardCards []string) []HoleCards {
	if len(handRange) == 0 {
		return []HoleCards{}
	}

	// Set number runs as 100. Fix the number of computations if it runs slow.
	// numRunsToSort := int(math.Min(math.Round(10000.0/float64(len(handRange))), 100))
	numRunsToSort := 100

	// handRanks contains the rankings of the hands for the random runouts that are played out belows
	handRanks := make(map[HoleCards][]int)
	for i := 0; i < len(handRange); i++ {
		handRanks[handRange[i]] = make([]int, 0, numRunsToSort)
	}
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	for i := 0; i < numRunsToSort; i++ {
		runout := RandomRunout(boardCards, r)
		unexcludedRange := UnexcludedRange(handRange, runout)
		scoredHoleCards := ScoreHoleCards(unexcludedRange, boardCards, runout)
		_ = scoredHoleCards
	}

	res := handRange
	return res
}

// ScoreHoleCards scores hole cards. Lower scores are better.
func ScoreHoleCards(unexcludedRange []HoleCards, boardCards []string, runout map[string]bool) map[HoleCards]uint32 {
	res := make(map[HoleCards]uint32)
	for _, holeCards := range unexcludedRange {
		fullHandCards := make([]string, len(boardCards))
		copy(fullHandCards, boardCards)
		fullHandCards = append(fullHandCards, holeCards[:]...)
		for card := range runout {
			fullHandCards = append(fullHandCards, card)
		}
		res[holeCards] = evaluator.HandScore(fullHandCards)
	}
	return res
}

//UnexcludedRange returns hole cards in handRange that are not contained in runout
func UnexcludedRange(handRange []HoleCards, runout map[string]bool) []HoleCards {
	res := make([]HoleCards, 0, len(handRange))
	for _, hand := range handRange {
		if !runout[hand[0]] && !runout[hand[1]] {
			res = append(res, hand)
		}
	}
	return res
}

func FullDeck() map[string]bool {
	res := make(map[string]bool)
	for _, card := range []string{
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
	} {
		res[card] = true
	}
	return res
}

// RandomRunout returns a random runout
func RandomRunout(boardCards []string, r *rand.Rand) map[string]bool {
	fullDeck := FullDeck()
	for _, card := range boardCards {
		fullDeck[card] = false
	}

	remainingCards := make([]string, 0, len(fullDeck)-len(boardCards))
	for card := range fullDeck {
		if fullDeck[card] == true {
			remainingCards = append(remainingCards, card)
		}
	}
	cardsToDeal := 5 - len(boardCards)
	res := make(map[string]bool)
	perm := r.Perm(len(remainingCards))
	for i := 0; i < cardsToDeal; i++ {
		res[remainingCards[perm[i]]] = true
	}
	return res
}

//ClassifyHands returns a list of hand rankings corresponding to the list of allhands
func ClassifyHands(allHands []HoleCards, boardCards []string) []string {
	res := make([]string, 0, len(allHands))
	for _, hand := range allHands {
		fullHand := append(boardCards, (hand[:])...)
		res = append(res, evaluator.HandRank(evaluator.HandScore(fullHand)))
	}
	return res
}
