package lay_anc_dir_hor

import (
	"uir/assets"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ui *ebitenui.UI
}

func NewGame() *Game {
	child := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(assets.ColorR),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition: widget.AnchorLayoutPositionEnd,
			}),
			widget.WidgetOpts.MinSize(50, 50),
		),
	)
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(assets.ColorL),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(child)

	return &Game{
		ui: &ebitenui.UI{Container: root},
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
