package cards

const (
	Clubs    Suit = "clubs"
	Diamonds Suit = "diamonds"
	Hearts   Suit = "hearts"
	Spades   Suit = "spades"
)

type (
	Suit string

	Stack struct {
		Cards []*Card
	}
	Card struct {
		// 2-10, 11, 12, 13, 14.
		Rank int  `json:"rank"`
		Suit Suit `json:"suit"`
	}
)

func NewStack() *Stack {
	return &Stack{
		Cards: NewShuffledCards(),
	}
}

func NewShuffledCards() []*Card {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	deck := make([]*Card, 0, 52)
	for _, s := range suits {
		for rank := 1; rank <= 13; rank++ {
			deck = append(deck, &Card{Rank: rank, Suit: s})
		}
	}
	return Shuffle(deck)
}

func (s *Stack) Deal(n int) []*Card {
	if n < 0 {
		n = 0
	}
	if n > len(s.Cards) {
		n = len(s.Cards)
	}
	addToHand := s.Cards[:n]
	s.Cards = s.Cards[n:]
	return addToHand
}
