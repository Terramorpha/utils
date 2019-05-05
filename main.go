package utils

import (
	"fmt"
	"image/color"
)

func SetForeground(text string, col color.RGBA) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", col.R, col.G, col.B, text)
}
