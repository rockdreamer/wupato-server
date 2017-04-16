package game

type Piece int8

const (
	_ Piece = iota
	Square
	I
	L
	LMirrored
	N
	NMirrored
	T
)

type pieceblocksT [5][5]Block

type PieceRotation int8

type RotatedPiece struct {
	Piece    Piece
	Rotation PieceRotation
}

var pieceblocks = map[Piece]pieceblocksT{
	Square: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Blue, Blue, Empty},
		[5]Block{Empty, Empty, Blue, Blue, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	I: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Orange, Orange, Orange, Orange},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	L: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Lime, Empty, Empty},
		[5]Block{Empty, Empty, Lime, Empty, Empty},
		[5]Block{Empty, Empty, Lime, Lime, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	LMirrored: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Magenta, Empty, Empty},
		[5]Block{Empty, Empty, Magenta, Empty, Empty},
		[5]Block{Empty, Magenta, Magenta, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	N: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Cyan, Empty},
		[5]Block{Empty, Empty, Cyan, Cyan, Empty},
		[5]Block{Empty, Empty, Cyan, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	NMirrored: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Yellow, Empty, Empty},
		[5]Block{Empty, Empty, Yellow, Yellow, Empty},
		[5]Block{Empty, Empty, Empty, Yellow, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
	T: pieceblocksT{
		[5]Block{Empty, Empty, Empty, Empty, Empty},
		[5]Block{Empty, Empty, Red, Empty, Empty},
		[5]Block{Empty, Empty, Red, Red, Empty},
		[5]Block{Empty, Empty, Red, Empty, Empty},
		[5]Block{Empty, Empty, Empty, Empty, Empty},
	},
}
