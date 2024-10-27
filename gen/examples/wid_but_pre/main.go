// START root
package wid_but_pre

import (
	"bytes"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

type Game struct {
	ui   *ebitenui.UI
	face text.Face
}

func NewGame() *Game {
	face := DefaultFont()
	button := widget.NewButton(
		widget.ButtonOpts.TextLabel("Button"),
		widget.ButtonOpts.TextFace(face),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    colornames.White,
			Hover:   colornames.Gainsboro,
			Pressed: colornames.White,
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(colornames.Darkslategray),
			Hover:   image.NewNineSliceColor(colornames.Darkslategray),
			Pressed: image.NewNineSliceColor(colornames.Mediumseagreen),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(200, 40),
		),
	)
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Gainsboro),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(button)

	return &Game{
		ui:   &ebitenui.UI{Container: root},
		face: DefaultFont(),
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func DefaultFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   20,
	}
}

// END root
