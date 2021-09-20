package tamago

type (
	// A tile holds a index reference to the color it should be.
	Tile    [8][8]uint8
	Tileset [384]Tile
)
