package tarot

import (
    "math/rand"
    "time"
)

type Deck struct {
    Cards []Card
}

func NewDeck() *Deck {
    rand.Seed(time.Now().UnixNano())
    cards := []Card{
        {Name: "Ace of Cups", Arcana: Minor, Suit: SuitCups, Value: 1, Description: "A new emotional start."},
        {Name: "The Fool", Arcana: Major, Value: 0, Description: "Beginning a journey.", Deity: "Zarnak"},
        // Add more cards here...
    }
    return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() {
    rand.Shuffle(len(d.Cards), func(i, j int) {
        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    })
}

func (d *Deck) Draw() *Card {
    if len(d.Cards) == 0 {
        return nil
    }
    card := d.Cards[0]
    d.Cards = d.Cards[1:]
    return &card
}