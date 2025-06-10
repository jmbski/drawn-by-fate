package components

type StatModifier struct {
	Value   float32
	StatKey float32
}

type Modifiable struct {
	Modifiers map[string]StatModifier
}
