package main

import (
	"os"

	"github.com/kofoworola/sketchtest/canvas"
)

func main() {
	canv := canvas.NewCanvas(
		canvas.NewFill([2]int{0, 0}, "-"),
		canvas.NewRectangle([2]int{15, 0}, 7, 6, "", "."),
		canvas.NewRectangle([2]int{0, 3}, 8, 4, "O", ""),
		canvas.NewRectangle([2]int{5, 5}, 5, 3, "x", "x"),
	)

	canv.Draw(os.Stdout)
}
