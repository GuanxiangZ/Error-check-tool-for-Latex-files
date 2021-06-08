package parser

type WordType int

const (
	Text WordType = iota
	Syntax
)

type Word struct {
	Raw  string   `json:"raw"`
	Type WordType `json:"type"`
}
