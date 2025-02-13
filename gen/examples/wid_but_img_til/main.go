// START root
package wid_but_img_til

import (
	"bytes"
	"gen/assets"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

type Game struct {
	ui   *ebitenui.UI
	face text.Face
}

func NewGame() *Game {
	face := DefaultFont()
	buttonImage, _ := loadButtonImage()
	button := widget.NewButton(
		widget.ButtonOpts.TextLabel("Button"),
		widget.ButtonOpts.TextFace(face),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    assets.ColorL,
			Hover:   assets.ColorL,
			Pressed: assets.ColorL,
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(assets.ColorD),
			Hover:   image.NewNineSliceColor(assets.ColorD),
			Pressed: image.NewNineSliceColor(assets.ColorG),
		}),
		widget.ButtonOpts.Image(buttonImage),
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
			image.NewNineSliceColor(assets.ColorL),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(button)

	return &Game{
		ui:   &ebitenui.UI{Container: root},
		face: DefaultFont(),
	}
}

// START this
func loadButtonImage() (*widget.ButtonImage, error) {
	idleImage := ebiten.NewImage(100, 100)
	vector.DrawFilledCircle(idleImage, float32(idleImage.Bounds().Dx())/2, float32(idleImage.Bounds().Dy())/2, 45, assets.ColorB, true)
	idle := image.NewNineSlice(idleImage, [3]int{0, idleImage.Bounds().Dx(), 0}, [3]int{0, idleImage.Bounds().Dy(), 0})

	hoverImage := ebiten.NewImage(100, 100)
	vector.DrawFilledCircle(hoverImage, float32(hoverImage.Bounds().Dx())/2, float32(hoverImage.Bounds().Dy())/2, 45, assets.ColorD, true)
	hover := image.NewNineSlice(hoverImage, [3]int{0, hoverImage.Bounds().Dx(), 0}, [3]int{0, hoverImage.Bounds().Dy(), 0})

	pressedImage := ebiten.NewImage(100, 100)
	vector.DrawFilledCircle(pressedImage, float32(pressedImage.Bounds().Dx())/2, float32(pressedImage.Bounds().Dy())/2, 45, assets.ColorG, true)
	pressed := image.NewNineSlice(pressedImage, [3]int{0, pressedImage.Bounds().Dx(), 0}, [3]int{0, pressedImage.Bounds().Dy(), 0})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

// END this

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
