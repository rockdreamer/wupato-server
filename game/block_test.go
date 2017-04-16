package game_test

import (
	"github.com/cheekybits/is"
	"github.com/rockdreamer/wupato-server/game"
	"testing"
)

func TestEnumerationValues(t *testing.T) {
	is := is.New(t)
	is.Equal(int(game.Empty), 1)
	is.Equal(int(game.Blue), 2)
	is.Equal(int(game.Orange), 3)
	is.Equal(int(game.Lime), 4)
	is.Equal(int(game.Magenta), 5)
	is.Equal(int(game.Cyan), 6)
	is.Equal(int(game.Yellow), 7)
	is.Equal(int(game.Red), 8)
}

func TestSerialization(t *testing.T) {
	is := is.New(t)
	is.Equal(game.Empty.String(), " ")
	is.Equal(game.Blue.String(), "B")
	is.Equal(game.Orange.String(), "O")
	is.Equal(game.Lime.String(), "L")
	is.Equal(game.Magenta.String(), "M")
	is.Equal(game.Cyan.String(), "C")
	is.Equal(game.Yellow.String(), "Y")
	is.Equal(game.Red.String(), "R")

	is.Equal(game.Empty, game.ParseBlock(" "))
	is.Equal(game.Blue, game.ParseBlock("B"))
	is.Equal(game.Orange, game.ParseBlock("O"))
	is.Equal(game.Lime, game.ParseBlock("L"))
	is.Equal(game.Magenta, game.ParseBlock("M"))
	is.Equal(game.Cyan, game.ParseBlock("C"))
	is.Equal(game.Yellow, game.ParseBlock("Y"))
	is.Equal(game.Red, game.ParseBlock("R"))
}
