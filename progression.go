package d2s

// progression struct will represent the characters in game progress of acts.
type progression byte

type readableProgression struct {
	Title string `json:"title"`
	Value int    `json:"value"`
}

// Readable will return a readable format for the progression byte.
func (p progression) Readable(c class, isExpansion bool) readableProgression {
	return readableProgression{}
}
