package tarot

type Suit int

const (
    SuitCups Suit = iota
    SuitSwords
    SuitWands
    SuitPentacles
)

type ArcanaType int

const (
    Major ArcanaType = iota
    Minor
)

type Card struct {
    Name       string
    Arcana     ArcanaType
    Suit       Suit
    Value      int    // 1–10 for Minor, 0 for Major
    Deity      string // Optional: deity associated with this card
    Description string
    IsReversed bool
}