package ui

import (
	"fyne.io/fyne/v2"
	"github.com/developernajib/learning-go-ztm/apptype"
	"github.com/developernajib/learning-go-ztm/pxcanvas"
	"github.com/developernajib/learning-go-ztm/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
