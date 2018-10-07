package sorter

import (
	"math"
	"math/rand"

	"github.com/kkwteh/joker/hand"
)

// def sort_range(hand_range, board_cards):
//     if len(hand_range) == 0:
//         return []

// 		num_runs_to_sort = min(round(10000 / len(hand_range)), 100)

// HoleCards type implements sort interface
type HoleCards struct {
	Cards [2]hand.Card
}

func (hc *HoleCards) Len() int           { return len(hc.Cards) }
func (hc *HoleCards) Less(i, j int) bool { return int(hc.Cards[i]) < int(hc.Cards[j]) }
func (hc *HoleCards) Swap(i, j int)      { hc.Cards[i], hc.Cards[j] = hc.Cards[j], hc.Cards[i] }

func SortRange(handRange []HoleCards, boardCards []hand.Card) []HoleCards {
	if len(handRange) == 0 {
		return []HoleCards{}
	}

	handRanks := make(map[HoleCards][]int)

	numRunsToSort := math.Min(math.Round(10000.0/float64(len(handRange))), 100)
	_ = handRanks
	_ = numRunsToSort
	_ = boardCards
	res := handRange
	return res
}

// FullDeck of 52 cards
func FullDeck() map[hand.Card]bool {
	res := make(map[hand.Card]bool)
	for i := 0; i < 52; i++ {
		res[hand.Card(i)] = true
	}
	return res
}

// RandomRunout returns a random runout
func RandomRunout(boardCards []hand.Card, r *rand.Rand) []hand.Card {
	fullDeck := FullDeck()
	for card := range boardCards {
		fullDeck[hand.Card(card)] = false
	}

	remainingCards := make([]hand.Card, 0, len(fullDeck)-len(boardCards))
	for card := range fullDeck {
		if fullDeck[card] == true {
			remainingCards = append(remainingCards, hand.Card(card))
		}
	}
	cardsToDeal := 5 - len(boardCards)
	res := make([]hand.Card, 0, cardsToDeal)
	perm := r.Perm(len(remainingCards))
	for i := 0; i < cardsToDeal; i++ {
		res = append(res, remainingCards[perm[i]])
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

func ClassifyHands(allHands []HoleCards, boardCards []hand.Card) []string {
	res := make([]string, 0, len(allHands))
	for i := 0; i < len(allHands); i++ {
		fullHand := append(boardCards, (allHands[i].Cards[:])...)
		res = append(res, hand.New(fullHand).Ranking().String())
	}
	return res
}

// func RiverValue(holeCards []hand.Card, boardCards []hand.Card, runout []hand.Card) {
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
