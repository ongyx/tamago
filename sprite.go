package tamago

type (
	// Sprite represents a sprite on the screen.
	Sprite struct {
		y, x, tile uint8
		options    Bits
	}
	SpriteData [40]Sprite
)
