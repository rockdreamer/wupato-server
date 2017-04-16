package game

type Block int8

const (
	_ Block = iota
	Empty
	Blue
	Orange
	Lime
	Magenta
	Cyan
	Yellow
	Red
)

var blockStrings = map[string]Block{
	" ": Empty,
	"B": Blue,
	"O": Orange,
	"L": Lime,
	"M": Magenta,
	"C": Cyan,
	"Y": Yellow,
	"R": Red,
}

func (b Block) String() string {
	for s, v := range blockStrings {
		if b == v {
			return s
		}
	}
	return "invalid"
}

func ParseBlock(s string) Block {
	return blockStrings[s]
}
