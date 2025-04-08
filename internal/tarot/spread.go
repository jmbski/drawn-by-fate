package tarot

type Spread struct {
    Name      string
    Positions []*Card
    Size      int
}

func NewThreeCardSpread(deck *Deck) *Spread {
    spread := &Spread{Name: "Three Card", Size: 3}
    deck.Shuffle()
    for i := 0; i < 3 && len(deck.Cards) > 0; i++ {
        spread.Positions = append(spread.Positions, deck.Draw())
    }
    return spread
}