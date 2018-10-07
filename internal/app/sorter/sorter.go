package sorter

import (
	"math/rand"
)

type HoleCards struct {
	Cards [2]string
}

// func SortRange(handRange []HoleCards, boardCards []string) []HoleCards {
// 	if len(handRange) == 0 {
// 		return []HoleCards{}
// 	}

// 	// Set number runs as 100. Fix the number of computations if it runs slow.
// 	// numRunsToSort := int(math.Min(math.Round(10000.0/float64(len(handRange))), 100))
// 	numRunsToSort := 100

// 	// handRanks contains the rankings of the hands for the random runouts that are played out belows
// 	handRanks := make(map[HoleCards][]int)
// 	for i := 0; i < len(handRange); i++ {
// 		handRanks[handRange[i]] = make([]int, 0, numRunsToSort)
// 	}
// 	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
// 	for i := 0; i < numRunsToSort; i++ {
// 		runout := RandomRunout(boardCards, r)
// 		unexcludedRange := UnexcludedRange(handRange, runout)
// 		hands := BuildHands(unexcludedRange, boardCards, runout)
// 		hands = hand.Sort(hand.SortingHigh, hand.ASC, hands...)
// 	}

// 	res := handRange
// 	return res
// }

// func BuildHands(unexcludedRange []HoleCards, boardCards []string, runout map[string]bool) ([]*hand.Hand, map[string]HoleCards) {
// 	hands := make([]*hand.Hand, 0, len(unexcludedRange))
// 	origHoleCards := make(map[string]HoleCards)

// 	for j := 0; j < len(unexcludedRange); j++ {
// 		fullHandCards := make([]string, len(boardCards))
// 		copy(fullHandCards, boardCards)
// 		fullHandCards = append(fullHandCards, unexcludedRange[j].Cards[:]...)
// 		for card := range runout {
// 			fullHandCards = append(fullHandCards, card)
// 		}

// 		newHand := hand.New(fullHandCards)
// 		[handLength]string(newstrings)
// 		hands = append(hands, hand.New(fullHandCards))
// 	}
// 	return hands, origHoleCards
// }

// func New(cards []Card, options ...func(*Config)) *Hand {
// 	c := &Config{}
// 	for _, option := range options {
// 		option(c)
// 	}
// 	combos := cardCombos(cards)
// 	hands := []*Hand{}
// 	for _, combo := range combos {
// 		hand := handForFiveCards(combo, *c)
// 		hands = append(hands, hand)
// 	}
// 	hands = Sort(c.sorting, DESC, hands...)
// 	hands[0].config = c
// 	return hands[0]
// }

//UnexcludedRange returns hole cards in handRange that are not contained in runout
func UnexcludedRange(handRange []HoleCards, runout map[string]bool) []HoleCards {
	res := make([]HoleCards, 0, len(handRange))
	for i := 0; i < len(handRange); i++ {
		if !runout[handRange[i].Cards[0]] && !runout[handRange[i].Cards[1]] {
			res = append(res, handRange[i])
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

// Python code to classify hands
// def classify_hands(all_hands, board_cards):
//     res = []
//     trey_board_cards = [Card.new(s) for s in board_cards]
//     for hole_cards in all_hands:
//         trey_hole_cards = [Card.new(s) for s in hole_cards]
//         score = EVALUATOR.evaluate(trey_board_cards, trey_hole_cards)
//         res.append(EVALUATOR.class_to_string(EVALUATOR.get_rank_class(score)))
// 		return res

// func RiverValue(holeCards []string, boardCards []string, runout []string) {
// 	allCards := append(append(holeCards, boardCards...), runout...)
// 	resHand := hand.New(allCards)
// }

// def river_value(hole_cards, board_cards, runout):
//     # fails if cards aren't distinct
//     trey_hole_cards = [Card.new(s) for s in hole_cards]
//     trey_board_cards = [Card.new(s) for s in board_cards]
//     trey_runout = [Card.new(s) for s in runout]
//     res = EVALUATOR.evaluate(trey_board_cards + trey_runout, trey_hole_cards)
//     return EVALUATOR.evaluate(trey_board_cards + trey_runout, trey_hole_cards)
