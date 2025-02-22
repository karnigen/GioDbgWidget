// go:build exclude

package GioDbgWidget

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// DbgWidget is a debug widget to draw a rectangle
// - Width: width of the border
// - R, G, B, A: color of the border
// - Min, Max: rectangle size

// Usage:
// 	DbgWidget{Width: 1, G: 0xff, Max: dims.Size}.Layout(gtx)

type DbgWidget struct {
	Width      float32
	R, G, B, A uint8 // color.NRGBA
	Min, Max   image.Point
}

// - Init: initialize the widget
//     - if Width <= 0, set to 1
//     - if A == 0, set to 0xff
//     - if R, G, B == 0, R set to 0xff
//     - if Max is not set, set to gtx.Constraints.Max
//     - if Max.X < 0, set to gtx.Constraints.Max.X
//     - if Max.Y < 0, set to gtx.Constraints.Max.Y

func (w *DbgWidget) Init(gtx layout.Context) {
	if w.Width <= 0 {
		w.Width = 1
	}

	if w.A == 0 {
		w.A = 0xff
	}
	if w.R == 0 && w.G == 0 && w.B == 0 {
		w.R = 0xff
	}

	if w.Max == (image.Point{}) { // if Max is not set, set to gtx.Constraints.Max
		w.Max = gtx.Constraints.Max
	} else if w.Max.X < 0 { // for negative values set to gtx.Constraints.Max
		w.Max.X = gtx.Constraints.Max.X
	} else if w.Max.Y < 0 { // for negative values set to gtx.Constraints.Max
		w.Max.Y = gtx.Constraints.Max.Y
	}

}

func (w DbgWidget) Layout(gtx layout.Context) layout.Dimensions {
	w.Init(gtx)
	rect := clip.Rect{Min: w.Min, Max: w.Max}
	// fmt.Println("wdbg: ", w)

	paint.FillShape(gtx.Ops, color.NRGBA{R: w.R, G: w.G, B: w.B, A: w.A},
		clip.Stroke{
			Path:  rect.Path(),
			Width: w.Width,
		}.Op(),
	)
	return layout.Dimensions{Size: w.Max}
}
